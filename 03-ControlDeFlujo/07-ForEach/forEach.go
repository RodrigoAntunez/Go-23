package main

import "fmt"

func main() {
	nombres := []string{"Juan", "Pedro", "Maria", "Jose"}

	// for i := 0; i < len(nombres); i++ {
	// 	fmt.Println(nombres[i])
	// }

	for indice, elementos := range nombres {
		fmt.Println(indice, elementos)
	}
	//Prubea de git
}
