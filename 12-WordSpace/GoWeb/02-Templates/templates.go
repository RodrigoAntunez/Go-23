package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Estructuras de datos
type User struct {
	Name   string
	Age    int
	Activo bool
	Admin  bool
	Cursos []Curso
}

type Curso struct {
	Nombre string
}

// Handler de la pagina principal
// Esta funcion se encarga de renderizar el template index.html
func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo")
	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"C++"}

	template, err := template.ParseFiles("indexIterador.html")

	cursos := []Curso{c1, c2, c3, c4}
	usuario := User{"Rodri Agustin", 20, true, false, cursos}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, usuario)
	}
}

func main() {
	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	//Servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("Servidor corriendo en http://localhost:3000")

	log.Fatal(server.ListenAndServe())
}
