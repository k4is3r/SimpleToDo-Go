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

func main() {
	t := &task{
		nombre:      "Completar mi curso de DeepLearning",
		descripcion: "Completar mi cursos de DeepLearning en semana de cuarentena",
	}
	fmt.Printf("%+v\n", t)
	t.marcarCompleta()
	t.actualizarNombre("finalizar curso DeepLearning")
	t.actualizarDescripcion("Completar curso lo antes posible")
	fmt.Printf("%+v\n", t)
}
