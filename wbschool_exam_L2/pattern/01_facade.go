package pattern

/*
Паттерн «Фасад»
Описание: Фасад (Facade) — структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотек или фреймворков.

Применимость:

Когда нужно упростить взаимодействие с сложной подсистемой.
Когда требуется уменьшить количество зависимостей между клиентами и сложной системой.
Плюсы:

Упрощает работу с подсистемой, скрывая сложные детали реализации.
Уменьшает количество зависимостей между клиентом и сложной системой.
Минусы:

Может стать "божественным объектом" (God Object), если не ограничивать его ответственность.
Пример использования:

В веб-приложении можно создать фасад для работы с несколькими сервисами: аутентификацией, логированием и управлением базой данных.
go

*/

import "fmt"

// Подсистема 1
type AuthService struct{}

func (a *AuthService) Login(user string) {
	fmt.Println("Logging in user:", user)
}

// Подсистема 2
type LoggerService struct{}

func (l *LoggerService) Log(message string) {
	fmt.Println("Log message:", message)
}

// Подсистема 3
type DatabaseService struct{}

func (d *DatabaseService) Query(query string) {
	fmt.Println("Executing query:", query)
}

// Фасад
type Facade struct {
	authService     *AuthService
	loggerService   *LoggerService
	databaseService *DatabaseService
}

func NewFacade() *Facade {
	return &Facade{
		authService:     &AuthService{},
		loggerService:   &LoggerService{},
		databaseService: &DatabaseService{},
	}
}

func (f *Facade) LoginAndLog(user, logMessage string) {
	f.authService.Login(user)
	f.loggerService.Log(logMessage)
}

func (f *Facade) ExecuteQuery(query string) {
	f.databaseService.Query(query)
}

func facade() {
	facade := NewFacade()
	facade.LoginAndLog("user1", "User logged in")
	facade.ExecuteQuery("SELECT * FROM users")
}