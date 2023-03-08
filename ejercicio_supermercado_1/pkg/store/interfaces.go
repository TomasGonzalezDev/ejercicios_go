package store

import "ejercicios_go/ejercicio_supermercado_1/internal/domain"

type StoreInterface interface {
	FindAll() ([]domain.Product, error)
	// Read devuelve un producto por su id
	FindById(id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Save(product domain.Product) (domain.Product, error)
	// Update actualiza un producto
	Update(product domain.Product) (domain.Product, error)
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
	Exists(codeValue string) (bool, error)
}