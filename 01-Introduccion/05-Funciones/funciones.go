package main

import "fmt"

func saludar() {
	fmt.Println("Hola mundo")
}

func saludar2(nombre string) {
	fmt.Println("Hola", nombre)
}

func sumar(a, b int) int {
	return a + b
}

func main() {
	saludar()
	saludar2("Juan")
	resultado := sumar(2, 3)
	fmt.Println(resultado)
}
