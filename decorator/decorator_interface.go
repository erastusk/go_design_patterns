package main

import "log"

type Decorator_Int interface {
	WhichFunc(a string) string
}
type Decorator_Int_ struct{}

func newDecorator_Int() *Decorator_Int_ {
	return &Decorator_Int_{}
}

func (d *Decorator_Int_) WhichFuncEx1(a string) string {
	log.Println("Calling Decorator_Int Ex1 -> ", a)
	return a
}

func (d *Decorator_Int_) WhichFuncEx2(a string) string {
	log.Println("Calling Decorator_Int Ex2 -> ", a)
	return a
}
