package main

import (
	"log"
	"time"
)

func main() {
	// Func
	t := Decorator_func(DecoratorFunc)
	_ = t("1st Decorator Func")
	_ = t("2nd Decorator Func")

	// Struct
	s := newDecorator_Struct()
	s_ := Decorator_struct(s.WhichFunc)
	_ = s_("1st Decorator Struct")
	_ = s_("2nd Decorator Struct")

	// Interface
	a := newDecorator_Int()
	b_ := Decorator_int(a.WhichFuncEx1)
	_ = b_("1st Decorator Int")
	c_ := Decorator_int(a.WhichFuncEx2)
	_ = c_("2nd Decorator Int")

	// Interface instance
	ab := newDecorator_D_()
	zb := Decorator_Any(ab)
	zb("Storex1")
	zb("Storex2")

	// Middleware
	abc := newDecorator_Int()
	bc := DecoratorMiddlewareLogger(abc.WhichFuncEx1)
	_ = bc("1st Decorator Middleware")
	cc := DecoratorMiddlewareLogger(a.WhichFuncEx2)
	_ = cc("2nd Decorator Middleware")
}

func Decorator_func(p func(a string) string) func(string) string {
	return func(a string) string {
		log.Println("Calling Decorator_Func")
		return p(a)
	}
}

func Decorator_struct(p func(a string) string) func(string) string {
	return func(a string) string {
		log.Println("Calling Decorator_struct")
		return p(a)
	}
}

func Decorator_int(p func(a string) string) func(string) string {
	return func(a string) string {
		log.Println("Calling Decorator_int")
		return p(a)
	}
}

func Decorator_Any(db DB) func(string) {
	return func(s string) {
		log.Println("Executing wrapper fun Decorator Decorator_Any", s)
		db.Store(s)
		db.StoringStuff()
	}
}

func DecoratorMiddlewareLogger(p func(a string) string) func(string) string {
	return func(a string) string {
		start := time.Now()
		log.Println("Beginning logging")
		log.Println("Ending logging")
		defer log.Println("Took: ", time.Since(start))
		return p(a)
	}
}
