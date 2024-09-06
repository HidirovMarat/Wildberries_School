package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	//делаем все символы строчными
	s = strings.ToLower(s)

	set := make(map[rune]bool)
	// если такой ключ встречался то возращаем false, иначе создаем такоей ключ.
	for _, v := range s {
		if ok := set[v]; ok {
			return false
		}
		set[v] = true
	}
	// если цикл прошел до конца то возращаем true
	return true
}

func main() {
	f := "abcdD"

	fmt.Println(f, " ", isUnique(f))

	f = "abCdefAaf"

	fmt.Println(f, " ", isUnique(f))

	f = "aabcd"

	fmt.Println(f, " ", isUnique(f))
}
