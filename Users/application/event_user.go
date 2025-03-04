package aplication

import "proyecto_hex/Users/domain"

type GetAllUsers struct {
	db domain.Iuser
}


func NewGetAllUsers(db domain.Iuser) *GetAllUsers {
	return &GetAllUsers{db: db}
}

// Execute obtiene todos los usuarios registrados.
func (uc *GetAllUsers) Execute() []domain.User {
	allUsers := uc.db.GetAllUsers()
	// Asignamos directamente los usuarios obtenidos, evitando el bucle innecesario.
	users := append([]domain.User{}, allUsers...)
	return users
}
