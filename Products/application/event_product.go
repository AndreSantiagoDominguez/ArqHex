package aplication

import "proyecto_hex/Products/domain"

// GetAllProducts es el caso de uso para obtener la lista de todos los productos.
type GetAllProducts struct {
	db domain.Iproduct // Corrección: Usamos la interfaz Iproduct
}

// NewGetAllProducts crea una nueva instancia del caso de uso.
func NewGetAllProducts(db domain.Iproduct) *GetAllProducts {
	return &GetAllProducts{db: db}
}

// Execute obtiene todos los productos registrados.
func (uc *GetAllProducts) Execute() []domain.Product {
	allProducts := uc.db.GetAllProduct() // Corrección: Llamamos al método correcto de la interfaz

	// No es necesario hacer una conversión de tipo aquí porque `GetAllProduct()` ya devuelve `[]domain.Product`
	return allProducts
}
