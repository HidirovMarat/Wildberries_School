package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Паттерн "Стратегия" определяет семейство алгоритмов, инкапсулирует их и делает их взаимозаменяемыми.
// Применимость:
// - Подходит, когда нужно выбрать или изменить алгоритм выполнения задачи во время выполнения программы.
// - Используется для разделения логики между разными алгоритмами.
//
// Плюсы:
// + Позволяет легко менять алгоритмы на лету.
// + Упрощает тестирование и поддержку.
// Минусы:
// - Усложняет код за счет необходимости создания дополнительных классов.
//
// Пример использования: системы сортировки, вычислительные задачи, системы скидок.

import "fmt"

// Strategy определяет интерфейс стратегии
type Strategy interface {
	Execute(a, b int) int
}

// AddStrategy реализует стратегию сложения
type AddStrategy struct{}

// Execute выполняет сложение
func (s *AddStrategy) Execute(a, b int) int {
	return a + b
}

// MultiplyStrategy реализует стратегию умножения
type MultiplyStrategy struct{}

// Execute выполняет умножение
func (s *MultiplyStrategy) Execute(a, b int) int {
	return a * b
}

// Context использует стратегию для выполнения операции
type Context struct {
	strategy Strategy
}

// SetStrategy задает текущую стратегию
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// ExecuteStrategy выполняет операцию, используя текущую стратегию
func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {
	context := &Context{}

	// Используем стратегию сложения
	context.SetStrategy(&AddStrategy{})
	fmt.Println("Result:", context.ExecuteStrategy(5, 3)) // Результат: 8

	//
}