package main

import "fmt"

// Reader интерфейс.
type Reader interface {
	Read() string
}

// Closer интерфейс.
type Closer interface {
	Close() error
}

// ReadCloser объединяет Reader и Closer.
type ReadCloser interface {
	Reader
	Closer
}

// File реализует ReadCloser.
type File struct {
	name string
}

func (f File) Read() string {
	return "чтение файла: " + f.name
}

func (f File) Close() error {
	fmt.Println("закрытие файла:", f.name)

	return nil
}

func main() {
	var rc ReadCloser = File{name: "test.txt"}

	fmt.Println(rc.Read())

	err := rc.Close()
	if err != nil {
		return
	}
}
