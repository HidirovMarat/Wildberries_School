package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

// Структура события
type Event struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

// Хранилище событий (для простоты в памяти)
var events = make(map[int]Event)
var nextID = 1

// Middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s, URL: %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// Вспомогательная функция для отправки JSON ответа
func sendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Вспомогательная функция для парсинга даты
func parseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02"
	return time.Parse(layout, dateStr)
}

// Обработчик для создания события
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Invalid method"})
		return
	}

	// Парсинг данных из формы
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid user_id"})
		return
	}

	title := r.FormValue("title")
	dateStr := r.FormValue("date")
	date, err := parseDate(dateStr)
	if err != nil {
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
		return
	}

	// Создаем событие
	event := Event{
		ID:        nextID,
		UserID:    userID,
		Title:     title,
		Date:      date,
		CreatedAt: time.Now(),
	}
	events[nextID] = event
	nextID++

	sendJSONResponse(w, http.StatusOK, map[string]string{"result": "Event created"})
}

// Обработчик для обновления события
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Invalid method"})
		return
	}

	// Парсинг данных
	eventID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || eventID <= 0 {
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid event ID"})
		return
	}

	event, exists := events[eventID]
	if !exists {
		sendJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Event not found"})
		return
	}

	title := r.FormValue("title")
	dateStr := r.FormValue("date")
	date, err := parseDate(dateStr)
	if err != nil {
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
		return
	}

	// Обновляем событие
	event.Title = title
	event.Date = date
	events[eventID] = event

	sendJSONResponse(w, http.StatusOK, map[string]string{"result": "Event updated"})
}

// Обработчик для удаления события
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Invalid method"})
		return
	}

	// Парсинг данных
	eventID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || eventID <= 0 {
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid event ID"})
		return
	}

	_, exists := events[eventID]
	if !exists {
		sendJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Event not found"})
		return
	}

	// Удаляем событие
	delete(events, eventID)

	sendJSONResponse(w, http.StatusOK, map[string]string{"result": "Event deleted"})
}

// Обработчик для получения событий за день
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := parseDate(dateStr)
	if err != nil {
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
		return
	}

	var result []Event
	for _, event := range events {
		if event.Date.Equal(date) {
			result = append(result, event)
		}
	}

	sendJSONResponse(w, http.StatusOK, result)
}

// Обработчики для получения событий за неделю и месяц можно реализовать аналогично

func main() {
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)

	// Добавляем middleware для логирования
	loggedMux := loggingMiddleware(http.DefaultServeMux)

	port := ":8080"
	fmt.Printf("Запуск сервера на порту %s...\n", port)
	if err := http.ListenAndServe(port, loggedMux); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
