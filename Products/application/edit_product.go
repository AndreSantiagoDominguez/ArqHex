package aplication

import (
	"proyecto_hex/Products/domain"
)

type EditProduct struct {
	db domain.Iproduct
}

func (uc *EditProduct) EditProduct(id int, param any) (any, error) {
	panic("unimplemented")
}

func NewEditProduct(db domain.Iproduct) *EditProduct {
	return &EditProduct{db: db}
}

func (uc *EditProduct) EditProductProcess(id int, product domain.Product) (uint, error) {
	return uc.db.EditProduct(id, product)

}
