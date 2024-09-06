package main

import (
	"context"
	"fmt"
)

func main() {
	// Создаем контекст с функцией отмены.
	ctx, cancel := context.WithCancel(context.Background())
	dataCh := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// ctx.Done() возвращает канал, который будет закрыт после
			// отмены контекста. Из закрытого канала будет получено нулевое
			// значение и горутина завершит свою работу.
			select {
			case <-ctx.Done():
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
	// Отменяем контекст.
	cancel()
}