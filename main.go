package main

import "fmt"

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	//toscana, Marguerita, atum com queijo
	var pizzas = []Pizza{
		{ID: 1, nome: "Toscana", preco: 49.50},
		{ID: 2, nome: "Marguerita", preco: 79.50},
		{ID: 3, nome: "Atum com queijo", preco: 69.50},
		{ID: 4, nome: "4Queijos", preco: 80.00},
	}
	fmt.Println(pizzas)
}
