package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Определение флагов
	after := flag.Int("A", 0, "печатать N строк после совпадения")
	before := flag.Int("B", 0, "печатать N строк до совпадения")
	context := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество совпадающих строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "инвертировать совпадения")
	fixed := flag.Bool("F", false, "точное совпадение со строкой")
	lineNum := flag.Bool("n", false, "печатать номер строки")
	flag.Parse()

	// Получаем паттерн (или точную строку) для поиска
	pattern := flag.Arg(0)
	if pattern == "" {
		fmt.Println("Ошибка: отсутствует строка для поиска")
		os.Exit(1)
	}

	// Если установлен флаг -C (контекст), применяем его ко флагам -A и -B
	if *context > 0 {
		*after = *context
		*before = *context
	}

	// Открываем стандартный ввод для чтения
	scanner := bufio.NewScanner(os.Stdin)

	// Конфигурация для игнорирования регистра
	var re *regexp.Regexp
	var err error
	if *fixed {
		// Для флага -F используем точное совпадение
		if *ignoreCase {
			pattern = strings.ToLower(pattern)
		}
	} else {
		// Для паттерна используем регулярное выражение
		flags := ""
		if *ignoreCase {
			flags = "(?i)"
		}
		re, err = regexp.Compile(flags + pattern)
		if err != nil {
			fmt.Println("Ошибка компиляции регулярного выражения:", err)
			os.Exit(1)
		}
	}

	// Чтение строк из стандартного ввода и обработка флагов
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Переменная для подсчета количества совпадений
	matchCount := 0

	// Функция для проверки совпадения строки
	matches := func(line string) bool {
		if *fixed {
			if *ignoreCase {
				return strings.ToLower(line) == pattern
			}
			return line == pattern
		}
		return re.MatchString(line)
	}

	// Обрабатываем строки с учетом флагов
	for i, line := range lines {
		matched := matches(line)

		// Инвертирование совпадений, если установлен флаг -v
		if *invert {
			matched = !matched
		}

		if matched {
			matchCount++
			// Выводим строку и дополнительные строки (флаги -A, -B, -C)
			if *count {
				continue
			}

			// Печать строк до совпадения (-B)
			start := i - *before
			if start < 0 {
				start = 0
			}

			// Печать строк после совпадения (-A)
			end := i + *after
			if end >= len(lines) {
				end = len(lines) - 1
			}

			for j := start; j <= end; j++ {
				if *lineNum {
					fmt.Printf("%d:", j+1)
				}
				fmt.Println(lines[j])
			}
		}
	}

	// Если установлен флаг -c, выводим количество совпадений
	if *count {
		fmt.Println(matchCount)
	}
}
