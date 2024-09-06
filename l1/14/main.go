package main

import (
	"fmt"
	"reflect"
)

func defineType(t interface{}) {

	switch t.(type) {
	case int:
		fmt.Println("This is int")
	case string:
		fmt.Println("This is string")
	case bool:
		fmt.Println("This is bool")
	case chan interface{}:
		fmt.Println("This is chan interface{}")
	case chan int:
		fmt.Println("This is chan int")
	case chan bool:
		fmt.Println("This is chan bool")
	case chan string:
		fmt.Println("This is chan string")
	default:
		//Можно просто использовать %T, она выводит значения типа)
		fmt.Printf("This is %T \n", t)
	}
}

func main() {
	// interface
	var a interface{}
	defineType(a)
	// int
	a = 5
	defineType(a)
	//string
	a = "d"
	defineType(a)
	// bool
	a = true
	defineType(a)
	// chan int
	a = make(chan int)
	defineType(a)
	// chan string
	a = make(chan string)
	defineType(a)
	// float64
	a = 5.5
	defineType(a)

	// можно использовать reflect.TypeOf(any)

	a = make(chan float32)

	fmt.Print("reflect.TypeOf - ", reflect.TypeOf(a))
}
