package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Структура для хранения флагов
type SortOptions struct {
	column      int  // -k: колонка для сортировки
	numeric     bool // -n: сортировка по числовому значению
	reverse     bool // -r: сортировка в обратном порядке
	unique      bool // -u: вывод только уникальных строк
	month       bool // -M: сортировка по названию месяца
	ignoreSpace bool // -b: игнорирование хвостовых пробелов
	checkSorted bool // -c: проверка отсортированности
	humanSort   bool // -h: сортировка с учетом суффиксов
}

// Чтение строк из стандартного ввода
func readLines() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Функция сортировки с поддержкой различных ключей
func sortLines(lines []string, opts SortOptions) ([]string, error) {
	// Если задано -u (уникальные строки)
	if opts.unique {
		lines = uniqueLines(lines)
	}

	// Если -M (сортировка по месяцу)
	if opts.month {
		sort.SliceStable(lines, func(i, j int) bool {
			return compareMonths(lines[i], lines[j], opts)
		})
	} else if opts.numeric {
		// Если -n (сортировка по числовому значению)
		sort.SliceStable(lines, func(i, j int) bool {
			return compareNumeric(lines[i], lines[j], opts)
		})
	} else {
		// Обычная строковая сортировка
		sort.SliceStable(lines, func(i, j int) bool {
			return compareStrings(lines[i], lines[j], opts)
		})
	}

	// Если -r (обратный порядок)
	if opts.reverse {
		reverseSlice(lines)
	}

	return lines, nil
}

// Удаление дубликатов строк
func uniqueLines(lines []string) []string {
	uniqueMap := make(map[string]bool)
	var uniqueLines []string
	for _, line := range lines {
		if !uniqueMap[line] {
			uniqueMap[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}

// Сравнение строк по месяцу
func compareMonths(a, b string, opts SortOptions) bool {
	monthFormat := "Jan"
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)
	monthA, errA := time.Parse(monthFormat, a[:3])
	monthB, errB := time.Parse(monthFormat, b[:3])

	if errA != nil || errB != nil {
		return false
	}
	return monthA.Before(monthB)
}

// Сравнение строк по числовому значению
func compareNumeric(a, b string, opts SortOptions) bool {
	trimmedA := strings.TrimSpace(a)
	trimmedB := strings.TrimSpace(b)
	numA, errA := strconv.ParseFloat(trimmedA, 64)
	numB, errB := strconv.ParseFloat(trimmedB, 64)

	if errA != nil || errB != nil {
		return false
	}
	return numA < numB
}

// Обычное сравнение строк
func compareStrings(a, b string, opts SortOptions) bool {
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)
	if opts.column > 0 {
		a = getColumn(a, opts.column)
		b = getColumn(b, opts.column)
	}
	return a < b
}

// Получить нужную колонку
func getColumn(s string, col int) string {
	fields := strings.Fields(s)
	if col-1 < len(fields) {
		return fields[col-1]
	}
	return ""
}

// Обратный порядок слайса
func reverseSlice(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func main() {
	var opts SortOptions

	// Флаги командной строки
	flag.IntVar(&opts.column, "k", 0, "column for sorting")
	flag.BoolVar(&opts.numeric, "n", false, "sort by numeric value")
	flag.BoolVar(&opts.reverse, "r", false, "sort in reverse order")
	flag.BoolVar(&opts.unique, "u", false, "output only unique lines")
	flag.BoolVar(&opts.month, "M", false, "sort by month name")
	flag.BoolVar(&opts.ignoreSpace, "b", false, "ignore trailing spaces")
	flag.BoolVar(&opts.checkSorted, "c", false, "check if the input is sorted")
	flag.BoolVar(&opts.humanSort, "h", false, "sort by numeric value with suffixes")
	flag.Parse()

	// Чтение строк с ввода
	lines, err := readLines()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Если флаг -c (проверка отсортированности)
	if opts.checkSorted {
		if sort.SliceIsSorted(lines, func(i, j int) bool { return compareStrings(lines[i], lines[j], opts) }) {
			fmt.Println("The input is sorted")
		} else {
			fmt.Println("The input is not sorted")
		}
		return
	}

	// Сортировка строк
	sortedLines, err := sortLines(lines, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sorting lines: %v\n", err)
		os.Exit(1)
	}

	// Вывод результата
	for _, line := range sortedLines {
		fmt.Println(line)
	}
}
