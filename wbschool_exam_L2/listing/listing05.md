Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
...

```
Программа выведет:

error

Интерфейсное значение err:

    В Go интерфейсное значение состоит из двух частей: динамического типа и динамического значения.
    В данном случае типом интерфейса err будет *customError, а его значением — nil.
    Важно отметить, что интерфейс считается равным nil только тогда, когда и динамический тип, и динамическое значение равны nil.
    В данном случае тип интерфейса — *customError, а значение — nil, поэтому выражение if err != nil вернёт true