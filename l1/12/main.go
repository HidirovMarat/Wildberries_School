package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	// создаем множество
	setOfstrs := make(map[string]bool)
	// итерируемся по последовательности строк, устанавливаем строку в качестве
	// ключа карты setOfstrs и присваиваем ключу значение в true, поскольку
	// карты не допускают дублирования ключей, получаем имитация множества.
	for _, v := range strs {
		setOfstrs[v] = true
	}
	// выводим элементы множества в консоль
	for k := range setOfstrs {
		fmt.Printf("%s ", k)
	}
	fmt.Printf("\n")
}
