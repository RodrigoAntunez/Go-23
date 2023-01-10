package main

import "fmt"

// ESTRUCTURA

//RELACION DE UNO A UNO
type User struct {
	Name   string
	Email  string
	Status bool
}

type Student struct {
	user   User // Relación de uno a uno.
	Codigo string
}

//RELACION DE UNO A MUCHOS
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func (s *Student) imprimir() {
	fmt.Println(s.user.Name)
	fmt.Println(s.Codigo)
}

func main() {
	//RELACION DE UNO A UNO

	/*alex := User{
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
	*/

	//RELACION DE UNO A MUCHOS

	video1 := Video{titulo: "01-Introduccion"}
	video2 := Video{titulo: "02-Instalacion"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	fmt.Println(video1.curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
