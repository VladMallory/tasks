package main

type MyStruct struct {
	a int
	b bool
	f float64
	c complex128
	s struct {
		name string
	}
}

func (s MyStruct) da1() {
	s.a = 10
	s.c = 30
	s.f = 10.10
	s.c = 30
}

func (s *MyStruct) da2() {
	s.a = 20
}

func main() {
}
