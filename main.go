package main

import (
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	Nome  string
	Preco float64
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()

}

func getPizzas(c *gin.Context) {
	var pizzas = []Pizza{
		{ID: 1, Nome: "Toscana", Preco: 49.50},
		{ID: 2, Nome: "Marguerita", Preco: 79.50},
		{ID: 3, Nome: "Atum com queijo", Preco: 69.50},
		{ID: 4, Nome: "4Queijos", Preco: 80.00},
		{ID: 5, Nome: "Calabresa", Preco: 90.00},
		{ID: 6, Nome: "Frango", Preco: 78.00},
		{ID: 7, Nome: "Carne de Sol", Preco: 110.00},
		{ID: 8, Nome: "3Queijos", Preco: 69.00},
		{ID: 9, Nome: "Italiana", Preco: 110.00},
		{ID: 10, Nome: "Chocolate", Preco: 170.00},
	}
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
