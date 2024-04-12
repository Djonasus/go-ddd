package netutils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrReport(err error, c *gin.Context) {
	log.Println(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
