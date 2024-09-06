package main

import (
	"fmt"
	"math/big"
)

type BigCalc struct{}

// Для операций сложения, вычитания, умножения и деления чисел больше 2^20,
// используем пакет math/big, который позволяет оперировать большими числами.
func (b BigCalc) Add(x, y *big.Float) *big.Float {
	// Метод Add типа big.Float складывает 2 big.Float числа.
	return new(big.Float).Add(x, y)
}

func (b BigCalc) Sub(x, y *big.Float) *big.Float {
	// Метод Sub типа big.Float вычитает 2 big.Float числа.
	return new(big.Float).Sub(x, y)
}

func (b BigCalc) Div(x, y *big.Float) *big.Float {
	// Метод Div типа big.Float делит 2 big.Float числа.
	return new(big.Float).Quo(x, y)
}

func (b BigCalc) Multiply(x, y *big.Float) *big.Float {
	// Метод Multiply типа big.Float умножает 2 big.Float числа.
	return new(big.Float).Mul(x, y)
}

func main() {
	var bigCalc BigCalc
	a, _ := new(big.Float).SetPrec(512).SetString("184467440737095516168")
	b, _ := new(big.Float).SetPrec(512).SetString("184467440737095516132")
	fmt.Println(bigCalc.Add(a, b))      // 3.689348814741910323e+20
	fmt.Println(bigCalc.Sub(a, b))      // 36
	fmt.Println(bigCalc.Multiply(a, b)) // 3.4028236692093846342648111928434910822176e+40
	fmt.Println(bigCalc.Div(a, b))      //1.0000000000000000001951563910473907981509639617973111532619491283479808243883991762849533550365238031373365706998554299891143573868155897886025635356895757 вот здесь скорее всего не точно
}
