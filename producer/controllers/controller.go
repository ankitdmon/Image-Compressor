package controllers

import "github.com/gin-gonic/gin"

func CreateProduct (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "product Created",
	})
}