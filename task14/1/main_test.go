package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

// TestLoggerInterface проверяет, существует ли интерфейс.
func TestLoggerInterface(t *testing.T) {
	var logger Logger
	// если бы интерфейса не было
	// была бы ошибка
	_ = logger
}

func TestLoggerCreation(t *testing.T) {
	stdoutLogger := &StdoutLogger{}
	mockLogger := &MockLogger{}

	// Проверка, создались ли они без ошибки
	// nil значит что что-то не так, ведь они
	// должны быть заполены &{} &{[] 0}
	if stdoutLogger == nil || mockLogger == nil {
		t.Errorf("ошибка создания структур")
	}
}

// TestStdoutLogger_Log вообще не понимаю, мы перехватываем и с чем-то сравниваем.
func TestStdoutLogger_Log(t *testing.T) {
	logger := &StdoutLogger{}

	// Перехватываем stdout
	old := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal("os.Pipe:", err)
	}

	os.Stdout = w

	logger.Log("test")

	if err := w.Close(); err != nil {
		log.Fatal("ошибка закрытия ")
	}

	os.Stdout = old

	var buf bytes.Buffer

	_, err = io.Copy(&buf, r)
	if err != nil {
		log.Fatal("")
	}

	if !strings.Contains(buf.String(), "test") {
		t.Errorf("expected 'test' in output, got: %s", buf.String())
	}
}
