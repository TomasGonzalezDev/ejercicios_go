package handler

import (
	"context"
	"ejercicios_go/ejercicio_supermercado_1/internal/domain"
	"ejercicios_go/ejercicio_supermercado_1/internal/dtos/requests"
	"ejercicios_go/ejercicio_supermercado_1/internal/product"
	"ejercicios_go/ejercicio_supermercado_1/pkg/myerrors"
	"ejercicios_go/ejercicio_supermercado_1/pkg/web"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


type Product struct{
	service product.Service
}

type response struct{
	Message string
}

func NewService(sevice product.Service) *Product{
	return &Product{
		service: sevice,
	}
}

func (e *Product) GetAll() gin.HandlerFunc{


	return func (c *gin.Context){

		ctx := context.Background()
		products, err := e.service.FindAll(ctx)
		
		if err != nil {
			switch err.(type){
			case myerrors.ServerError:
				web.Failure(c,500,err.Error())
				return
			case myerrors.ResourseNotFound:
				web.Failure(c,404,err.Error())
				return
			default:
				web.Failure(c,500,"unexpected error")
				return
			}
		}

		web.Success(c,products)

	}
}

func (e *Product) GetProductById() gin.HandlerFunc{

	return func (c *gin.Context){

		ctx := context.Background()

		id , errConvert := strconv.Atoi(c.Param("id"))

		if errConvert != nil {
			web.Failure(c,400,"invalid id")
			return
		}

		product, err := e.service.FindById(ctx, id)

		if err != nil {
			switch err.(type){
			case myerrors.ServerError:
				web.Failure(c,500,err.Error())
				return
			case myerrors.ResourseNotFound:
				web.Failure(c,404,err.Error())
				return
			default:
				web.Failure(c,500,"unexpected error")
				return
			}
		}

		web.Success(c,product)

	}
}

func (e *Product) GetProductWithPriceGt() gin.HandlerFunc{

	return func (c *gin.Context){

		ctx := context.Background()
		price, errConvert := strconv.ParseFloat(c.Query("priceGt"),5)

		if errConvert != nil {
			web.Failure(c,400,"invalid priceGt")
			return
		}

		products, err := e.service.FindProductByPriceGt(ctx,price)

		if err != nil {
			switch err.(type){
			case myerrors.ServerError:
				web.Failure(c,500,err.Error())
				return
			case myerrors.ResourseNotFound:
				web.Failure(c,404,err.Error())
				return
			default:
				web.Failure(c,500,"unexpected error")
				return
			}
		}

		web.Success(c,products)
	}
}

func validateEmptys(product *domain.Product) error{
	if product.CodeValue == "" || 
		product.Expiration == "" ||
	  	product.Name == "" ||
	 	product.Price == 0 ||
		product.Quantity == 0 {
			return errors.New("fields cannot be empty")
	}
	return nil
}

func validateExpirationDate(exp string) error {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) || (list[1] < 1 || list[1] > 12) || (list[2] < 1 || list[2] > 9999)
	if condition {
		return errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return nil
}

func (e *Product) SaveProduct() gin.HandlerFunc{

	return func(c *gin.Context) {

		expectToken := os.Getenv("TOKEN")
		actualtoken := c.GetHeader("token")
		if actualtoken != expectToken {
			web.Failure(c,401,"unoaturized")
			return
		}
		ctx := context.Background()
		var product domain.Product
		if err := c.ShouldBind(&product); err != nil{
			web.Failure(c,400,"invalid json")
			return
		}
		//validando si hay campos vacios
		if err := validateEmptys(&product); err != nil {
			web.Failure(c,400,err.Error())
			return
		}

		//validando fecha de expiracion
		fmt.Println(validateExpirationDate(product.Expiration))

		if err := validateExpirationDate(product.Expiration); err != nil{
			web.Failure(c,400,err.Error())
			return
		}
		

		savedProduct , err := e.service.Save(ctx,product)

		if err != nil {
			switch err.(type){
			case myerrors.ServerError:
				web.Failure(c,500,err.Error())
				return
			case myerrors.ResourseNotFound:
				web.Failure(c,404,err.Error())
				return
			case myerrors.DuplicatedError:
				web.Failure(c,400,err.Error())
				return
			default:
				web.Failure(c,500,"unexpected error")
				return
			}
		}

		web.Success(c,savedProduct)

	}
}

func (e *Product) UpdateProduct() gin.HandlerFunc{
	return func (c *gin.Context){

		expectToken := os.Getenv("TOKEN")
		actualtoken := c.GetHeader("token")
		if actualtoken != expectToken {
			web.Failure(c,401,"unoaturized")
			return
		}

		var product domain.Product
		ctx := context.Background()
		id, convertError := strconv.Atoi(c.Param("id"))
		if convertError != nil {
			web.Failure(c,400,"invalid id")
			return
		}
		if err := c.ShouldBind(&product); err != nil {
			web.Failure(c,400,"invalid json")
			return
		}

		//validando si hay campos vacios
		if err := validateEmptys(&product); err != nil {
			web.Failure(c,400,err.Error())
			return
		}

		//validando fecha de expiracion
		if err := validateExpirationDate(product.Expiration); err != nil{
			web.Failure(c,400,err.Error())
			return
		}
		

		product, err := e.service.Update(ctx,product,id)

		if err != nil {
			switch err.(type){
			case myerrors.ServerError:
				web.Failure(c,500,err.Error())
				return
			case myerrors.ResourseNotFound:
				web.Failure(c,404,err.Error())
				return
			case myerrors.DuplicatedError:
				web.Failure(c,400,err.Error())
				return
			default:
				web.Failure(c,500,"unexpected error")
				return
			}
		}

		web.Success(c,product)

	}
}

func (e *Product) UpdateProductName() gin.HandlerFunc{
	return func (c *gin.Context){

		expectToken := os.Getenv("TOKEN")
		actualtoken := c.GetHeader("token")
		if actualtoken != expectToken {
			web.Failure(c,401,"unoaturized")
			return
		}
		ctx := context.Background()
        var namedto requests.ProductName
		id, convertError := strconv.Atoi(c.Param("id"))
		if convertError != nil {
			web.Failure(c,400,"invalid id")
			return
		}
		if err := c.ShouldBind(&namedto); err != nil {
			web.Failure(c,400,"invalid json")
			return
		} 

		if strings.Trim(namedto.Name, " ") == "" {
			web.Failure(c,400,"invalid name")
			return
		}


		product, err := e.service.UpdateName(ctx,namedto.Name,id)

		if err != nil {
			switch err.(type){
			case myerrors.ServerError:
				web.Failure(c,500,err.Error())
				return
			case myerrors.ResourseNotFound:
				web.Failure(c,404,err.Error())
				return
			default:
				web.Failure(c,500,"unexpected error")
				return
			}
		}

		web.Success(c,product)

	}
}

func (e *Product) DeleteProduct() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx := context.Background()

		expectToken := os.Getenv("TOKEN")
		actualtoken := c.GetHeader("token")
		if actualtoken != expectToken {
			web.Failure(c,401,"unoaturized")
			return
		}

		id, convertError := strconv.Atoi(c.Param("id"))
		if convertError != nil {
			web.Failure(c,400,"invalid id")
			return
		}

		if err := e.service.Delete(ctx,id); err != nil{
			switch err.(type){
			case myerrors.ServerError:
				web.Failure(c,500,err.Error())
				return
			case myerrors.ResourseNotFound:
				web.Failure(c,404,err.Error())
				return
			default:
				web.Failure(c,500,"unexpected error")
				return
			}
		}

		web.Success(c,response{"producto eliminado"})
	}
}