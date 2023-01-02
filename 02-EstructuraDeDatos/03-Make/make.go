package main

import "fmt"

func main() {
	// make([]T, length, capacity)
	numeros := make([]int, 0, 3)

	// Para agregar elementos al make se los agrega con append

	numeros = append(numeros, 1)
	fmt.Println(numeros, len(numeros), cap(numeros))

	numeros = append(numeros, 2)
	fmt.Println(numeros, len(numeros), cap(numeros))

	// Aca hay que modificar la longitud del slice, porque no se puede modificar la capacidad
	//numeros2 := make([]int, 0, 3)
	numeros2 := make([]int, 3, 3)

	numeros2[0] = 1
	fmt.Println(numeros2, len(numeros2), cap(numeros2))
}
