package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Определяем флаги
	fieldsFlag := flag.String("f", "", "выбрать поля (колонки)")
	delimiterFlag := flag.String("d", "\t", "использовать другой разделитель")
	separatedFlag := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	// Парсим поля
	fieldIndices, err := parseFields(*fieldsFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при разборе полей:", err)
		os.Exit(1)
	}

	// Вызов основной функции с использованием os.Stdin и os.Stdout
	err = runCut(os.Stdin, os.Stdout, fieldIndices, *delimiterFlag, *separatedFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
}

// runCut выполняет основную логику утилиты cut
func runCut(input io.Reader, output io.Writer, fieldIndices []int, delimiter string, separated bool) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		// Если -s установлен и строка не содержит разделитель, пропускаем её
		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		// Разбиваем строку по разделителю
		columns := strings.Split(line, delimiter)

		// Выводим только выбранные поля
		var selectedFields []string
		for _, idx := range fieldIndices {
			if idx < len(columns) {
				selectedFields = append(selectedFields, columns[idx])
			}
		}
		fmt.Fprintln(output, strings.Join(selectedFields, delimiter))
	}

	return scanner.Err()
}

// parseFields разбирает флаг -f и возвращает индексы колонок
func parseFields(fields string) ([]int, error) {
	var indices []int
	fieldsList := strings.Split(fields, ",")
	for _, field := range fieldsList {
		var index int
		_, err := fmt.Sscanf(field, "%d", &index)
		if err != nil {
			return nil, fmt.Errorf("неверный формат поля: %s", field)
		}
		indices = append(indices, index-1) // переводим в 0-индексацию
	}
	return indices, nil
}
