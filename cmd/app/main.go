package main

import (
	infrastructure "ddd/internal/Infrastructure"
	"ddd/internal/Interface/rest"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(infrastructure.UserModel{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// Запускаем слой инфраструтктуры
	userInf := infrastructure.NewUserInfrastructure(db)

	// Сервис - устарел, вместо него контроллер, который является презентором

	// Контроллер
	rest.NewUserController(r, &userInf)

	r.Run()
}
