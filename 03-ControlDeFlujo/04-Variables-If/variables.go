package main

import "fmt"

func main() {
	// Declaración de variables en el if
	// El scope de las variables en el if es solo dentro del if
	if nombre, edad := "Alex", 26; nombre == "Alex" {
		fmt.Println("Hola ", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}

	//Obtener valor de mapa
	users := make(map[string]string)

	users["Alex"] = "alex@gmail.com"
	users["Juan"] = "juan@gmail.com"

	if email, ok := users["Alex"]; ok {
		fmt.Println("El email de Alex es: ", email)
	} else {
		fmt.Println("No se encontró el usuario")
	}
}
