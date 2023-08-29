package main

import "log"

type Decorator_Struct struct{}

func newDecorator_Struct() *Decorator_Struct {
	return &Decorator_Struct{}
}

func (d *Decorator_Struct) WhichFunc(a string) string {
	log.Println("Calling Decorator_Struct -> ", a)
	return a
}
