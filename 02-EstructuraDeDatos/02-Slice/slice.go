package main

import "fmt"

func main() {
	// Slicen, se diferencia de los arrays en que no se especifica el tamaño
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(numeros, len(numeros))

	// para agregar un elemento al final del slice
	numeros = append(numeros, 11)

	fmt.Println(numeros, len(numeros))
}
