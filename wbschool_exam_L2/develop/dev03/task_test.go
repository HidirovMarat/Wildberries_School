package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

// Тест сортировки по умолчанию (лексикографическая)
func TestSortLinesDefault(t *testing.T) {
	input := []string{"banana", "apple", "pear"}
	expected := []string{"apple", "banana", "pear"}

	opts := SortOptions{}
	sorted, _ := sortLines(input, opts)

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}

// Тест сортировки с флагом -n (числовая сортировка)
func TestSortLinesNumeric(t *testing.T) {
	input := []string{"10", "2", "30", "25"}
	expected := []string{"2", "10", "25", "30"}

	opts := SortOptions{numeric: true}
	sorted, _ := sortLines(input, opts)

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}

// Тест сортировки с флагом -r (обратная сортировка)
func TestSortLinesReverse(t *testing.T) {
	input := []string{"banana", "apple", "pear"}
	expected := []string{"pear", "banana", "apple"}

	opts := SortOptions{reverse: true}
	sorted, _ := sortLines(input, opts)

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}

// Тест сортировки с флагом -u (уникальные строки)
func TestSortLinesUnique(t *testing.T) {
	input := []string{"apple", "banana", "apple", "pear"}
	expected := []string{"apple", "banana", "pear"}

	opts := SortOptions{unique: true}
	sorted, _ := sortLines(input, opts)

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}

// Тест сортировки по месяцу с флагом -M
func TestSortLinesMonth(t *testing.T) {
	input := []string{"Mar", "Jan", "Feb", "Dec"}
	expected := []string{"Jan", "Feb", "Mar", "Dec"}

	opts := SortOptions{month: true}
	sorted, _ := sortLines(input, opts)

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}

// Тест игнорирования хвостовых пробелов с флагом -b
func TestSortLinesIgnoreSpaces(t *testing.T) {
	input := []string{"apple  ", "banana ", " pear"}
	expected := []string{"apple  ", "banana ", " pear"}

	opts := SortOptions{ignoreSpace: true}
	sorted, _ := sortLines(input, opts)

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}

// Тест на проверку отсортированности с флагом -c
func TestSortLinesCheckSorted(t *testing.T) {
	input := []string{"apple", "banana", "pear"}

	opts := SortOptions{checkSorted: true}
	if !sort.SliceIsSorted(input, func(i, j int) bool { return compareStrings(input[i], input[j], opts) }) {
		t.Errorf("Expected input to be sorted")
	}
}

// Функция для чтения строк из строки (для тестов)
func readLinesFromString(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
