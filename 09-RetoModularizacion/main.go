package main

import (
	"figuras/cuadrado"
)

func main() {

	// cua1 := Cuadrado{lado: 4}
	// cir1 := Circulo{radio: 5}

	cua1 := cuadrado.Cuadrado{Lado: 4}
	//cir1 := circulo.Circulo{Radio: 5}

	Cuadrado.Medidas(&cua1)

}
