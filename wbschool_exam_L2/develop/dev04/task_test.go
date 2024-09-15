package main

import (
	"reflect"
	"testing"
)

// Тест для проверки корректной работы функции с набором анаграмм
func TestFindAnagramsBasic(t *testing.T) {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	expected := map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

	result := findAnagrams(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Тест для проверки пустого ввода
func TestFindAnagramsEmpty(t *testing.T) {
	words := []string{}
	expected := map[string][]string{}

	result := findAnagrams(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Тест для проверки, что одиночные слова не попадают в результат
func TestFindAnagramsSingleWords(t *testing.T) {
	words := []string{"пятак", "слово", "слиток", "столик", "листок", "тяпка"}
	expected := map[string][]string{
		"пятак":  {"пятак", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

	result := findAnagrams(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Тест для проверки работы с одинаковыми словами
func TestFindAnagramsDuplicateWords(t *testing.T) {
	words := []string{"пятак", "пятак", "тяпка", "листок", "слиток", "листок", "столик"}
	expected := map[string][]string{
		"пятак":  {"пятак", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

	result := findAnagrams(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Тест для проверки работы с уже отсортированными словами
func TestFindAnagramsSorted(t *testing.T) {
	words := []string{"тяпка", "пятка", "пятак"}
	expected := map[string][]string{
		"пятак": {"пятак", "пятка", "тяпка"},
	}

	result := findAnagrams(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
