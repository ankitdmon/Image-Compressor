package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

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

    db := initializers.DB

    productImages := strings.Split(product.ProductImages, ",")

    productImagesJSON, err := json.Marshal(productImages)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to create product",
        })
        return
    }

    query := `
        INSERT INTO products (user_id, product_name, product_description, product_images, product_price, compressed_images, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, NULL, NOW(), NOW())`
    result := db.Exec(query, product.UserID, product.ProductName, product.ProductDescription, productImagesJSON, product.ProductPrice)

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
