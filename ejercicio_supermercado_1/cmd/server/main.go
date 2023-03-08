package main

import (
	"ejercicios_go/ejercicio_supermercado_1/cmd/server/handler"
	"ejercicios_go/ejercicio_supermercado_1/internal/product"
	"ejercicios_go/ejercicio_supermercado_1/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){

	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	jsonStorage := store.NewJsonStore("data.json")
	productRepopository := product.NewRepository(jsonStorage)
	productService := product.NewService(productRepopository)
	productHandler := handler.NewService(productService)

	productRoutes := router.Group("products")
	productRoutes.GET("",productHandler.GetAll())
	productRoutes.GET("/:id",productHandler.GetProductById())
	productRoutes.GET("/search",productHandler.GetProductWithPriceGt())
	productRoutes.POST("",productHandler.SaveProduct())
	productRoutes.PUT("/:id",productHandler.UpdateProduct())
	productRoutes.PATCH("/:id",productHandler.UpdateProductName())
	productRoutes.DELETE("/:id",productHandler.DeleteProduct())

	router.Run(":8080")

	
	

	

	//se define el router
	//router := gin.Default()

	//gopher := router.Group("/products")

	//trae todos los productos
	//gopher.GET("",service.AllProducts)

	//trae productos por id
	//gopher.GET("/:id",service.ProductById)

	//trae productos mayores al precio del parametro priceGt
	//gopher.GET("/search",service.ProductPriceGt)

	//graba un nuevo producto
	//gopher.POST("", service.SaveProduct)


	//router.Run(":8080")

	
}