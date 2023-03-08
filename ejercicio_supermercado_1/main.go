package main

import (
	"ejercicios_go/ejercicio_supermercado_1/service"

	"github.com/gin-gonic/gin"
)

func main(){

	//se define el router
	router := gin.Default()

	gopher := router.Group("/products")

	//trae todos los productos
	gopher.GET("",service.AllProducts)

	//trae productos por id
	gopher.GET("/:id",service.ProductById)

	//trae productos mayores al precio del parametro priceGt
	gopher.GET("/search",service.ProductPriceGt)

	//graba un nuevo producto
	gopher.POST("", service.SaveProduct)


	router.Run(":8080")

	
}