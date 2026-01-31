package main

type Logger struct {
	level string
}

type yookassa struct {
	// Встроил другую структуру
	// Теперь она имеет те же что что и Logger
	Logger
	UUID string
}

func main() {
	// Это не наследование, в go его нету
	// Это встраивание

	// Го использует композицию + интерфейсы
	// File содержит Reader
	// Методы Reader доступны у File
	/*
		type Reader struct{}

		func (Reader) Read() {}

		type File struct {
			Reader // встраивание
		}
	*/

	// Полиморфизм через интерфейсы
	/*
		type Reader interface {
		    Read()
		}

		type File struct{}

		func (File) Read() {}

		Теперь
		var r Reader
		r = File{}
	*/

	yookassaClient := yookassa{}
	// Мы смогли обратится к полю из
	// Logger, это встраивание
	yookassaClient.level = "info"
}
