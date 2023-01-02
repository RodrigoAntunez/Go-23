package main

import "fmt"

func main() {
	// Esta es la forma de declarar un mapa
	dias := make(map[int]string)

	//Agregar elementos al mapa
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"
	dias[7] = "Domingo"
	fmt.Println(dias)

	//Modificar elementos del mapa
	dias[7] = "DOMINGO"
	fmt.Println(dias)

	//Borrar elementos del mapa
	delete(dias, 7)
	fmt.Println(dias)
}
