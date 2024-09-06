package main

import (
	"fmt"  // пакет для форматированного ввода/вывода
	"math" // пакет для математических функций
)

// Point - структура для представления точки в декартовой системе координат
type Point struct {
	x float64 // Координата X
	y float64 // Координата Y
}

// NewPoint - функция-конструктор для создания новой точки
func NewPoint(x float64, y float64) *Point {
	// Создание нового экземпляра структуры Point
	point := Point{
		x: x, // Инициализация поля x значением x
		y: y, // Инициализация поля y значением y
	}
	// Возврат адреса созданной точки с помощью указателя
	return &point
}

// (p *Point) Distance(p2 *Point) float64 - метод структуры Point
func (p *Point) Distance(p2 *Point) float64 {
	// Расчет расстояния между двумя точками по формуле Евклида
	return math.Sqrt(math.Pow((p.x-p2.x), 2) + math.Pow((p.y-p2.y), 2))
	//  - math.Pow - возведение в степень
	//  - math.Sqrt - квадратный корень
}

func main() {
	// Создание точек с помощью функции-конструктора NewPoint
	p1 := NewPoint(5, 6)
	p2 := NewPoint(2, 10)

	// Вычисление расстояния между точками p1 и p2 с помощью метода Distance
	distance := p1.Distance(p2)

	// Печать расстояния
	fmt.Print(distance)
}
