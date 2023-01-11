package main

//Aplicacion que se encarga de ver si un servidor esta funcionando o no

import (
	"fmt"
	"net/http"
	"time"
)

// Funcion que se va a encargar de ver si el servidor esta funcionando
func obversarServidor(servidor string) {
	_, err := http.Get(servidor)

	if err != nil {
		fmt.Println(servidor, "no esta funcionando")
	} else {
		fmt.Println(servidor, "esta funcionando")
	}
}

func main() {
	inicio := time.Now()

	// Arreglo de servidores
	servidores := []string{
		"http://google.com",
		"http://platzi.com",
		"http://facebook.com",
		"http://digitalsport.com.ar",
	}

	for _, servidor := range servidores {
		go obversarServidor(servidor)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion: ", tiempoPaso)
}
