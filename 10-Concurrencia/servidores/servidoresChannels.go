package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) { //3
	_, err := http.Get(servidor)

	if err != nil {
		//fmt.Println(servidor, "No esta disponible")
		canal <- servidor + " No esta disponible" //4
	} else {
		//fmt.Println(servidor, "Esta funcionando")
		canal <- servidor + " Esta funcionando" //4
	}
}

func main() {
	inicio := time.Now()

	canal := make(chan string) // 2

	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
	}

	for _, sevidor := range servidores {
		go revisarServidor(sevidor, canal) //1
	}

	for i := 0; i < len(servidores); i++ { //5
		fmt.Println(<-canal)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución: ", tiempoPaso)
}
