package main

import "fmt"

//Struct Persona
type Persona struct {
	Nombre string
	Edad   int
}

//Crear objetos desde la estructura persona

func main() {
	persona1 := Persona{"Juan", 20}
	fmt.Println(persona1)

	//Si queremos moficar el valor de una propiedad
	persona1.Nombre = "Pedro"

	fmt.Println(persona1)

	persona2 := Persona{
		Nombre: "Maria",
		Edad:   30,
	}
	fmt.Println(persona2)
}
