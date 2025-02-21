package aplication

import (
	"proyecto_hex/Products/domain"
)

type EditProduct struct {
	db domain.Iproduct
}

func NewEditProduct(db domain.Iproduct) *EditProduct {
	return &EditProduct{db: db}
}

func (uc *EditProduct) EditProduct(id int, product domain.Product) (uint, error) {
	return uc.db.EditProduct(id, product)

}
