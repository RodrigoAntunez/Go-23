package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//usarname:password@tcp(localhost:3306)/datbase
//user:123456@tcp(127.0.0.1:3306)/goweb_db
const url = "root:Sql2022@tcp(localhost:3306)/goweb_db" //url de conexion a la base de datos
// Guarda la conexion
var db *sql.DB

// Realiza la conexion
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa a la base de datos")
	db = conection
}

// Cierra la conexion
func Close() {
	db.Close()
}

// Verifica si la conexion esta activa
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Conexion activa")
}

//crea la tabla

func CreateTable(schema string) {
	db.Exec(schema) //ejecuta el esquema de la tabla
}
