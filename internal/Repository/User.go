package repository

import entity "ddd/internal/Entity"

// Репозиторий
type UserRepository interface {
	Create(user *entity.User) error
	FindByLogin(login string) (*entity.User, error)
	Delete(*entity.User) error
	Update(*entity.User) error
}
