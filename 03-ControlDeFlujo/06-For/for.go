package main

import "fmt"

func main() {
	//Bucle infinito

	// for{
	// 	println("Hola")
	// }

	//Bucle tipo while
	numeros := 12455
	c := 0

	for numeros > 0 {
		numeros /= 10
		c++
	}

	fmt.Println("El numero tiene", c, "digitos")

	//Bucle tipo for

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
