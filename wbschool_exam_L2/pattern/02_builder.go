package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн «Строитель»
Описание: Строитель (Builder) — порождающий паттерн проектирования, который позволяет пошагово создавать сложные объекты.

Применимость:

Когда объект может иметь множество конфигураций.
Когда требуется создать сложный объект пошагово.
Плюсы:

Упрощает создание сложных объектов.
Позволяет пошагово строить объект.
Минусы:

Увеличивает количество кода из-за введения дополнительных классов.
Пример использования:

Построение сложного объекта, такого как HTML-документ, или объект конфигурации.
*/

import "fmt"

// Продукт
type House struct {
	walls   string
	roof    string
	floors  int
	garage  bool
}

// Интерфейс строителя
type Builder interface {
	BuildWalls()
	BuildRoof()
	BuildFloors()
	BuildGarage()
	GetResult() House
}

// Конкретный строитель
type ConcreteBuilder struct {
	house House
}

func (b *ConcreteBuilder) BuildWalls() {
	b.house.walls = "Brick walls"
}

func (b *ConcreteBuilder) BuildRoof() {
	b.house.roof = "Metal roof"
}

func (b *ConcreteBuilder) BuildFloors() {
	b.house.floors = 2
}

func (b *ConcreteBuilder) BuildGarage() {
	b.house.garage = true
}

func (b *ConcreteBuilder) GetResult() House {
	return b.house
}

// Директор
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) Construct() {
	d.builder.BuildWalls()
	d.builder.BuildRoof()
	d.builder.BuildFloors()
	d.builder.BuildGarage()
}

func builder() {
	builder := &ConcreteBuilder{}
	director := NewDirector(builder)
	director.Construct()

	house := builder.GetResult()
	fmt.Printf("House: %+v\n", house)
}
