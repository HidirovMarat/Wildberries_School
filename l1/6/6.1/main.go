package main

import "fmt"

func main() {
	// Канал для управляемого заверщения горутины.
	done := make(chan struct{})
	// Функция отмены, которая закроет канал Done.
	cancel := func() {
		close(done)
	}
	dataCh := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// чтение из открытого канала приостанавливается до появления в нем
			// данных, а чтение из закрытого канала всегда возвращает нулевое
			// значение для используемого каналом типа. Это означает, что
			// выполнение ветви, отвечающей за чтение из канала done, будет
			// приостановлено до закрытия этого канала.
			select {
			case <-done:
				return
			default:
				dataCh <- i
			}
		}
		close(dataCh)
	}()
	for i := range dataCh {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	// Закрываем канал done с помощью функции отмены.
	cancel()
}