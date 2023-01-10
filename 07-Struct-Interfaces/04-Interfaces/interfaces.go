package main

import "fmt"

//Los interface maneja los metodos indenticos de las structura
type Animal interface {
	mover() string
}

type Perro struct{}

type Pez struct{}

type Pajaro struct{}

func (*Perro) mover() string {
	return "Soy perro y camino"
}

func (*Pez) mover() string {
	return "Soy un pez y nado"
}

func (*Pajaro) mover() string {
	return "Soy un pajaro y vuelo"
}

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}
func main() {
	perro := Perro{}
	moverAnimal(&perro)

	pez := Pez{}
	moverAnimal(&pez)

	pajaro := Pajaro{}
	moverAnimal(&pajaro)
}