package main

import "fmt"

func sumar(a, b int) {
	fmt.Println("Ingresa el primer numero: ")
	fmt.Scanln(&a)
	fmt.Println("Ingresa el segundo numero: ")
	fmt.Scanln(&b)
	fmt.Println("La suma es: ", a+b)
}

func division(a, b int) {
	fmt.Println("Ingresa el primer numero: ")
	fmt.Scanln(&a)
	fmt.Println("Ingresa el segundo numero: ")
	fmt.Scanln(&b)
	fmt.Println("La division es: ", a/b)
	fmt.Println("El residuo es: ", a%b)
}

func main() {
	//sumar(0, 0)
	division(0, 0)
}
