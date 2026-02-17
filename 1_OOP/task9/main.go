package main

import (
	"fmt"
)

func Describe(v []any) {
	for _, value := range v {
		switch t := value.(type) {
		case string:
			fmt.Println("string", t)
		case int:
			fmt.Println("int", t)
		case float64:
			fmt.Println("float64", t)
		}
	}

	//switch t := v.(type) {
	//case string:
	//	fmt.Println("string", t)
	//case int:
	//	fmt.Println("int", t)
	//case float64:
	//	fmt.Println("float64", t)
	//}
}

type Named interface {
	Name() string
}

type User struct {
	name string
}

func (u User) Name() string {
	return u.name
}

func main() {
	name := "aaa"
	num := 12345
	slice := []any{name, num}

	Describe(slice)
}
