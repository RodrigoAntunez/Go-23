package main

import (
	"fmt"
	"strings"
)

// Funciones anonimas, son funciones que no tienen nombre
// Las funciones anonimas se utilizan para pasar como parametro a otra funcion

//Closure
//Es una funcion que se ejecuta en el momento y que se puede ejecutar varias veces

func repeat(numero int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, numero)
	}
}

func main() {
	/*Funcion anonima
	 1- *
	func() {
		fmt.Println("Hola desde una funcion anonima")
	}()

	2- *
	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola %s desde una funcion anonima", nombre)
	}
	fmt.Println(myfunc("Rodri"))
	*/
	repeat2 := repeat(5)
	fmt.Println(repeat2("Hola \n"))

}
