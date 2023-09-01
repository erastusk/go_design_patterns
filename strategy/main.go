package main

func main() {
	mongo_conn := Mongodb{}
	gate_mongo := NewGateway(mongo_conn)
	gate_mongo.db.conn()

	mysql_conn := Mysqldb{}
	gate_mysql := NewGateway(mysql_conn)
	gate_mysql.db.conn()
}
