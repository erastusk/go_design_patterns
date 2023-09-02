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

// Gateway
type Gateway struct {
	db Connection
}

func NewGateway(db Connection) *Gateway {
	return &Gateway{
		db: db,
	}
}

// Middlware concept - GG
type DataProdStruct struct {
	id  int
	lat int
	lon int
}

type ProduceData interface {
	Produce(DataProdStruct)
}

type ProdData struct{}

func (p ProdData) Produce(d DataProdStruct) {
	fmt.Printf("ProData %+v\n", d)
}

type MiddlewareData struct {
	next ProdData
}

func (p MiddlewareData) Produce(d DataProdStruct) {
	fmt.Printf("MiddlewareData %+v\n", d)
}

func MiddlewareProduceData(pr MiddlewareData, t DataProdStruct) func(DataProdStruct) {
	return func(DataProdStruct) {
		fmt.Printf("Middleware happening here\n")
		pr.next.Produce(t)
	}
}
