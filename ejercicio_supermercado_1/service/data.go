package service

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)


func SaveProduct(ctx *gin.Context){

	var id int
	var product Product

	if ctx.ShouldBind(&product) != nil {
		ctx.String(400,"bad request")
		return
	}

	products, err := loadData()
	if err != nil{
		ctx.String(500,err.Error())
		return
	}

	//validacion de que los atributos de body no sean nulos
	if product.CodeValue == "" || 
		product.Expiration == "" ||
	  	product.Name == "" ||
	 	product.Price == 0 ||
		product.Quantity == 0 {
			ctx.String(400,"ningun campo puede ser vacio")
			return
		}

	//validar que el codigo del producto sea unico

	if !IsAUnicCodeValue(products, product.CodeValue){
		ctx.String(400,"code value already exist")
		return
	}
	//validar fecha
	if flag, message := IsAValidDate(product.Expiration); !flag{
		ctx.String(400,message)
		return
	}
	

	id = len(products) + 1

	product.Id = id

	ctx.IndentedJSON(200,product)

	savePoduct(product)

	
}

func AllProducts(ctx *gin.Context){

	products, err := loadData()

	if err != nil{
		ctx.String(500,"server error")
		return
	}
	

	ctx.IndentedJSON(200,products)
}

func ProductById(ctx *gin.Context){

	products, err := loadData()

	if err != nil{
		ctx.String(500,"server error")
		return
	}

	
	var filteredProducts []Product = []Product{} 

	id , err1 := strconv.Atoi(ctx.Param("id"))

	if err1 != nil{
		ctx.String(400,"bad request")
		return
	}

	

	for i := range products{
		if products[i].Id == id {
			filteredProducts = append(filteredProducts, products[i])
		}
	}


	ctx.IndentedJSON(200,filteredProducts)
}

func ProductPriceGt(ctx *gin.Context){
	products, err := loadData()

	if err != nil{
		ctx.String(500,"server error")
		return
	}

	
	var filteredProducts []Product = []Product{} 

	price , err1 := strconv.ParseFloat(ctx.Query("priceGt"),2)

	if err1 != nil{
		ctx.String(400,"bad request")
		return
	}

	

	for i := range products{
		if products[i].Price > price {
			filteredProducts = append(filteredProducts, products[i])
		}
	}


	ctx.IndentedJSON(200,filteredProducts)
}

func IsAValidDate(date string) (bool, string){
	if result, _ := regexp.Match("[0-3][0-9]/[0-1][0-9]/[0-9]{4}", []byte(date)); !result{
		return false, "formato de fecha no valido"
	}

	dateParams := strings.Split(date,"/")
	anno, _ := strconv.Atoi(dateParams[2])
	mes, _ := strconv.Atoi(dateParams[1])
	day, _ := strconv.Atoi(dateParams[0])
	//theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	productDate := time.Date(anno, time.Month(mes), day,0,0,0,0,time.Local)
	currentDate := time.Now()
	if !productDate.After(currentDate){
		return false, "la fecha debe ser futura"
	}

	return true, "ok"


}

func IsAUnicCodeValue(products []Product, codeValue string) bool{
	for i := range products {
		if products[i].CodeValue == codeValue{
			return false
		}
	}
	return true
}

func loadData()([]Product, error) {

	//file, err := os.Open("data.json")

	//if err != nil {
	//	return nil, err
	//}

	var products []Product
	

	data, err := os.ReadFile("data.json")

	if err != nil {
		return nil, err
	}

 	if err := json.Unmarshal(data, &products); err != nil{

		return nil, err
	}
	
	return products, nil
}

func savePoduct(product Product) (Product, error){

	file, err := os.OpenFile("data.json", os.O_CREATE | os.O_WRONLY, os.ModePerm)
	products, _ := loadData()

	product.Id = len(products) + 1
	if err != nil {
		return Product{} , err
	}

	

	products = append(products,product)
	byteJson, _:= json.Marshal(products)
	fmt.Println(product)
	file.Write(byteJson)

	

	return product, nil
}

type Product struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	CodeValue string `json:"code_value"`
	IsPublised bool `json:"is_publised"`
	Expiration string `json:"expiration"`
	Price float64 `json:"price"`
}