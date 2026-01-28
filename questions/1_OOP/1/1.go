package main

type test struct {
}

func (T test) M() {}

//func (*T test) M() {} // ❌ ошибка компиляции: duplicate method M

func main() {
}
