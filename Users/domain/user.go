package domain

type User struct {
    Id          int    `json:"id"`
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    Age         int    `json:"age"`
    PhoneNumber string `json:"phone_number"`
}

// NewUser es el constructor para crear un nuevo usuario.
func NewUser(firstName string, lastName string, age int, phoneNumber string) *User {
	return &User{
		Id:          0,
		FirstName:   firstName,
		LastName:    lastName,
		Age:         age,
		PhoneNumber: phoneNumber,
	}
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) SetFirstName(firstName string) {
	u.FirstName = firstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) SetLastName(lastName string) {
	u.LastName = lastName
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) GetPhoneNumber() string {
	return u.PhoneNumber
}

func (u *User) SetPhoneNumber(phoneNumber string) {
	u.PhoneNumber = phoneNumber
}