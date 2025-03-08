package aplication

import (
	"proyecto_hex/Purchases/domain"
)

type CreatePurchase struct {
	db domain.IPurchase
}

func NewCreatePurchase(db domain.IPurchase) *CreatePurchase{
	return &CreatePurchase{
		db:db}
}

func (cp *CreatePurchase) Shopping(purchase *domain.Purchase) error {
	_, err := cp.db.Shopping(purchase)
	return err
}