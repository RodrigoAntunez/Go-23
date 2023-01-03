package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int

	for _, valor := range valores {
		num, error := strconv.Atoi(valor)

		if error != nil {
			fmt.Println("Error al convertir el valor")
		} else {
			result += num
		}
	}
	return result
}

func main() {
	var expresion string
	var resultado int

	fmt.Print("=>")
	fmt.Scanln(&expresion)

	resultado = sumar(expresion)
	fmt.Printf("=> %d \n", resultado)
}
