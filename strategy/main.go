package main

func main() {
	mongo_conn := Mongodb{}
	gate_mongo := NewGateway(mongo_conn)
	gate_mongo.db.conn()

	mysql_conn := Mysqldb{}
	gate_mysql := NewGateway(mysql_conn)
	gate_mysql.db.conn()
	sd := ProdData{}
	s := MiddlewareData{
		next: sd,
	}
	tr := DataProdStruct{
		id:  1,
		lat: 2,
		lon: 3,
	}
	z := MiddlewareProduceData(s, tr)
	z(tr)
	z(tr)
}
