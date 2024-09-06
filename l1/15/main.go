/*
Утечка памяти:

Когда вы создаете срез строки v[:100], этот срез ссылается на тот же массив байтов, что и исходная строка v. Это означает, что даже если срез очень маленький, он будет удерживать в памяти всю большую строку v, что может привести к значительным утечкам памяти, особенно если createHugeString создает очень большие строки.
Потенциальная паника:

Если длина строки, возвращаемой createHugeString, меньше 100, попытка создать срез v[:100] вызовет панику.
*/

package main

import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	// Создаем строку указанного размера
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1 << 10)

	// Проверяем, что строка достаточно длинная
	if len(v) < 100 {
		justString = v
	} else {
		// Создаем новый срез, чтобы избежать утечек памяти
		justString = string(v[:100])
	}
}

func main() {
	someFunc()
	fmt.Println(justString) // Выводим строку, чтобы убедиться, что все работает
}
