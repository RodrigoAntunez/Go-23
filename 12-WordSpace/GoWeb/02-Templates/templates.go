package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Handler de la pagina principal
// Esta funcion se encarga de renderizar el template index.html
func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo")
	template, err := template.ParseFiles("index.html")
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
