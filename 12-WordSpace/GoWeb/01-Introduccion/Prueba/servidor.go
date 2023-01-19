package main

import (
	"fmt"
	"net/http"
	"log"
)

func Hola(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola mundo")
}

func Chau(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Chau mundo desde funcion")
}


func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/1", Chau)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}	

	fmt.Println("Servidor corriendo en http://localhost:3000")

	log.Fatal(server.ListenAndServe())
}