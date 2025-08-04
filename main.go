package main

import "fmt"

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	nomePizzaria := "Pizzaria Go"
	instagram, telefone := "@pizzaria_go", 6399203-4419
	fmt.Println(nomePizzaria, instagram, telefone)
}
