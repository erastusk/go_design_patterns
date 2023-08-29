package main

import "log"

type DB interface {
	Store(a string) string
	StoringStuff()
}
type Decorator_D struct{}

func newDecorator_D_() *Decorator_D {
	return &Decorator_D{}
}

func (d *Decorator_D) Store(a string) string {
	log.Println("Calling DB store Ex1 -> \n ", a)
	return a
}

func (d *Decorator_D) StoringStuff() {
	log.Println("Calling DB storingstuff")
}
