package main

import "log"

func DecoratorFunc(a string) string {
	log.Println("Decorator Func Call - > ", a)
	return a
}
