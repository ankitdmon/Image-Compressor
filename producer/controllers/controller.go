package controllers

import (
	"encoding/json"
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

	// Convert to JSON array
	compressedImagesJSON, err := json.Marshal(product.CompressedImages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process compressed product images",
		})
		return
	}

	product.ProductImages = []string{string(productImagesJSON)}
	product.CompressedImages = []string{string(compressedImagesJSON)}

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