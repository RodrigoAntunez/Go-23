package models

type Persona struct {
	nombre string
	edad   int
}

// El constructor es un método que se usa para inicializar los atributos de la estructura

func (p *Persona) Constructor(nombre string, edad int) {
	p.nombre = nombre
	p.edad = edad
}

// Estos métodos no son necesarios, pero se pueden usar para acceder a los atributos de la estructura
func (p *Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) Getead() int {
	return p.edad
}

// Modificadores de acceso para los atributos de la estructura
func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona) SetEdad(edad int) {
	p.edad = edad
}
