package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Тест функции downloadPage
func TestDownloadPage(t *testing.T) {
	// Создаем тестовый HTTP сервер
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body>Test page</body></html>"))
	}))
	defer ts.Close()

	// Загрузка страницы с тестового сервера
	body, err := downloadPage(ts.URL)
	if err != nil {
		t.Fatalf("Ошибка загрузки страницы: %v", err)
	}

	expected := "<html><body>Test page</body></html>"
	if body != expected {
		t.Errorf("Ожидалось: %s, получено: %s", expected, body)
	}

	// Проверяем, что файл был создан
	if _, err := os.Stat("index.html"); os.IsNotExist(err) {
		t.Errorf("Файл index.html не был создан")
	}

	// Удаляем файл после теста
	os.Remove("index.html")
}

// Тест функции parseHTML и загрузки ресурсов
func TestParseHTML(t *testing.T) {
	// Создаем тестовый HTTP сервер
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Write([]byte(`<html><body><img src="/test.png"></body></html>`))
		} else if r.URL.Path == "/test.png" {
			w.Write([]byte("image data"))
		}
	}))
	defer ts.Close()

	// Загружаем главную страницу
	body, err := downloadPage(ts.URL)
	if err != nil {
		t.Fatalf("Ошибка загрузки страницы: %v", err)
	}

	// Парсим HTML и загружаем ресурсы
	err = parseHTML(ts.URL, body)
	if err != nil {
		t.Errorf("Ошибка парсинга HTML: %v", err)
	}

	// Проверяем, что ресурс был загружен
	if _, err := os.Stat("test.png"); os.IsNotExist(err) {
		t.Errorf("Файл test.png не был загружен")
	}

	// Удаляем файлы после теста
	os.Remove("index.html")
	os.Remove("test.png")
}

// Тест функции resolveURL для преобразования относительных ссылок в абсолютные
func TestResolveURL(t *testing.T) {
	baseURL := "http://example.com"

	tests := []struct {
		resourceURL string
		expected    string
	}{
		{"/test.png", "http://example.com/test.png"},
		{"http://example.com/test.png", "http://example.com/test.png"},
	}

	for _, test := range tests {
		result := resolveURL(baseURL, test.resourceURL)
		if result != test.expected {
			t.Errorf("Ожидалось: %s, получено: %s", test.expected, result)
		}
	}
}
