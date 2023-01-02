package main

import "fmt"

func main() {
	// Esta es la forma de declarar un array
	var numeros [5]int
	numeros[0] = 1

	// Esta es la forma de declarar un array de tipo literal y asignarle valores
	arrays := [3]string{"Hola", "Hola", "Hola"}

	// Esta es la forma de declarar un array de tipo literal sin longitud y asignarle valores
	colores := [...]string{
		"Rojo",
		"Verde",
		"Azul",
	}

	fmt.Println(numeros)
	fmt.Println(arrays)
	fmt.Println(colores, len(colores)) // len() es una función que devuelve la longitud del array

}
