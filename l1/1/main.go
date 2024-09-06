package main

import (
	"fmt"
)

// определяем структуру Human  и  методы для нее
type Human struct {
	name     string
	lastName string
	age      int
	sex      string
}

func (h Human) Add(a int) int {
	return a + 10
}

func (h *Human) Rename(newName string) {
	h.name = newName
}

// определяем структуру Action
type Action struct {
	Human
}

func (a Action) Add(number int) int {
	return number + 20
}

func main() {
	action := Action{
		Human{
			age:      10,
			name:     "alex",
			sex:      "man",
			lastName: "morti",
		},
	}
	// используе поля из  human
	action.age = 20
	// используем метод из human
	action.Rename("ka")
	fmt.Print(action.name)
	// В случае коллизии - приоритет за "дочерним"
	fmt.Print(action.Add(8))

	// можно явная вызвать нужный метод
	fmt.Print(action.Human.Add(5))
}
