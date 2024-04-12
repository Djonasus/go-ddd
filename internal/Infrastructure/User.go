package infrastructure

import (
	entity "ddd/internal/Entity"
	repository "ddd/internal/Repository"

	"gorm.io/gorm"
)

// Инфраструктура - Связь репозиториев с БД
type UserInfrastructure struct {
	Db *gorm.DB
}

func NewUserInfrastructure(db *gorm.DB) repository.UserRepository {
	return &UserInfrastructure{
		Db: db,
	}
}

// Реализации интерфейса

// Create implements user.UserRepository.
func (ui *UserInfrastructure) Create(user *entity.User) error {
	modeledUser := UserEntityToModel(user)
	res := ui.Db.Create(modeledUser)
	return res.Error
}

// Delete implements user.UserRepository.
func (ui *UserInfrastructure) Delete(u *entity.User) error {
	modeledUser := UserEntityToModel(u)
	res := ui.Db.Delete(modeledUser)
	return res.Error
}

// FindByLogin implements user.UserRepository.
func (ui *UserInfrastructure) FindByLogin(login string) (*entity.User, error) {
	modeledUser := UserEntityToModel(&entity.User{Login: login})
	res := ui.Db.Limit(1).Find(modeledUser)
	return UserModelToEntity(modeledUser), res.Error
}

// Update implements user.UserRepository.
func (ui *UserInfrastructure) Update(u *entity.User) error {
	modeledUser := UserEntityToModel(u)
	res := ui.Db.Save(modeledUser)
	return res.Error
}

func UserModelToEntity(um *UserModel) *entity.User {
	return &entity.User{
		Login:    um.Login,
		Password: um.Password,
		IsBanned: um.IsBanned,
		Money:    um.Money,
	}
}

func UserEntityToModel(u *entity.User) *UserModel {
	return &UserModel{
		Login:    u.Login,
		Password: u.Password,
		IsBanned: u.IsBanned,
		Money:    u.Money,
	}
}

// Модель для БД
type UserModel struct {
	gorm.Model
	Login    string
	Password string
	IsBanned bool
	Money    int
}
