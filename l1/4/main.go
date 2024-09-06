package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
)

// Воркер завершит работы только тогда, когда канал будет закрыт
// и все данные из канала будут прочитаны.
func worker(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Stop worker")
	for {
		v, ok := <-ch
		if !ok {
			return
		}
		fmt.Println(v)
	}
}

// Фунцкия для запуска необходимого количества воркеров.
func startWorkers(n int, ch <-chan int, wg *sync.WaitGroup) {
	fmt.Printf("Start %d workers\n", n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ch, wg)
	}
}

// Функция для получения из консоли от пользователя количества
// воркеров к запуску.
func GetNumberOfWorkers() int {
	for {
		fmt.Print("Enter the number of workers:")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Enter a valid number greater than 0")
			continue
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter a valid number greater than 0")
			continue
		}
		if num < 1 {
			fmt.Println("Enter a valid number greater than 0")
			continue
		}
		return num
	}
}

func main() {
	// Cоздание канала для сигналов ОС.
	syscallChan := make(chan os.Signal, 1)
	// Регистрация syscallChan для получения уведомлений об SIGINT.
	signal.Notify(syscallChan, os.Interrupt)
	// Cоздание объекта WaitGroup для синхронизации горутин.
	var wg sync.WaitGroup
	// Получаем число воркеров к запуску.
	n := GetNumberOfWorkers()
	// Канал для постоянной записи данных.
	dataChan := make(chan int, n*2)
	// Запускаем воркеры.
	startWorkers(n, dataChan, &wg)
	// При получении значения из канала syscallChan,
	// канал dataChan закрывается, основная горутина заблокирована,
	// пока все воркеры дочитывают данные из канала и заверщаются,
	// после чего программа заверщает свою работу.
	for i := 0; ; i++ {
		select {
		case <-syscallChan:
			close(dataChan)
			wg.Wait()
			return
		case dataChan <- i:
		}
	}
}