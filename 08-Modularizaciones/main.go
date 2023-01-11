package main

import (
	"fmt"
	"package/mensajes"
	"package/models"
)

func main() {
	mensajes.Hola()
	p1 := models.Persona{}
	p1.Constructor("Juan", 20)
	fmt.Println(p1)

	fmt.Println(p1.GetNombre())
	p1.SetNombre("Pedro")
	fmt.Println(p1)
}
