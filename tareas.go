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
	lista := &taskList{
		tasks: []*task{
			t1, t2, t3,
		},
	}
	lista.agregarALista(t4)

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index:", i, "nombre:", lista.tasks[i].nombre)
	}
	fmt.Println("**********************************************")
	for index, tarea := range lista.tasks {
		fmt.Println("Index:", index, "nombre:", tarea.nombre)
	}
}
