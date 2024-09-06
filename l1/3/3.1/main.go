package main

import (
	"fmt"
	"sync"
)

func m(a chan<- int, nums []int) {
	// закрываем канал, чтобы не было deadlock 
	defer close(a)
	var wg sync.WaitGroup

	for _, num := range nums {
		// добавлям в общий счетчик колическов горутин для wait
		wg.Add(1)
		
		go func (x int)  {
			// именьшаем счетчик горутин
			defer wg.Done()
			a <- x * x
		}(num) 
	}
	// ждем чтобы все горутины завершились
	wg.Wait()
	
}

func main() {
	a := make(chan int)
	nums := []int{1, 2, 3, 4, 5, 10}
	// так как wait ждет окончание всех горутин нам надо отправить его отдельную горутину 
	go m(a, nums)

	sum := 0
	// считаем из канала и добавлем в sum
	for val := range a {
		sum += val
	}

	fmt.Print(sum)
}
