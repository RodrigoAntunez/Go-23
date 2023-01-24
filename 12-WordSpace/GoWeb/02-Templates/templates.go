package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//FUNCIONES
func Saludar(name string) string {
	return "Hola desde una funcion " + name
}

// Estructuras de datos
/*type User struct {
	Name   string
	Age    int
	Activo bool
	Admin  bool
	Cursos []Curso
}

type Curso struct {
	Nombre string
}
*/

// Handler de la pagina principal
// Esta funcion se encarga de renderizar el template index.html
func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo")
	// c1 := Curso{"Go"}
	// c2 := Curso{"Python"}
	// c3 := Curso{"Java"}
	// c4 := Curso{"C++"}

	//FUNCIONES, aca van todas las funciones
	funciones := template.FuncMap{
		"saludar": Saludar,
	}

	template, err := template.New("indexIterador.html").Funcs(funciones).
		ParseFiles("indexIterador.html")

	// cursos := []Curso{c1, c2, c3, c4}
	// usuario := User{"Rodri Agustin", 20, true, false, cursos}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, nil)
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
