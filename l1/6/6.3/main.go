package main

import (
	"fmt"
)

// При заверщении основной горутины, остальные горутины также завершатся.
func main() {
	dataCh := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			dataCh <- i
		}
		close(dataCh)
	}()
	for i := range dataCh {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
