package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// Тест для createEventHandler
func TestCreateEventHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/create_event", strings.NewReader("user_id=1&title=Meeting&date=2024-09-18"))
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createEventHandler)

	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус %v, но получен %v", http.StatusOK, status)
	}

	// Проверяем содержимое ответа
	expected := `{"result":"Event created"}`
	if rr.Body.String() != expected {
		t.Errorf("Ожидался ответ %v, но получен %v", expected, rr.Body.String())
	}
}

// Тест для updateEventHandler
func TestUpdateEventHandler(t *testing.T) {
	// Сначала создаем событие для теста
	events[1] = Event{ID: 1, UserID: 1, Title: "Old Meeting", Date: parseDateOrPanic("2024-09-18")}

	req, err := http.NewRequest("POST", "/update_event", strings.NewReader("id=1&title=Updated Meeting&date=2024-09-19"))
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateEventHandler)

	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус %v, но получен %v", http.StatusOK, status)
	}

	// Проверяем содержимое ответа
	expected := `{"result":"Event updated"}`
	if rr.Body.String() != expected {
		t.Errorf("Ожидался ответ %v, но получен %v", expected, rr.Body.String())
	}
}

// Тест для deleteEventHandler
func TestDeleteEventHandler(t *testing.T) {
	// Сначала создаем событие для теста
	events[2] = Event{ID: 2, UserID: 1, Title: "Test Event", Date: parseDateOrPanic("2024-09-18")}

	req, err := http.NewRequest("POST", "/delete_event", strings.NewReader("id=2"))
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteEventHandler)

	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус %v, но получен %v", http.StatusOK, status)
	}

	// Проверяем содержимое ответа
	expected := `{"result":"Event deleted"}`
	if rr.Body.String() != expected {
		t.Errorf("Ожидался ответ %v, но получен %v", expected, rr.Body.String())
	}

	// Проверяем, что событие действительно удалено
	if _, exists := events[2]; exists {
		t.Errorf("Ожидалось, что событие будет удалено, но оно все еще существует")
	}
}

// Тест для eventsForDayHandler
func TestEventsForDayHandler(t *testing.T) {
	// Сначала создаем несколько событий
	events[3] = Event{ID: 3, UserID: 1, Title: "Meeting", Date: parseDateOrPanic("2024-09-18")}
	events[4] = Event{ID: 4, UserID: 1, Title: "Workshop", Date: parseDateOrPanic("2024-09-18")}

	req, err := http.NewRequest("GET", "/events_for_day?date=2024-09-18", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(eventsForDayHandler)

	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус %v, но получен %v", http.StatusOK, status)
	}

	// Проверяем содержимое ответа
	expected := `[{"id":3,"user_id":1,"title":"Meeting","date":"2024-09-18T00:00:00Z","created_at":"0001-01-01T00:00:00Z"},{"id":4,"user_id":1,"title":"Workshop","date":"2024-09-18T00:00:00Z","created_at":"0001-01-01T00:00:00Z"}]`
	if rr.Body.String() != expected {
		t.Errorf("Ожидался ответ %v, но получен %v", expected, rr.Body.String())
	}
}

// Вспомогательная функция для парсинга даты или паники
func parseDateOrPanic(dateStr string) time.Time {
	date, err := parseDate(dateStr)
	if err != nil {
		panic(err)
	}
	return date
}
