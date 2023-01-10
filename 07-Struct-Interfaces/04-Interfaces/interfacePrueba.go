package main

import "fmt"

type Pantallas interface {
	pantalla() string
}

type Celular struct{}

type Computadora struct{}

type Tablet struct{}

func (*Celular) pantalla() string {
	return "Soy un celular y tengo una pantalla de 5 pulgadas"
}

func (*Computadora) pantalla() string {
	return "Soy una computadora y tengo una pantalla de 15 pulgadas"
}

func (*Tablet) pantalla() string {
	return "Soy una tablet y tengo una pantalla de 10 pulgadas"
}

func pulgadasPantallas(pantalla Pantallas) {
	fmt.Println(pantalla.pantalla())
}

func main() {
	Celular := Celular{}
	pulgadasPantallas(&Celular)

	Computadora := Computadora{}
	pulgadasPantallas(&Computadora)

	Tablet := Tablet{}
	pulgadasPantallas(&Tablet)
}
