package main

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

func main() {

}

func unboxLine(line []rune) ([]rune, bool) {
	res := make([]rune, 0)
	last := ' '

	for i := 0; i < len(line); i++ {
		current := line[i]
		if last == ' ' && isNumber(current) {
			return nil, false
		}
		if last == ' ' && isLetter(current) {
			last = current
			continue
		}
		if last == ' ' && current == '/' {
			last = '/'
		}
	}

	return res
}

func printRuneKOnce(k rune, numsCount int) []rune {
	numsK := make([]rune, numsCount)
	for i := 0; i < numsCount; i++ {
		numsK = append(numsK, k)
	}

	return numsK
}

func isNumber(r rune) bool {
	return byte(r) >= 48 && byte(r) <= 57
}

func isLetter(r rune) bool {
	return byte(r) >= 97 && byte(r) <= 122
}
