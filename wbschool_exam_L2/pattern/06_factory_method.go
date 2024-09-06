package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Паттерн "Фабричный метод" определяет интерфейс для создания объектов, но позволяет подклассам изменять тип создаваемых объектов.
// Применимость:
// - Подходит, когда необходимо делегировать создание объектов подклассам.
// - Используется, когда объекты создаваемого класса могут изменяться в зависимости от логики.
//
// Плюсы:
// + Позволяет избегать жесткой привязки к конкретным классам продуктов.
// + Облегчает добавление новых типов продуктов.
// Минусы:
// - Может усложнить код, если слишком много подклассов.
//
// Пример использования: создание объектов в зависимости от условий, фреймворки для GUI.

import "fmt"

// Transport - интерфейс для всех видов транспорта
type Transport interface {
	Deliver()
}

// Truck представляет грузовик как транспортное средство
type Truck struct{}

// Deliver реализует доставку грузовиком
func (t *Truck) Deliver() {
	fmt.Println("Delivery by truck")
}

// Ship представляет корабль как транспортное средство
type Ship struct{}

// Deliver реализует доставку кораблем
func (s *Ship) Deliver() {
	fmt.Println("Delivery by ship")
}

// Logistics - интерфейс логистики, который определяет фабричный метод CreateTransport
type Logistics interface {
	CreateTransport() Transport
}

// RoadLogistics реализует логистику для дорог
type RoadLogistics struct{}

// CreateTransport создает транспорт - грузовик
func (r *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

// SeaLogistics реализует логистику для моря
type SeaLogistics struct{}

// CreateTransport создает транспорт - корабль
func (s *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}

func factoryMethod() {
	var logistics Logistics

	// Используем дорожную логистику
	logistics = &RoadLogistics{}
	transport := logistics.CreateTransport()
	transport.Deliver() // Доставка грузовиком

	// Используем морскую логистику
	logistics = &SeaLogistics{}
	transport = logistics.CreateTransport()
	transport.Deliver() // Доставка кораблем
}
