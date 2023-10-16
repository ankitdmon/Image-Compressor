package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ankitdmon/producer/initializers"
	"github.com/ankitdmon/producer/models"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Convert to JSON Array
	productImagesJSON, err := json.Marshal(product.ProductImages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process product images",
		})
		return
	}

	// // Convert to JSON array
	// compressedImagesJSON, err := json.Marshal(product.CompressedImages)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Failed to process compressed product images",
	// 	})
	// 	return
	// }

	product.ProductImages = []string{string(productImagesJSON)}
	//product.CompressedImages = []string{string(compressedImagesJSON)}

	db := initializers.DB

	result := db.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Created",
		"product": product,
	})
}

func GetProducts(c *gin.Context) {
	var products []models.Product

	// MySQL Query
	query := `
        SELECT 
            p.product_id,
            p.product_name,
            u.name,
            u.mobile,
            p.product_price
        FROM
            products p
        LEFT JOIN
            users u ON u.id = p.user_id
    `

	fmt.Println("Query: ", query)

	db := initializers.DB

	result := db.Raw(query).Find(&products)
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
