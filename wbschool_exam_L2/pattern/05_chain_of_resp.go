package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Паттерн "Цепочка вызовов" позволяет передавать запрос последовательно через цепочку обработчиков, пока один из них не обработает его.
// Применимость:
// - Подходит, когда нужно передавать запрос по цепочке обработчиков без явной привязки к какому-либо из них.
// - Используется в логировании, обработке событий, фильтрации данных.
//
// Плюсы:
// + Уменьшает зависимость между отправителем и получателем запроса.
// + Позволяет динамически изменять цепочку обработчиков.
// Минусы:
// - Нет гарантии, что запрос будет обработан.
//
// Пример использования: обработка событий в графическом интерфейсе, фильтрация запросов, обработка ошибок.

import "fmt"

// Handler определяет интерфейс для обработки запроса
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// BaseHandler реализует базовый функционал установки следующего обработчика
type BaseHandler struct {
	next Handler
}

// SetNext задает следующий обработчик в цепочке
func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// Handle передает запрос следующему обработчику, если он есть
func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

// ConcreteHandlerA обрабатывает запросы типа "A"
type ConcreteHandlerA struct {
	BaseHandler
}

// Handle обрабатывает запросы "A" или передает их дальше
func (h *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("Handler A handled the request")
	} else {
		h.BaseHandler.Handle(request)
	}
}

// ConcreteHandlerB обрабатывает запросы типа "B"
type ConcreteHandlerB struct {
	BaseHandler
}

// Handle обрабатывает запросы "B" или передает их дальше
func (h *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("Handler B handled the request")
	} else {
		h.BaseHandler.Handle(request)
	}
}

func chainOfResp() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	// Устанавливаем цепочку: handlerA -> handlerB
	handlerA.SetNext(handlerB)

	// Пробуем обработать разные запросы
	handlerA.Handle("A") // Обработается Handler A
	handlerA.Handle("B") // Обработается Handler B
	handlerA.Handle("C") // Не будет обработано
}
