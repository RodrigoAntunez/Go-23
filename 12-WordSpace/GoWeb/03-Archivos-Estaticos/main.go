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
var errorTemplate = template.Must(template.ParseFiles("templates/Error/error.html"))

func handlerErro(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

func RenderTemplate(rw http.ResponseWriter, name string, data interface{}) {

	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		//http.Error(rw, "No es posible retornar el template", http.StatusInternalServerError)
		handlerErro(rw, http.StatusInternalServerError)
	}

}

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
	// }

	RenderTemplate(rw, "indexIterador.html", usuario)
}

func Registro(rw http.ResponseWriter, r *http.Request) {

	RenderTemplate(rw, "registro.html", nil)

}

func main() {
	//ARCHIVOS ESTATICOS
	staticFile := http.FileServer(http.Dir("statics")) //Ruta de los archivos estaticos

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	//Mux de archivos estaticos
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile)) //La ruta de donde va a cargar los handlers

	//Servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("Servidor corriendo en http://localhost:3000")

	log.Fatal(server.ListenAndServe())
}
