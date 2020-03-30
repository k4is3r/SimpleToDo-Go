package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

type task struct {
	nombre      string
	descripcion string
	ceompletado bool
}

func (t *task) marcarCompleta() {
	t.ceompletado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (t *taskList) elminiarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)

}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripcion:", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletada() {
	for _, tarea := range t.tasks {
		if tarea.ceompletado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripcion:", tarea.descripcion)
		}
	}
}

func main() {
	t1 := &task{
		nombre:      "Completar mi curso de DeepLearning",
		descripcion: "Completar mi cursos de DeepLearning en semana de cuarentena",
	}
	t2 := &task{
		nombre:      "Completar mi curso de EthicalHacking",
		descripcion: "Completar mi cursos de EthicalHacking en semana de cuarentena",
	}
	t3 := &task{
		nombre:      "Completar mi curso de Raspberry",
		descripcion: "Completar mi cursos de Raspberry en semana de cuarentena",
	}
	t4 := &task{
		nombre:      "Completar mi curso de GCP",
		descripcion: "Completar mi curso de GCP semana de cuarentena",
	}
	t5 := &task{
		nombre:      "Completar mi curso de Java",
		descripcion: "Completar mi cursos de Java en semana de cuarentena",
	}
	t6 := &task{
		nombre:      "Completar mi curso de C#",
		descripcion: "Completar mi curso de C# semana de cuarentena",
	}
	lista := &taskList{
		tasks: []*task{
			t1, t2, t3,
		},
	}
	lista2 := &taskList{
		tasks: []*task{
			t4, t5, t6,
		},
	}

	mapaTareas := make(map[string]*taskList)
	mapaTareas["edward"] = lista
	mapaTareas["jose"] = lista2

	fmt.Println("Tareas Edward")
	mapaTareas["edward"].imprimirLista()
	fmt.Println("Tareas jose")
	mapaTareas["jose"].imprimirLista()
}
