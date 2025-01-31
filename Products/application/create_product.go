package aplication

import (
	"proyecto_hex/Products/domain"
)

type CreateProduct struct {
	db domain.Iproduct
}

func NewCreateProduct(db domain.Iproduct) *CreateProduct {
	return &CreateProduct{db: db}
}

func (uc *CreateProduct) CreateProduct(product domain.Product) (uint, error) {
	return uc.db.Save(&product)
}