package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID int `json:"id"`
	// Если это делать с приватными полями, то
	// пакет json не увидет полей структуры
	// из-за привытных полей и не сможет прочитать json
	// name string `json:"name"` будет не правильно
	Name string `json:"name"`
}

func main() {
	b, _ := json.Marshal(User{ID: 1, Name: "Ivan"})
	fmt.Println(string(b))
}
