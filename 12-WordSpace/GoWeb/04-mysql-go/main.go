package main

import "mysqlgo/db"

func main() {
	// ...
	db.Connect()
	db.Ping()
	db.Close()
}
