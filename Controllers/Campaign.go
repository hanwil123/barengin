package Controllers

import (
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/Models"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func AddProduct(c *gin.Context, app *firebase.App) {
	client, err := app.Database(c)
	if err != nil {
		panic(err)
	}
	var data map[string]string
	errs := c.ShouldBind(&data)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}
	priceStr := data["price"]
	priceFloat, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
		return
	}
	productName := data["nameProduct"]
	productPrefix := productName[:2]
	rand.NewSource(time.Now().UnixNano())
	randomDigits := rand.Intn(90000) + 10000                      // Generate a random 5-digit number
	productID := fmt.Sprintf("%s%d", productPrefix, randomDigits) // Combine prefix and digits
	randomUserID := uint(rand.Intn(10000) + 1)
	randomCampaignID := uint(rand.Intn(1000) + 1)
	products := Models.Products{
		ID:          uint64(randomUserID),
		ProductID:   productID,
		NameProduct: productName,
		Price:       priceFloat,
		CampaignID:  uint64(randomCampaignID),
	}
	ref := client.NewRef("products")
	addProduct, errr := ref.Push(c, products)
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid add Product"})
	} else {
		// Call UpdateProduct immediately after successful addition
		go UpdateProduct(c, app, productID) // Use a goroutine to avoid blocking

		c.JSON(http.StatusOK, gin.H{"message": "Product added successfully", "productID": addProduct})
	}

}
func UpdateProduct(c *gin.Context, app *firebase.App, productID string) {
	client, err := app.Database(c)
	if err != nil {
		panic(err)
	}

	// Get updated data from request
	var updateData map[string]interface{}
	if err := c.ShouldBind(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
		return
	}

	productRef := client.NewRef("products").Child(productID)
	errr := productRef.Update(c, updateData)
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
func DeleteProduct(c *gin.Context, app *firebase.App, productID string) {
	client, err := app.Database(c)
	if err != nil {
		panic(err)
	}

	productRef := client.NewRef("products").Child(productID)

	errss := productRef.Delete(c)
	if errss != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
