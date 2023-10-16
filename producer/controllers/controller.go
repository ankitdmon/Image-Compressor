package controllers

import (
	"net/http"

	"github.com/ankitdmon/producer/initializers"
	"github.com/ankitdmon/producer/messaging"
	"github.com/ankitdmon/producer/models"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var createProduct models.Product

	if err := c.ShouldBindJSON(&createProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	product, err := models.CreateProduct(createProduct.UserID, createProduct.ProductName, createProduct.ProductDescription, createProduct.ProductImages, createProduct.ProductPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating product"})
		return
	}

	productID, err := models.GetProductId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting product id"})
		return
	}
	if productID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error getting product id"})
		return
	}

	err = messaging.PublishToRabbitMQ(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to RabbitMQ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Created",
		"product": product,
	})
}

func GetProducts(c *gin.Context) {
	productId := c.Param("id")

	var products []models.Product

	// MySQL Query
	query := `
		SELECT 
			p.product_id,
			p.product_name,
			u.name,
			u.mobile,
			p.product_price,
			p.product_images
		FROM
			products p
		LEFT JOIN
			users u ON u.id = p.user_id
		WHERE p.product_id = ?`

	// fmt.Println("Query: ", query)

	db := initializers.DB

	result := db.Raw(query, productId).Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
