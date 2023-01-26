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
type User struct {
	Name string
	Age  int
	//Activo bool
	//Admin  bool
	//Cursos []Curso
}

// type Curso struct {
// 	Nombre string
// }

var templates = template.Must(template.ParseGlob("templates/**/*.html"))

// Handler de la pagina principal
// Esta funcion se encarga de renderizar el template index.html
func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo")
	// c1 := Curso{"Go"}
	// c2 := Curso{"Python"}
	// c3 := Curso{"Java"}
	// c4 := Curso{"C++"}

	//FUNCIONES, aca van todas las funciones
	// funciones := template.FuncMap{
	// 	"saludar": Saludar,
	// }

	// MUST, si hay un error en el template, se rompe la aplicacion. Tambien sirve para no trabajar con el error en los handlers

	//template := template.Must(template.ParseFiles("indexIterador.html", "base.html"))

	// cursos := []Curso{c1, c2, c3, c4}
	usuario := User{"Rodri Agustin", 20}

	// if err != nil {
	// 	panic(err)
	// } else {
	err := templates.ExecuteTemplate(rw, "indexIterador.html", usuario)
	// }

	if err != nil {
		panic(err)
	}
}

func Registro(rw http.ResponseWriter, r *http.Request) {

	err := templates.ExecuteTemplate(rw, "registro.html", nil)

	if err != nil {
		panic(err)
	}

}

func main() {
	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	//Servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("Servidor corriendo en http://localhost:3000")

	log.Fatal(server.ListenAndServe())
}
