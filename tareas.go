package main

type task struct {
	nombre      string
	descripcion string
	ceompletado bool
}

func main() {
	t := task{
		nombre: "Completar mi curso de DeepLearning"
	}
}