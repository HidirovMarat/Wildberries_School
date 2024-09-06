package main

import "fmt"

func main() {
	a := 10
	b := 50

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a = a + b
	// Так как в переменной 'a' хранится a + b , то b = a - b <=> ((a + b) - b = a). То есть в 'b' теперь хранится значение 'a' 
	b = a - b
	// Так как в переменной 'a' хранится a + b, а в переменной 'b' значение 'a', то a = a - b <=> ((a + b) - a = b). То есть в 'b' теперь хранится значение 'a'
	a = a - b

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}