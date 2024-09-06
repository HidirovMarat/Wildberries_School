package main

import "fmt"

type SomeInterface interface {
	ReqiredMethod()
}

// Функция принимающая аргументом переменную реализующую SomeInterface.
func SomeFunction(i SomeInterface) {
	i.ReqiredMethod()
}

// Структура с необходимым методом, но под других названием.
type StructWithReqiredLogic struct {
	num int
}

func (s StructWithReqiredLogic) ReqiredMethodWithAnotherName() {
	fmt.Println(s.num)
}

// В структуру адаптер встаривается структура с необходимой логикой.
type Adapter struct {
	StructWithReqiredLogic
}

// Структуру адаптер реализует SomeInterface, внутри необходимого метода,
// вызывает метод встроенной структуры выполняющий необходимую логику.
func (a Adapter) ReqiredMethod() {
	a.ReqiredMethodWithAnotherName()
}

func main() {
	notSuitableStructure := StructWithReqiredLogic{10}
	// SomeFunction(notSuitableStructure) - не будет работать,
	// переменная типа StructWithReqiredLogic не соответствует
	// интерфейсу SomeInterface.
	suitableStructure := Adapter{notSuitableStructure}
	// suitableStructure соответствует
	// интерфейсу SomeInterface.
	SomeFunction(suitableStructure)
}
