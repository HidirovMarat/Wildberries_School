package main

import (
	"testing"
	"time"
)

// Функция-утилита для создания канала, который закроется через заданное время
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func TestOr(t *testing.T) {
	start := time.Now()

	// Используем функцию or для объединения нескольких каналов
	done := or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second), // Этот канал должен закрыться первым
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	<-done // Ожидаем завершения

	// Проверяем, что прошло примерно 1 секунда (канал должен закрыться через 1 секунду)
	elapsed := time.Since(start)
	if elapsed < 1*time.Second || elapsed > 1*time.Second+100*time.Millisecond {
		t.Errorf("Ожидалось, что функция завершится за примерно 1 секунду, но прошло %v", elapsed)
	}
}
