package aplication

import (
	"proyecto_hex/Users/domain"
)

type EditUser struct {
	db domain.Iuser
}

func NewEditUser(db domain.Iuser) *EditUser {
	return &EditUser{db: db}
}

func (uc *EditUser) EditUser(id int, user domain.User)(uint, error){
	return uc.db.EditUser(id, user)
}