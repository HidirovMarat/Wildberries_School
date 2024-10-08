package main

import (
	"fmt"  // для форматированного ввода/вывода
	"sync" // для синхронизации (RWMutex)
)

func main() {
	// Структура для хранения счетчика с доступом для чтения и записи
	var counter = struct {
		sync.RWMutex
		m map[int]int
	}{m: make(map[int]int)}

	// Запуск 10000 горутин для записи в счетчик
	for i := 0; i < 10000; i++ {
		go func(i int) {
			// Блокировка счетчика для записи
			counter.Lock()
			defer counter.Unlock() // Разблокировка после выхода из функции
			counter.m[0] = i       // Установка значения счетчика по ключу 0 на текущее значение i
		}(i)
	}

	// Запуск 10000 горутин для чтения из счетчика
	for i := 0; i < 10000; i++ {
		go func() {
			// Блокировка счетчика только для чтения
			counter.RLock()
			defer counter.RUnlock()   // Разблокировка после выхода из функции
			fmt.Println(counter.m[0]) // Печать текущего значения счетчика по ключу 0
		}()
	}
}
