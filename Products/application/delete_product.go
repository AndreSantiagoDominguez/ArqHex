package aplication

import (
	"proyecto_hex/Products/domain"
)

type DeleteProduct struct {
	db domain.Iproduct
}

func NewDeleteProduct(db domain.Iproduct) *DeleteProduct {
	return &DeleteProduct{db: db}
}

func (uc *DeleteProduct) Run(id int) (uint, error) {
	return uc.db.DeleteProduct(id)
}