package main

import "figuras/figuras"

func main() {

	// cua1 := Cuadrado{lado: 4}
	// cir1 := Circulo{radio: 5}
	cua1 := figuras.Cuadrado{Lado: 4}
	figuras.Medidas(&cua1)

}
