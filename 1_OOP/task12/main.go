package main

import (
	"fmt"
)

type Server struct {
	Port int
	Host string
	TLS  bool
}

type Option func(*Server)

func NewServer(opts ...func(*Server)) *Server {
	s := &Server{
		Host: "localhost",
		Port: 8080,
		TLS:  false,
	}

	// opt это одна функция
	// мы берем func (s *Server) { s.Host = "google"}
	// и уже работаем с этой одной функцией
	for _, opt := range opts {
		// func
		opt(s)
	}

	return s
}

func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTLS(tls bool) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func main() {
	s1 := NewServer()
	fmt.Println(s1)

	s2 := NewServer(
		WithHost("google"),
		WithPort(443),
		WithTLS(true),
	)

	fmt.Println(s2)
}
