package main

import "fmt"

//Struct Persona
type Persona struct {
	Nombre string
	Edad   int
}

//Crear METODOS desde la estructura persona

//Metodo para imprimir los datos de la persona
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.Nombre, p.Edad)
}

// Metodo para modicar el nombre de la persona y no hacerlo de forma explicita
func (p *Persona) modificarNombre(nombre string) {
	p.Nombre = nombre
}

//HERENCIA

type Empleado struct {
	Persona
}

//Crear objetos desde la estructura persona

func main() {
	persona1 := Persona{"Juan", 20}
	//fmt.Println(persona1)
	persona1.imprimir()

	//Si queremos moficar el valor de una propiedad
	//persona1.Nombre = "Pedro"
	persona1.modificarNombre("Rodri")
	persona1.imprimir()
	//fmt.Println(persona1)

	persona2 := Persona{
		Nombre: "Maria",
		Edad:   30,
	}
	//fmt.Println(persona2)
	persona2.imprimir()

	Empleado1 := Empleado{}
	Empleado1.Nombre = "Agustin"
	Empleado1.Edad = 23
	Empleado1.imprimir()
}
