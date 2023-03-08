package product

import (
	"context"
	"ejercicios_go/ejercicio_supermercado_1/internal/domain"
	"ejercicios_go/ejercicio_supermercado_1/pkg/myerrors"
	"ejercicios_go/ejercicio_supermercado_1/pkg/store"
)

type Repository interface{
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindById(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, codeValue string) (bool, error)
	Save(ctx context.Context, product domain.Product) (domain.Product ,error)
	FindProductByPriceGt(ctx context.Context, double float64) ([]domain.Product, error)
	Update(ctx context.Context,product domain.Product, id int) (domain.Product, error)
	UpdateName(ctx context.Context,name string, id int) (domain.Product, error)
	Delete(ctx context.Context, id int) (error)
}

type reposiotry struct{
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository{
	return &reposiotry{storage}
}

func (r *reposiotry) FindAll(ctx context.Context) ([]domain.Product,error){
	
	products, err :=r.storage.FindAll()
	//products, err := loadData()

	if err != nil {
		return nil ,err
	}

	return products , nil
}

func (r *reposiotry) FindById(ctx context.Context, id int) (domain.Product, error){

	product, err := r.storage.FindById(id)

	if err != nil {
		return domain.Product{} ,err
	}

	return product, nil
}

func (r *reposiotry) Exists(ctx context.Context, codeValue string) (bool, error){

	result, err := r.storage.Exists(codeValue)

	if err != nil {
		return false , myerrors.ServerError{}
	}

	
	return result, nil
}

func (r *reposiotry) Save(ctx context.Context, product domain.Product)(domain.Product, error){

	product, err := r.storage.Save(product)

	if err != nil {
		return domain.Product{} , err
	}

	return product, nil
}

func (r *reposiotry) FindProductByPriceGt(ctx context.Context, double float64) ([]domain.Product, error){

	products, err := r.storage.FindAll()

	var filteredProducts []domain.Product

	if err != nil {
		return nil, err
	}

	for i := range products{
		if products[i].Price >= double{
			filteredProducts = append(filteredProducts, products[i])
		}
	}

	return filteredProducts, nil
}

func (r *reposiotry) Update(ctx context.Context,product domain.Product, id int) (domain.Product, error){
	product.Id = id
	updateProduct, err := r.storage.Update(product)

	if err != nil {
		return domain.Product{}, err
	}

	return updateProduct, nil
}
func (r *reposiotry) UpdateName(ctx context.Context,name string, id int) (domain.Product, error){
	product, errl := r.storage.FindById(id)

	if errl != nil {
		return domain.Product{} , errl
	}

	product.Name = name

	product, err := r.storage.Update(product)
	if err != nil {
		return domain.Product{} , errl
	}
	
	return product, nil
}
func (r *reposiotry) Delete(ctx context.Context, id int) (error){
	return r.storage.Delete(id)
}


