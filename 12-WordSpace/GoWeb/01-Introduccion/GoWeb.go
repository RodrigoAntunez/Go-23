package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Crear un servidor web
	//http.ListenAndServe("localhost:8080", nil)
	fmt.Println("El Servidor web esta en el puerto http://localhost:8080")
	fmt.Println("Run server: http://localhost:8080/")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
