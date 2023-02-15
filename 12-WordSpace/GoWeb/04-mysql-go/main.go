package main

import (
	"mysqlgo/db"
	"mysqlgo/models"
)

func main() {
	// ...
	db.Connect()
	db.CreateTable(models.UserSchema) //crea la tabla
	db.Ping()
	db.Close()
}
