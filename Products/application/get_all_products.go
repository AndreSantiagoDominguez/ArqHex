package aplication

import (
	"proyecto_hex/Products/domain"
)

type GetAll struct {
	db domain.Iproduct
}

func NewGetAllProduct(db domain.Iproduct) *GetAll {
	return &GetAll{db: db}
}

func (uc *GetAll) GetAllProduct() ([]domain.Product) {
	return uc.db.GetAllProduct()
}