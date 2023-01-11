package main

import (
	"fmt"
	"net/http"
)

func Servidores(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "No esta funcionando")
	} else {
		fmt.Println(servidor, "Esta funcionando")
	}
}

func main() {
	estado := []string{
		"http://google.com",
		"http://platzi.com",
	}

	for _, servidor1 := range estado {
		Servidores(servidor1)
	}
}
