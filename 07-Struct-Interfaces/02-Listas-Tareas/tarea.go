package main

import "fmt"

// ESTRUCTURA

//Lista de tarea
type TaskList struct {
	Tasks []*Task
}

//Tarea
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

func (tl *TaskList) addTask(t *Task) {
	tl.Tasks = append(tl.Tasks, t)
}

func (tl *TaskList) deleteTask(index int) {
	tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
}

func main() {
	t1 := Task{
		Name:        "Estudiar Go",
		Description: "Estudiar el curso de Go de Udemy",
		Completed:   true,
	}
	//t1.toPrint()

	t2 := Task{
		Name:        "Estudiar Python",
		Description: "Estudiar el curso de Python de YOUTUBE",
		Completed:   false,
	}
	//t2.toPrint()

	lista := TaskList{}
	lista.addTask(&t1)
	lista.addTask(&t2)

	fmt.Println(lista)

	t3 := Task{
		Name:        "Estudiar CSS",
		Description: "Estudiar el curso de Python de YOUTUBE",
		Completed:   false,
	}

	lista.addTask(&t3)

	lista.deleteTask(2)

	for i, task := range lista.Tasks {
		fmt.Println(i, task)
	}
}
