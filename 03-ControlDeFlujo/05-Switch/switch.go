package main

import "fmt"

func main() {
	var vocal string

	fmt.Print("Ingrese una letra: ")
	fmt.Scan(&vocal)

	// switch {
	// case vocal == "a" || vocal == "A":
	// 	fmt.Println("Es una vocal")
	// case vocal == "e" || vocal == "E":
	// 	fmt.Println("Es una vocal")
	// case vocal == "i" || vocal == "I":
	// 	fmt.Println("Es una vocal")
	// case vocal == "o" || vocal == "O":
	// 	fmt.Println("Es una vocal")
	// case vocal == "u" || vocal == "U":
	// 	fmt.Println("Es una vocal")
	// default:
	// 	fmt.Println("No es una vocal")
	// }

	switch vocal {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println("Es una vocal")
	default:
		fmt.Println("No es una vocal")
	}
}
