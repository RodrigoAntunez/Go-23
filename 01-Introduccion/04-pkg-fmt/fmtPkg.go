package main

import "fmt"

func main() {

	Hola := "Hola"
	Mundo := "Mundo"

	//Print imprime el valor de la viriable pero sin salto de linea
	fmt.Print(Hola)
	fmt.Print(Mundo)

	//Println imprime el valor de la viriable con salto de linea
	fmt.Println(Hola, Mundo)

	nombre := "Rodri"
	edad := 24

	//Printf imprime el valor de la viriable con formato
	fmt.Printf("Hola %s tienes %d años", nombre, edad)

	//Se pone V por si no se conoce el valor de la variable a imprimir
	fmt.Printf("Hola %v tienes %v años", nombre, edad)

	//Para guardar el valor que se va a imprimir en una variable
	mensaje := fmt.Sprintf("Hola %s tienes %d años", nombre, edad)
	fmt.Println(mensaje)

	//Para imprimir el tipo de dato de una variable
	fmt.Printf("Nombre: %T", nombre)

	//Para ingresar datos por teclado
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("Hola", nombre)
}
