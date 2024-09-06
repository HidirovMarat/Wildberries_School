package main

import (
	"fmt"
	"sync"
	"time"
)

// Определяем тип структуры-счетчика.
type CountStruct struct {
	count int
	// Мьютекс для структуры-счетчика.
	mu sync.Mutex
}

// Метод для конкурентного инкрементирования структуры-счетчика.
func (c *CountStruct) ConcurrentIncrement() {
	// Устанавливаем блокировку мьютекса, изменять структуры сможет только
	// одна горутина.
	c.mu.Lock()
	// Снимаем блокировку мьютекса после завершения метода.
	defer c.mu.Unlock()
	c.count++
}

func main() {
	// Создаем указатель на структуру
	countStruct := new(CountStruct)
	// Запускаем горутины для инкрементирования счетчика структуры в
	// конкурентной среде.
	for i := 0; i < 2000; i++ {
		go countStruct.ConcurrentIncrement()
	}
	// Ждем одну секунду чтобы усели завершиться другие горутины.
	time.Sleep(time.Second)
	// Выводим итоговый результат в консоль.
	fmt.Println(countStruct.count)
}