package aplication

import(
	"proyecto_hex/Users/domain"
)

type CreateUser struct {
	db domain.Iuser
}

func NewCreateUser(db domain.Iuser) *CreateUser {
	return &CreateUser{db: db}
}

func (uc *CreateUser) CreateUser(user domain.User) (uint, error) {
	return uc.db.Save(&user)
}