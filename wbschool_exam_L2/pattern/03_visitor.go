package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Паттерн «Посетитель»
Описание: Посетитель (Visitor) — поведенческий паттерн, который позволяет добавлять новые операции к существующим объектам, не изменяя их структуру.

Применимость:

Когда нужно добавить новую операцию к классам без изменения их структуры.
Когда есть множество классов с разными интерфейсами, и требуется выполнить над ними одну и ту же операцию.
Плюсы:

Упрощает добавление новых операций без изменения существующих классов.
Минусы:

Усложняет поддержку кода при большом количестве посетителей и классов.
Пример использования:

Поддержка разных типов объектов в дереве (например, файловая система), к которым нужно применить одну и ту же операцию.
go
*/

import "fmt"

// Интерфейс посетителя
type Visitor interface {
	VisitElementA(*ElementA)
	VisitElementB(*ElementB)
}

// Элемент интерфейс
type Element interface {
	Accept(Visitor)
}

// Конкретный элемент A
type ElementA struct{}

func (e *ElementA) Accept(v Visitor) {
	v.VisitElementA(e)
}

// Конкретный элемент B
type ElementB struct{}

func (e *ElementB) Accept(v Visitor) {
	v.VisitElementB(e)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitElementA(e *ElementA) {
	fmt.Println("Visited ElementA")
}

func (v *ConcreteVisitor) VisitElementB(e *ElementB) {
	fmt.Println("Visited ElementB")
}

func visitor() {
	elements := []Element{&ElementA{}, &ElementB{}}
	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}
