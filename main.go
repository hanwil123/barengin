package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"main.go/Controllers"

	"net/http"
)

func main() {
	conf := &firebase.Config{
		DatabaseURL: "https://barengin-9e0ca-default-rtdb.asia-southeast1.firebasedatabase.app",
	}
	opt := option.WithCredentialsFile("./barengin-9e0ca-firebase-adminsdk-hbzc9-93af8c930b.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	// Create a Firebase auth client instance
	authClient, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	apps, errs := firebase.NewApp(context.Background(), conf, opt)
	database, errs := apps.Database(context.Background())
	if errs != nil {
		panic(errs)
	}
	// Create a new Gin router
	r := gin.Default()

	// Apply middleware for authentication

	// Define routes
	r.POST("/register", func(c *gin.Context) {
		Controllers.RegisterUser(c, authClient)
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})
	r.POST("/login", func(c *gin.Context) {
		Controllers.LoginUser(c, authClient)
		c.JSON(http.StatusOK, gin.H{"message": "User Login successfully"})
	})
	r.POST("/product", func(c *gin.Context) {
		Controllers.AddProduct(c, database)
	})
	r.GET("/product", func(c *gin.Context) {

		// Call UpdateProduct with the necessary arguments
		Controllers.GetProducts(c, database) // Pass productID as an argument
		c.JSON(http.StatusOK, gin.H{"message": "Get Product Successfully"})
	})
	r.DELETE("/product", func(c *gin.Context) {
		productID := c.Param("productID")

		// Call UpdateProduct with the necessary arguments
		Controllers.DeleteProduct(c, database, productID) // Pass productID as an argument
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	// Run the server
	r.Run(":8080")

}
