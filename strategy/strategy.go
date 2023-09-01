package main

import "fmt"

type Connection interface {
	conn()
}

// Mongodb
type Mongodb struct{}

func (m Mongodb) conn() {
	fmt.Println("Mongodb connection")
}

// Mysqldb
type Mysqldb struct{}

func (m Mysqldb) conn() {
	fmt.Println("Mysql connection")
}

type Gateway struct {
	db Connection
}

func NewGateway(db Connection) *Gateway {
	return &Gateway{
		db: db,
	}
}
