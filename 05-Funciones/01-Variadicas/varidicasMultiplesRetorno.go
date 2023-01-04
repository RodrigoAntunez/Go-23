package main

import "fmt"

// Funciones variadicas, son funciones que pueden recibir un numero indefinidos de parametros
// Esta es la forma de retornar multiples valores

func sumar(nombre string, numeros ...int) (string, int) {
	mensaje := fmt.Sprintf("Hola %s, la suma de los numeros es: ", nombre)
	var total int
	for _, numero := range numeros {
		total += numero
	}
	return mensaje, total
}

func main() {
	mensaje, total := sumar("Juan", 4, 5, 9)
	fmt.Println(mensaje, total)
}
