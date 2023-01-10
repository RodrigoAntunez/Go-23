package main

import "fmt"

// ESTRUCTURA
type User struct {
	Name   string
	Email  string
	Status bool
}

type Student struct {
	user   User // Relación de uno a uno.
	Codigo string
}

func (s *Student) imprimir() {
	fmt.Println(s.user.Name)
	fmt.Println(s.Codigo)
}

func main() {
	alex := User{
		Name:   "Alex",
		Email:  "alex@gmail.com",
		Status: true,
	}

	rodri := User{
		Name:   "Rodri",
		Email:  "rodri@gmai.com",
		Status: true,
	}

	alexStudent := Student{ //Este es una relacion entre dos estructuras. usuario y estudiante.
		user:   alex,
		Codigo: "123456",
	}

	fmt.Println(rodri)
	alexStudent.imprimir()
}
