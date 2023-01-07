package main

import "fmt"

// ESTRUCTURA
type Task struct {
	Name        string
	Description string
	Completed   bool
}

// METODOS
func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletado: %t\n", t.Name, t.Description, t.Completed)
}

func (t *Task) markComplete() {
	t.Completed = true
}

func main() {
	t1 := Task{
		Name:        "Estudiar Go",
		Description: "Estudiar el curso de Go de Udemy",
		Completed:   true,
	}
	t1.toPrint()

	t2 := Task{
		Name:        "Estudiar Python",
		Description: "Estudiar el curso de Python de YOUTUBE",
		Completed:   false,
	}
	t2.toPrint()
}
