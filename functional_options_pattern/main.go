package main

import (
	"log"
	"time"
)

func main() {
	s := newServerOptions(WithTimeout(time.Second*5), WithPort("6000"))
	log.Printf("%+v", s)
}

// Struct  you want its options defined dynamically
type ServerOptions struct {
	timeout time.Duration
	port    string
	network string
	tls     bool
}

// Func takes in functions of type Option because it returns a func with same signature as Option
// newServerOptions is using a variadic function because of the 3 dots, which means zero or more
func newServerOptions(opts ...Option) *ServerOptions {
	// Declare empty struct, here you could define your defaults incase user doesn't pass them OR
	// Adding them as part of the function newServerOptions(), This way it errors out if missing.
	s := ServerOptions{
		port:    "8080",
		network: "tcp",
		tls:     true,
		timeout: time.Second * 10,
	}
	// Range through all passed option function
	for _, fn := range opts {
		// Option function fn will call the *With option provided, eg WithTimeout func and executes it
		fn(&s)
	}
	return &s
}

// Define a function type that matches the return func of your option function, all option functions return must match Option signature
// This type function modifies your option Struct so it will need to accept its pointer.
type Option func(*ServerOptions)

// Define functions for each option configurable that modify your struct.
// Each of the Option functions act as a “option constructor” and returns a function. This returned func takes a pointer to option struct as an argument, and returns nothing. It then modifies the appropriate struct field using the closure function principle -> A function that returns a function
// WithTimeout becomes of type Option because it modifies *ServerOptions or rather because its return func is the same signature as Option func
func WithTimeout(t time.Duration) func(s *ServerOptions) {
	return func(s *ServerOptions) {
		s.timeout = t
	}
}

func WithPort(t string) func(s *ServerOptions) {
	return func(s *ServerOptions) {
		s.port = t
	}
}
