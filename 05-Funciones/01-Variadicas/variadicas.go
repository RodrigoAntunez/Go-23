package main

import "fmt"

//Funciones variadicas, son funciones que pueden recibir un numero indefinidos de parametros
//Para declarar una funcion variadica, se debe colocar el tipo de dato seguido de tres puntos
func sumar(numeros ...int) int {
	var total int
	for _, numero := range numeros {
		total += numero
	}
	return total
}
func main() {
	result := sumar(4, 5, 9)
	fmt.Println(result)
}
