package entity

import "fmt"

// Сущность
type User struct {
	Login    string
	Password string
	IsBanned bool
	Money    int
}

// Фабрика
func NewUser(login, password string) (*User, error) {
	if login == "" || password == "" {
		return nil, fmt.Errorf("incorrect user data")
	}
	return &User{
		Login:    login,
		Password: password,
		IsBanned: false,
		Money:    100,
	}, nil
}
