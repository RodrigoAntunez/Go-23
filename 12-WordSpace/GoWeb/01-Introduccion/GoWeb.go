package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler:
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo HTTP es: ", r.Method) // Para saber si es un get o un post
	fmt.Fprintln(rw, "Hola Mundo desde un Handler")
}

func Chau(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Chau Mundo desde un Handler")
}

func PaginaNTF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Error interno del servidor", http.StatusNotFound)
} //Este es un paquete de Go

func main() {

	//Router
	http.HandleFunc("/", Hola)
	http.HandleFunc("/chau", Chau)
	http.HandleFunc("/ntf", PaginaNTF)

	//Crear un servidor web
	//http.ListenAndServe("localhost:8080", nil)
	fmt.Println("El Servidor web esta en el puerto http://localhost:3000")
	fmt.Println("Run server: http://localhost:3000/")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))

	// Router y Handler
	//Router: Es el encargado de recibir las peticiones y enviarlas a un Handler
	//Handler: Es el encargado de procesar la petición y devolver una respuesta(son funciones que responden a una ruta)

}
