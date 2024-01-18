package Controllers

import (
	"firebase.google.com/go/v4/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/Models"
	"math/rand"
	"net/http"
	"strconv"
)

func AddCampaign(c *gin.Context, database *db.Client, name string, target uint64, discount uint64) (string, error) {
	// Generate random ID for campaign
	randomCampaignID := uint(rand.Intn(1000) + 1)

	// Create a new campaign
	newCampaign := Models.Campaigns{
		ID:           uint64(randomCampaignID),
		NameCampaign: name,
		Target:       target,
		Discount:     discount,
	}

	// Push the campaign to the "campaigns" node
	campaignsRef := database.NewRef("campaigns")
	addCampaign, err := campaignsRef.Push(c, newCampaign)
	if err != nil {
		return "", err
	}

	return addCampaign.Key, nil
}

func AddProduct(c *gin.Context, database *db.Client) {
	var data map[string]string
	errs := c.ShouldBindJSON(&data)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind data"})
		return
	}

	// Parse price and stock
	priceStr := data["price"]
	price, err := strconv.ParseUint(priceStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
		return
	}
	stockStr := data["stock"]
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock format"})
		return
	}

	// Generate product ID
	productName := data["nameProduct"]
	productPrefix := productName[:2]
	randomDigits := rand.Intn(90000) + 10000
	productID := fmt.Sprintf("%s%d", productPrefix, randomDigits)

	// Generate random ID for user and campaign
	randomUserID := uint(rand.Intn(10000) + 1)

	// Create a new product
	newProduct := Models.Products{
		ID:          uint64(randomUserID),
		ProductID:   productID,
		CampaignID:
		NameProduct: productName,
		Stock:       uint64(stock),
		Price:       price,
	}

	// Push the product to the "products" node
	productsRef := database.NewRef("products")
	addProduct, err := productsRef.Push(c, newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add product"})
		return
	}

	// Call AddCampaign to add a campaign
	nameCampaign := data["nameCampaign"]
	targetStr := data["target"]
	target, err := strconv.ParseUint(targetStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target format"})
		return
	}
	discountStr := data["discount"]
	discount, err := strconv.ParseUint(discountStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discount format"})
		return
	}

	campaignID, err := AddCampaign(c, database, nameCampaign, target, discount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add campaign"})
		return
	}

	// Return success message along with product and campaign IDs
	c.JSON(http.StatusOK, gin.H{
		"message":    "Product and campaign added successfully",
		"productID":  addProduct.Key,
		"campaignID": campaignID,
	})
}

func GetProducts(c *gin.Context, database *db.Client) {
	// Get a reference to the "products" node
	productsRef := database.NewRef("products")

	// Retrieve all products under "products"
	var products map[string]Models.Products
	if err := productsRef.Get(c, &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	// Convert the map to a slice of Products
	var productList []Models.Products
	for _, product := range products {
		productList = append(productList, product)
	}

	// Send the fetched products in the response
	c.JSON(http.StatusOK, gin.H{"message": "Get products successfully", "products": productList})
}

func DeleteProduct(c *gin.Context, database *db.Client, productID string) {

	productRef := database.NewRef("products").Child(productID)

	errss := productRef.Delete(c)
	if errss != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
