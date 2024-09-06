package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
 Паттерн "Команда" используется для инкапсуляции запроса как объекта, позволяя параметризовать объекты различными запросами.
 Применимость:
 - Подходит для создания команд, которые можно отменить или повторить.
 - Используется для логирования операций или очередей задач.

 Плюсы:
 + Упрощает добавление новых команд.
 + Позволяет легко реализовать отмену или повтор операции.
 Минусы:
	- Может привести к увеличению числа классов, особенно для каждой новой команды.
Пример использования: системы undo/redo, выполнение задач в очереди, макросы.
*/

// Command определяет интерфейс для выполнения команды
type Command interface {
	Execute()
}

// Light представляет объект, с которым будет работать команда
type Light struct{}

// On включает свет
func (l *Light) On() {
	fmt.Println("Light is ON")
}

// Off выключает свет
func (l *Light) Off() {
	fmt.Println("Light is OFF")
}

// LightOnCommand - команда для включения света
type LightOnCommand struct {
	light *Light
}

// NewLightOnCommand создает новую команду для включения света
func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light}
}

// Execute выполняет команду для включения света
func (c *LightOnCommand) Execute() {
	c.light.On()
}

// LightOffCommand - команда для выключения света
type LightOffCommand struct {
	light *Light
}

// NewLightOffCommand создает новую команду для выключения света
func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light}
}

// Execute выполняет команду для выключения света
func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// RemoteControl представляет собой пульт управления
type RemoteControl struct {
	command Command
}

// SetCommand задает команду для выполнения
func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

// PressButton выполняет заданную команду
func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func command() {
	light := &Light{}
	remote := &RemoteControl{}

	lightOn := NewLightOnCommand(light)
	lightOff := NewLightOffCommand(light)

	// Включаем свет
	remote.SetCommand(lightOn)
	remote.PressButton()

	// Выключаем свет
	remote.SetCommand(lightOff)
	remote.PressButton()
}
