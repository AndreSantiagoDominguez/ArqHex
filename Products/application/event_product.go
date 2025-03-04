package aplication

import "proyecto_hex/Products/domain"

// GetAllProducts es el caso de uso para obtener la lista de todos los productos.
type GetAllProducts struct {
	db domain.Iproduct 
}


func NewGetAllProducts(db domain.Iproduct) *GetAllProducts {
	return &GetAllProducts{db: db}
}

// Execute obtiene todos los productos registrados.
func (uc *GetAllProducts) Execute() []domain.Product {
	allProducts := uc.db.GetAllProduct() 

	
	return allProducts
}
