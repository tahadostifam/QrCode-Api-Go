package controllers

import (
	"fmt"
	"net/http"
	"strings"

	models "$1/models"

	"github.com/gin-gonic/gin"
)

func RootAction(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func MakeUserAction(c *gin.Context) {
	user := &models.User{
		Name:     "Taha",
		Username: "tahadostifam",
		Password: "1234",
		Age:      15,
	}

	callback := fmt.Sprintf(
		strings.TrimSpace(`
name : %s
username : %s
age : %d
	`), user.Name, user.Username, user.Age)

	c.String(http.StatusOK, callback)
}
