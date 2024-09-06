package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	// фиксируем время
	start := time.Now()
	// цикл без условия for true
	for {
		// если время прошло, то выходим из цикла и завершаем функцию
		if time.Since(start) >= duration {
			break
		}
	}
}

func main() {
	fmt.Println("3 секунды")
	sleep(10 * time.Second)
	fmt.Println("ВСЕ!!!!")
}
