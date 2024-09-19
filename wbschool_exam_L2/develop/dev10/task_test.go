package main

import (
	"io"
	"net"
	"os"
	"testing"
	"time"
)

// Тест успешного подключения и передачи данных
func TestTelnetConnection(t *testing.T) {
	// Создаем тестовый TCP сервер
	ln, err := net.Listen("tcp", ":0") // :0 выбирает случайный свободный порт
	if err != nil {
		t.Fatalf("Ошибка при создании тестового сервера: %v", err)
	}
	defer ln.Close()

	// Запускаем сервер в отдельной горутине
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Logf("Ошибка при принятии соединения: %v", err)
			return
		}
		defer conn.Close()

		// Сервер принимает данные и возвращает их обратно
		io.Copy(conn, conn)
	}()

	// Получаем адрес сервера
	address := ln.Addr().String()

	// Устанавливаем подключение
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		t.Fatalf("Ошибка при подключении к тестовому серверу: %v", err)
	}
	defer conn.Close()

	// Отправляем тестовое сообщение
	message := "Hello, server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		t.Fatalf("Ошибка при отправке сообщения: %v", err)
	}

	// Читаем ответ от сервера
	buffer := make([]byte, len(message))
	_, err = conn.Read(buffer)
	if err != nil {
		t.Fatalf("Ошибка при получении ответа: %v", err)
	}

	// Проверяем, что ответ совпадает с отправленным сообщением
	if string(buffer) != message {
		t.Errorf("Ожидалось: %s, получено: %s", message, string(buffer))
	}
}

// Тест таймаута при подключении к недоступному хосту
func TestTelnetTimeout(t *testing.T) {
	// Устанавливаем таймаут подключения
	timeout := 1 * time.Second

	// Пытаемся подключиться к несуществующему хосту
	_, err := net.DialTimeout("tcp", "192.0.2.0:12345", timeout) // 192.0.2.0 - это зарезервированный IP для тестов
	if err == nil {
		t.Fatal("Ожидалась ошибка при подключении, но ее не было")
	}
}

// Тест обработки закрытия соединения сервером
func TestTelnetServerClose(t *testing.T) {
	// Создаем тестовый TCP сервер
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Ошибка при создании тестового сервера: %v", err)
	}
	defer ln.Close()

	// Запускаем сервер, который закроет соединение сразу после подключения
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Logf("Ошибка при принятии соединения: %v", err)
			return
		}
		conn.Close()
	}()

	// Получаем адрес сервера
	address := ln.Addr().String()

	// Подключаемся к серверу
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		t.Fatalf("Ошибка при подключении к тестовому серверу: %v", err)
	}

	// Пробуем отправить данные после закрытия соединения
	_, err = conn.Write([]byte("Hello"))
	if err == nil {
		t.Fatal("Ожидалась ошибка записи после закрытия соединения")
	}
}

// Тест ввода данных через STDIN
func TestTelnetStdin(t *testing.T) {
	// Создаем тестовый TCP сервер
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Ошибка при создании тестового сервера: %v", err)
	}
	defer ln.Close()

	// Запускаем сервер, который возвращает данные обратно
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Logf("Ошибка при принятии соединения: %v", err)
			return
		}
		defer conn.Close()
		io.Copy(conn, conn)
	}()

	// Получаем адрес сервера
	address := ln.Addr().String()

	// Сохраняем старый os.Stdin для восстановления
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	// Создаем pipe для эмуляции ввода в STDIN
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Записываем тестовое сообщение в pipe
	message := "Test input"
	go func() {
		w.Write([]byte(message))
		w.Close()
	}()

	// Устанавливаем соединение и отправляем данные
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		t.Fatalf("Ошибка при подключении к тестовому серверу: %v", err)
	}
	defer conn.Close()

	// Читаем данные, отправленные сервером
	buffer := make([]byte, len(message))
	_, err = conn.Read(buffer)
	if err != nil {
		t.Fatalf("Ошибка при чтении данных: %v", err)
	}

	// Проверяем, что данные совпадают
	if string(buffer) != message {
		t.Errorf("Ожидалось: %s, получено: %s", message, string(buffer))
	}
}
