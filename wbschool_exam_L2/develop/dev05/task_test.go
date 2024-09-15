package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// Функция для запуска main с тестовыми данными и флагами
func runGrep(input string, args ...string) string {
	// Создаем pipe для захвата вывода
	r, w, _ := os.Pipe()

	// Сохраняем стандартный вывод и ввод
	oldStdout := os.Stdout
	oldStdin := os.Stdin

	// Подменяем стандартный ввод и вывод
	os.Stdout = w
	os.Stdin = bytes.NewReader([]byte(input))

	// Устанавливаем флаги
	os.Args = append([]string{"grep"}, args...)

	// Запускаем main
	main()

	// Восстанавливаем стандартный вывод и ввод
	w.Close()
	os.Stdout = oldStdout
	os.Stdin = oldStdin

	// Читаем результат из pipe
	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

// Тест для проверки основного функционала поиска строки
func TestBasicSearch(t *testing.T) {
	input := "hello world\nfoo bar\nbaz qux\n"
	expected := "foo bar\n"
	result := runGrep(input, "foo")

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Тест для проверки игнорирования регистра (-i)
func TestIgnoreCase(t *testing.T) {
	input := "Hello World\nfoo bar\nBaz qux\n"
	expected := "Hello World\n"
	result := runGrep(input, "-i", "hello")

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Тест для проверки инверсии результата (-v)
func TestInvertMatch(t *testing.T) {
	input := "hello world\nfoo bar\nbaz qux\n"
	expected := "hello world\nbaz qux\n"
	result := runGrep(input, "-v", "foo")

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Тест для проверки вывода строки с номером (-n)
func TestLineNumbers(t *testing.T) {
	input := "hello world\nfoo bar\nbaz qux\n"
	expected := "2:foo bar\n"
	result := runGrep(input, "-n", "foo")

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Тест для проверки точного совпадения (-F)
func TestFixedString(t *testing.T) {
	input := "hello\nhello world\nhello!\n"
	expected := "hello\n"
	result := runGrep(input, "-F", "hello")

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Тест для проверки подсчета строк (-c)
func TestCount(t *testing.T) {
	input := "hello world\nfoo bar\nbaz qux\n"
	expected := "1\n"
	result := runGrep(input, "-c", "foo")

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Тест для проверки вывода строк до и после совпадения (-A, -B)
func TestAfterBeforeContext(t *testing.T) {
	input := "line1\nline2\nmatch\nline4\nline5\n"
	expected := "line2\nmatch\nline4\n"
	result := runGrep(input, "-A", "1", "match")

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}

	expected = "line1\nline2\nmatch\n"
	result = runGrep(input, "-B", "2", "match")

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Тест для проверки вывода контекста (-C)
func TestContext(t *testing.T) {
	input := "line1\nline2\nmatch\nline4\nline5\n"
	expected := "line1\nline2\nmatch\nline4\n"
	result := runGrep(input, "-C", "2", "match")

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}
