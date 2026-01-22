package main

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Error(service string, msg string)
	Log(msg string)
}

type StdoutLogger struct {
	path string // как вариант, можно указать путь
}

func (s StdoutLogger) Log(msg string) {
	// Пишу в stdout
	if _, err := fmt.Fprintln(os.Stdout, msg); err != nil {
		fmt.Println(err)
	}
}

func (s StdoutLogger) Error(service string, msg string) {
	log.Printf("в сервисе: %s. error: %s", service, msg)
}

type MockLogger struct {
	messages []string
	count    int
}

func (m *MockLogger) Log(msg string) {
	// Как пример. Не знаю что обычно тут делают
	m.messages = append(m.messages, msg)
	m.count++
}

func (m *MockLogger) Error(service string, msg string) {
	log.Printf("в сервисе: %s. error: %s", service, msg)
}

func main() {
	var logger Logger
	logger = StdoutLogger{}
	logger.Log("Что-то было")
	logger.Error("сервис", "что-то случилось плохое")

	logger = &MockLogger{}
	logger.Error("сервис", "что-то упало")
	logger.Log("залогировал")
}
