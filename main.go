package main

import (
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)

	//toscana, Marguerita, atum com queijo
	//var pizzas = []Pizza{
	//	{ID: 1, nome: "Toscana", preco: 49.50},
	//	{ID: 2, nome: "Marguerita", preco: 79.50},
	//	{ID: 3, nome: "Atum com queijo", preco: 69.50},
	//	{ID: 4, nome: "4Queijos", preco: 80.00},
	//	{ID: 5, nome: "Calabresa", preco: 90.00},
	//	{ID: 6, nome: "Frango", preco: 78.00},
	//	{ID: 7, nome: "Carne de Sol", preco: 110.00},
	//	{ID: 8, nome: "3Queijos", preco: 69.00},
	//}
	//fmt.Println(pizzas)
	router.Run()


}

func getPizzas(c *gin.Context) {
	
}

func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pizzas": "Toscana, Calebresa",
		})
	})
