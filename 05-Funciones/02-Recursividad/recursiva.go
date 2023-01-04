package main

import "fmt"

//Funciones recursivas, son funciones que se llaman a si mismas
//Para declarar una funcion recursiva, se debe colocar el tipo de dato seguido de tres puntos

func factorial(numero int) int {
	if numero == 0 {
		return 1
	}
	f := numero * factorial(numero-1) //Aca se llama a si misma

	return f
}

func main() {
	result := factorial(3)
	fmt.Println(result)
}
