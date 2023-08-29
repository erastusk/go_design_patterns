package main

import (
	"log"
)

func main() {
	s := ServerOpt{}

	t := WithPort("Port")
	t(&s)
	log.Printf("%+v", s)
	_ = WithOptionTest(8)
	newServerTest(WithOptionTest(9))
}

type ServerOpt struct {
	port  string
	fname string
	lname string
}

type (
	Option      func(*ServerOpt)
	Option_test func(string) string
)

func WithPort(a string) Option {
	return func(s *ServerOpt) {
		s.port = a
	}
}

func WithFname(a string) Option {
	return func(s *ServerOpt) {
		s.fname = a
	}
}

func WithLname(a string) Option {
	return func(s *ServerOpt) {
		s.lname = a
	}
}

func WithOptionTest(int) func(string) string {
	return func(x string) string {
		log.Println("WithOptionTest executed")
		return x
	}
}

func newServerTest(opts ...Option_test) {
	for _, o := range opts {
		o("")
	}
}
