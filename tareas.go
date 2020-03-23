package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	ceompletado bool
}

func main() {
	t := task{
		nombre:      "Completar mi curso de DeepLearning",
		descripcion: "Completar mi cursos de DeepLearning en semana de cuarentena",
	}
	fmt.Printf("%+v\n", t)
}
