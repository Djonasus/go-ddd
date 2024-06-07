package infrastructure

import (
	entity "ddd/internal/Entity"
	repository "ddd/internal/Repository"
	"fmt"

	"gorm.io/gorm"
)

// Модель для БД
type UserModel struct {
	gorm.Model
	Login    string
	Password string
	IsBanned bool
	Money    int
}

// Преобразование модели в сущность
func UserModelToEntity(um *UserModel) *entity.User {
	return &entity.User{
		Login:    um.Login,
		Password: um.Password,
		IsBanned: um.IsBanned,
		Money:    um.Money,
	}
}

// Преобразование сущности в модель
func UserEntityToModel(u *entity.User) *UserModel {
	return &UserModel{
		Login:    u.Login,
		Password: u.Password,
		IsBanned: u.IsBanned,
		Money:    u.Money,
	}
}

// Инфраструктура - Связь репозиториев с БД
type UserInfrastructure struct {
	Db *gorm.DB
}

// Фабрика
func NewUserInfrastructure(db *gorm.DB) repository.UserRepository {
	return &UserInfrastructure{
		Db: db,
	}
}

// Реализации интерфейса

// Create implements UserRepository.
func (ui *UserInfrastructure) Create(user *entity.User) error {
	// modeledUser := UserEntityToModel(user)
	result := ui.Db.Limit(1).Where(UserModel{Login: user.Login}).Find(&UserModel{})
	if result.RowsAffected != 0 {
		return fmt.Errorf("user already exists")
	}
	res := ui.Db.Create(UserEntityToModel(user))
	return res.Error
}

// Delete implements UserRepository.
func (ui *UserInfrastructure) Delete(u *entity.User) error {
	modeledUser := UserEntityToModel(u)
	res := ui.Db.Delete(modeledUser)
	return res.Error
}

// FindByLogin implements UserRepository.
func (ui *UserInfrastructure) FindByLogin(login string) (*entity.User, error) {
	modeledUser := UserEntityToModel(&entity.User{Login: login})
	res := ui.Db.Limit(1).Find(modeledUser)
	return UserModelToEntity(modeledUser), res.Error
}

// Update implements UserRepository.
func (ui *UserInfrastructure) Update(u *entity.User) error {
	modeledUser := UserEntityToModel(u)
	res := ui.Db.Save(modeledUser)
	return res.Error
}
