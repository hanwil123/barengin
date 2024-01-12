package Middleware

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"main.go/Controllers"
	"net/http"
)

func AuthUserMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	opt := option.WithCredentialsFile("./barengin-9e0ca-firebase-adminsdk-hbzc9-93af8c930b.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	// Create a Firebase auth client instance
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Firebase auth c  lient: %v", err)
	}

	verifToken, err := authClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
	Controllers.RegisterUser(c, authClient)
	c.Set("user", verifToken)
}
