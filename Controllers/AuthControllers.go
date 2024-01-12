package Controllers

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/Models"
	"math/rand"
	"net/http"
	"time"
)

func RegisterUser(c *gin.Context, authClient *auth.Client) {
	var data map[string]string
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

	// Create user in Firebase Authentication
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password(string(password)).
		DisplayName(user.Name)

	createdUser, err := authClient.CreateUser(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a reference to the user in the Realtime Database

	c.JSON(http.StatusOK, createdUser)
}
func LoginUser(c *gin.Context, authClient *auth.Client) {
	var data map[string]string
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify if the email and password are provided
	email := data["email"]
	password := data["password"]
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Verify user credentials using Firebase Authentication

	u, err := authClient.GetUserByEmail(context.Background(), email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// You can customize the response based on your needs
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": u})
}
