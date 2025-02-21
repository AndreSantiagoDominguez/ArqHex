package aplication

import (
	"proyecto_hex/Users/domain"
)

type GetAll struct {
	db domain.Iuser
}
func NewGetAllUser(db domain.Iuser) *GetAll {
	return &GetAll{db: db}
}

func (uc *GetAll) GetAllUser() ([]domain.User) {
	return uc.db.GetAllUsers()
}