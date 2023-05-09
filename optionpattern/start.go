/*
Functions accept optional arguments
Go: Struct with many fields. So no need to define many constructors,
instead 1 constructor with any variation of arguments.

https://dsysd-dev.medium.com/writing-better-code-in-go-with-the-option-pattern-bb9283131407
*/
package optionpattern

import (
	"fmt"
	"time"
)

type Server struct {
	host string
	port int
	timeout time.Duration
}

type ServerOption func(*Server)

func withHost(host string) ServerOption {
	return func(s *Server) {
		s.host = host
	}
}

func withPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func withTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func NewServer(opts ...ServerOption) *Server {
	server := &Server{}
	
	for _, opt := range opts {
		opt(server)
	}

	return server
}

func DoWork() {
	server := NewServer(withHost("localhost"), withPort(8000), withTimeout(5 * time.Second))
	fmt.Println(server)
}