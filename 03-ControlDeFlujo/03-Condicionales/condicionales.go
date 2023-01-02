package main

import "fmt"

func main() {
	/*
		APP para restaurantes
		1- Descuentos de 10% hastas 100 de consumo
		2- Descuentos de 20% de 100 a 200 de consumo
		3- Descuentos de 30% de 200 a 300 de consumo
		4- igv 19%
	*/

	var consumo, descuento float64
	var datosDeDescuentos string

	//Entrada de datos
	fmt.Print("Ingrese el consumo: ")
	fmt.Scan(&consumo)

	//Proceso
	if consumo >= 0 && consumo <= 100 {
		//Descuento de 10%
		datosDeDescuentos = "Descuento de 10%"
		descuento = consumo * 0.1
	} else {
		fmt.Println("No hay descuento")
	}

	//Operaciones
	montoDeDescuento := consumo - descuento
	igv := montoDeDescuento * 0.19
	total := montoDeDescuento + igv

	//Salida de datos
	fmt.Println("fACTURA DE CONSUMO")
	fmt.Println("Datos de descuento: ", datosDeDescuentos)
	fmt.Println("Monto de descuento: ", descuento)
	fmt.Println("Monto de descuento: ", montoDeDescuento)
	fmt.Println("IGV: ", igv)
	fmt.Println("Total: ", total)
}
