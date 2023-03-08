package product

import (
	"context"
	"ejercicios_go/ejercicio_supermercado_1/internal/domain"
	"ejercicios_go/ejercicio_supermercado_1/pkg/myerrors"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)


type Service interface{
	FindAll(ctx context.Context)([]domain.Product, error)
	FindById(ctx context.Context, id int)(domain.Product,error)
	Save(ctx context.Context, product domain.Product)(domain.Product, error)
	FindProductByPriceGt(ctx context.Context, double float64) ([]domain.Product, error)
	Update(ctx context.Context,product domain.Product, id int) (domain.Product, error)
	UpdateName(ctx context.Context,name string, id int) (domain.Product, error)
	Delete(ctx context.Context, id int) (error)
}

type service struct{
	repository Repository
}

func NewService(repository Repository) Service{
	return &service{
		repository,
	}
}

func (s *service) FindAll(ctx context.Context) ([]domain.Product,error){
	return s.repository.FindAll(ctx)
}

func (s *service) FindById(ctx context.Context, id int) (domain.Product, error){
	return s.repository.FindById(ctx,id)
}

func (s *service)FindProductByPriceGt(ctx context.Context, double float64) ([]domain.Product, error){

	
	return s.repository.FindProductByPriceGt(ctx,double)
}

func (s *service) Save(ctx context.Context, product domain.Product) (domain.Product, error){

	
	//code value
	result, err := s.repository.Exists(ctx,product.CodeValue)

	if err != nil {
		return domain.Product{}, err
	}
	if result {
		return domain.Product{} , myerrors.DuplicatedError{Message: "code value already exist"}
	}
	


	return s.repository.Save(ctx,product)
}

func (s *service) Update(ctx context.Context,product domain.Product, id int) (domain.Product, error){
	
	//code value
	result, err := s.repository.Exists(ctx,product.CodeValue)

	if err != nil {
		return domain.Product{}, err
	}
	

	oldProduct, err := s.repository.FindById(ctx,id)
	if err != nil {
		return domain.Product{}, err
	}
	if result && oldProduct.CodeValue != product.CodeValue {
		return domain.Product{} , myerrors.DuplicatedError{Message: "code value already exist"}
	}
	
	return s.repository.Update(ctx,product,id)
}
func (s *service) UpdateName(ctx context.Context,name string, id int) (domain.Product, error){
	
	return s.repository.UpdateName(ctx,name,id)
}
func (s *service) Delete(ctx context.Context, id int) (error){

	return s.repository.Delete(ctx, id)
}

func IsAValidExpirationDate(date string) (error){
	if result, _ := regexp.Match(`(3[0-1]|[0-2][0-9])/(0[0-9]|1[0-2])/([0-9]{4})`, []byte(date)); !result{
		return errors.New("formato de fecha no valido")
	}

	dateParams := strings.Split(date,"/")
	anno, _ := strconv.Atoi(dateParams[2])
	mes, _ := strconv.Atoi(dateParams[1])
	day, _ := strconv.Atoi(dateParams[0])
	productDate := time.Date(anno, time.Month(mes), day,0,0,0,0,time.Local)
	currentDate := time.Now()
	if !productDate.After(currentDate){
		return errors.New("la fecha debe ser futura")
	}

	return nil
}