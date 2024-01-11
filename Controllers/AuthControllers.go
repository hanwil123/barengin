package Controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/Databases"
	"main.go/Models"
	"math/rand"
	"net/http"
	"time"
)

func RegisterUser(c *gin.Context) {
	var data map[string]string
	err := c.ShouldBind(&data)
	if err != nil {
		panic(err)
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	rand.NewSource(time.Now().UnixNano())
	randomUserID := uint(rand.Intn(10000) + 1)
	user := Models.AuthUsers{
		ID:       uint64(randomUserID),
		Email:    data["email"],
		Password: password,
		Name:     data["name"],
	}
	result := Databases.UDB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
