package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

const DefaultNtp = "pool.ntp.org"

// GetNTPTime — функция, получающая точное время с NTP-сервера, принимает адрес сервера
// или пустую строку аргументом. Если функции передана пустая строка, будет использован
// сервер по-умолчанию.
func GetNTPTime(ntpServer string) (time.Time, error) {
	if ntpServer == "" {
		ntpServer = DefaultNtp
	}
	return ntp.Time(ntpServer)
}

func main() {
	t, err := GetNTPTime("")
	if err != nil {
		log.Fatalf("NTP time fetch error: %s", err.Error())
	}
	fmt.Println("Current OS time:  " + time.Now().Format(time.UnixDate))
	fmt.Println("Precise NTP time: " + t.Format(time.UnixDate))
}
