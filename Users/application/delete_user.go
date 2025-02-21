package aplication

import (
	"proyecto_hex/Users/domain"
)

type DeleteUser struct {
	db domain.Iuser
}

func NewDeleteUser(db domain.Iuser) *DeleteUser {
	return &DeleteUser{db: db}
}

func (uc *DeleteUser) Run(id int) (uint, error){
	return uc.db.DeleteUser(id)
}