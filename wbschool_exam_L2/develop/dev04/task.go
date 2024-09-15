package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Функция для сортировки рун в слове
func sortedString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Функция поиска множеств анаграмм
func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)
	usedWords := make(map[string]bool)

	for _, word := range words {
		loweredWord := strings.ToLower(word)
		sorted := sortedString(loweredWord)

		// Проверяем, если слово уже использовалось, пропускаем его
		if usedWords[loweredWord] {
			continue
		}

		// Если в мапе уже есть анаграммы для этого ключа (отсортированное слово), добавляем новое слово
		if _, exists := anagramMap[sorted]; exists {
			anagramMap[sorted] = append(anagramMap[sorted], loweredWord)
		} else {
			// Если это первое слово с данным ключом, создаем новую запись
			anagramMap[sorted] = []string{loweredWord}
		}

		// Помечаем слово как использованное
		usedWords[loweredWord] = true
	}

	// Создадим результирующую мапу, удаляя множества из одного элемента и сортируя списки анаграмм
	result := make(map[string][]string)
	for _, group := range anagramMap {
		if len(group) > 1 {
			sort.Strings(group)
			result[group[0]] = group
		}
	}

	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слиток", "столик"}
	anagrams := findAnagrams(words)
	for key, group := range anagrams {
		fmt.Printf("%s: %v\n", key, group)
	}
}
