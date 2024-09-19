package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	// Парсинг аргументов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут подключения")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Использование: go-telnet --timeout=10s <host> <port>")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	// Устанавливаем таймаут для подключения
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Printf("Ошибка подключения: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Подключено к %s\n", address)

	// Канал для завершения программы
	done := make(chan struct{})

	// Горутина для чтения из соединения и вывода в STDOUT
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			fmt.Println("Ошибка чтения из соединения:", err)
		}
		fmt.Println("\nСоединение закрыто сервером")
		done <- struct{}{}
	}()

	// Горутина для чтения из STDIN и отправки данных в соединение
	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			fmt.Println("Ошибка отправки данных:", err)
		}
		done <- struct{}{}
	}()

	// Ожидание завершения одной из горутин
	<-done
	fmt.Println("Программа завершена")
}
