package store

import (
	"ejercicios_go/ejercicio_supermercado_1/internal/domain"
	"ejercicios_go/ejercicio_supermercado_1/pkg/myerrors"
	"encoding/json"
	"os"
)

type jsonStore struct{
	pathToFile string
}



func NewJsonStore(path string) StoreInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}

func (e *jsonStore) FindAll() ([]domain.Product, error) {
	products, err := e.loadProducts()
	if err != nil {
		return nil, err
	}
	return products,nil
}

func (e *jsonStore) FindById(id int) (domain.Product, error) {
	products, err := e.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, myerrors.ResourseNotFound{Message: "product not found"}
}

func (e *jsonStore) Save(product domain.Product) (domain.Product,error) {
	products, err := e.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}
	product.Id = products[len(products)-1].Id + 1

	products = append(products, product)
	if err := e.saveProducts(products); err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (e *jsonStore) Update(product domain.Product) (domain.Product,error) {
	products, err := e.loadProducts()
	if err != nil {
		return domain.Product{},err
	}
	for i, p := range products {
		if p.Id == product.Id {
			products[i] = product
			return product,e.saveProducts(products)
		}
	}
	return domain.Product{}, myerrors.ResourseNotFound{Message:"product not found"}
}

func (s *jsonStore) Delete(id int) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == id {
			products = append(products[:i], products[i+1:]...)
			return s.saveProducts(products)
		}
	}
	return myerrors.ResourseNotFound{Message: "product not found"}
}
func (s *jsonStore) Exists(codeValue string) (bool, error) {
	products, err := s.loadProducts()
	if err != nil {
		return false, err
	}
	for _, p := range products {
		if p.CodeValue == codeValue {
			return true, nil
		}
	}
	return false, nil
}



//Carga los productos de un json
func (e *jsonStore) loadProducts()([]domain.Product, error){

	var products []domain.Product
	

	data, err := os.ReadFile(e.pathToFile)

	if err != nil {
		return nil, myerrors.ServerError{Message: "repository error cod 1"}

	}

 	if err := json.Unmarshal(data, &products); err != nil{

		return nil, myerrors.ServerError{Message: "repository error cod 2"}
	}
	
	return products, nil
}
//guarda los productos en un json
func (e *jsonStore) saveProducts(products []domain.Product) error {
	bytes, err := json.Marshal(products)
	if err != nil {
		return myerrors.ServerError{Message: "repository error cod 3"}
	}
	return os.WriteFile(e.pathToFile, bytes, 0644)
}