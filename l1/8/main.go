package main

import "fmt"

// true если 1, false если 0 надо
func SetBit(n int64, i uint, set bool) int64 {
	if set {
		// мы сдвигаем 1 => 00..1..000 такой вид
		// затем используем | и в позиции i будет 1
		return n | (1 << i)
	}
	// &^ это and not
	// 1 &^ 0 = 1
	// 1 &^ 1 = 0
	// 0 &^ 1 = 0
	// 0 &^ 0 = 0
	// мы сдвигаем 1 => 00..1..000 такой вид и где 0  остается таким же как и первый значения, а в индексе i будет 0
	return n &^ (1 << i)
}

func main() {
	var n int64 = 4
	var i uint = 2
	fmt.Print(SetBit(n, i, false))
}
