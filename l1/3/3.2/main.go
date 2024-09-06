package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	sum := int64(0)

	for _, v := range a {
		// запускаем горутины и в каждой вычисляем квадрат числа чтобы прибавить к sum
		go func(val int64) {
			// используем atomic.AddInt64 для атомарного прибавление, то есть потокобезопасного 
			atomic.AddInt64(&sum, val*val)
		}(int64(v))
	}
	// ЖДЕМ
	time.Sleep(time.Second)
	fmt.Print(sum)
}
