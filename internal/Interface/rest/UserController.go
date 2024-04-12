package rest

import (
	entity "ddd/internal/Entity"
	repository "ddd/internal/Repository"
	"ddd/pkg/netutils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Контроллер для HTTP запросов
type UserController struct {
	repo repository.UserRepository
}

func NewUserController(r *gin.Engine, ur *repository.UserRepository) *UserController {

	c := &UserController{
		repo: *ur,
	}

	r.POST("/create", c.Create)
	// r.GET("/delete/:login", c.Delete)

	return c
}

func (uc *UserController) Create(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	log.Println(login)
	log.Println(password)

	user, err := entity.NewUser(login, password)
	if err != nil {
		netutils.ErrReport(err, c)
		return
	}

	err = uc.repo.Create(user)
	if err != nil {
		netutils.ErrReport(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})

}

// func (uc *UserController) Delete(c *gin.Context) {
// 	login := c.Query("login")

// 	err := uc.service.Delete
// }
