package domain

type Iuser interface {
	Save(user *User) (uint, error)
	GetAllUsers() []User;
	EditUser(id int, user User) (uint, error)
	DeleteUser(id int) (uint, error)
}

