package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// UnpackString осуществляет распаковку строки с повторами символов и поддержкой escape-последовательностей
func UnpackString(input string) (string, error) {
	var result strings.Builder
	runes := []rune(input)
	escaped := false

	for i := 0; i < len(runes); i++ {
		curr := runes[i]

		// Если встретили escape (\), то переходим в режим escaped
		if curr == '\\' && !escaped {
			escaped = true
			continue
		}

		// Если текущий символ цифра, проверяем, является ли он частью escape-последовательности или нет
		if unicode.IsDigit(curr) {
			if i == 0 || (!escaped && unicode.IsDigit(runes[i-1])) {
				return "", errors.New("invalid string: starts with or contains consecutive numbers without letters")
			}

			count, _ := strconv.Atoi(string(curr))
			result.WriteString(strings.Repeat(string(runes[i-1]), count-1)) // -1, т.к. 1 раз символ уже добавлен
			escaped = false
			continue
		}

		// Если это обычный символ или escaped символ
		result.WriteRune(curr)
		escaped = false
	}

	return result.String(), nil
}

func main() {
	// Пример использования
	unpacked, err := UnpackString(`a4bc2d5e`)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Unpacked string:", unpacked)
	}
}
