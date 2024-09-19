package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunCut(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		fields    string
		delimiter string
		separated bool
		expected  string
	}{
		{
			name:      "Default delimiter with specific fields",
			input:     "one\ttwo\tthree\nfour\tfive\tsix\n",
			fields:    "1,3",
			delimiter: "\t",
			separated: false,
			expected:  "one\tthree\nfour\tsix\n",
		},
		{
			name:      "Custom delimiter with specific fields",
			input:     "one|two|three\nfour|five|six\n",
			fields:    "1,3",
			delimiter: "|",
			separated: false,
			expected:  "one|three\nfour|six\n",
		},
		{
			name:      "Separated flag with no delimiter",
			input:     "one|two|three\nfour|five|six\nnofields\n",
			fields:    "1,2",
			delimiter: "|",
			separated: true,
			expected:  "one|two\nfour|five\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Имитация ввода с помощью строки
			input := strings.NewReader(tt.input)
			// Буфер для вывода
			output := &bytes.Buffer{}

			// Парсим флаг -f
			fieldIndices, err := parseFields(tt.fields)
			if err != nil {
				t.Fatalf("Ошибка при разборе полей: %v", err)
			}

			// Вызываем функцию runCut
			err = runCut(input, output, fieldIndices, tt.delimiter, tt.separated)
			if err != nil {
				t.Fatalf("Ошибка выполнения runCut: %v", err)
			}

			// Сравниваем полученный вывод с ожидаемым
			if output.String() != tt.expected {
				t.Errorf("expected:\n%s\ngot:\n%s", tt.expected, output.String())
			}
		})
	}
}
