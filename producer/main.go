package main

import (
	"fmt"

	"github.com/ankitdmon/producer/initializers"
	"github.com/ankitdmon/producer/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Hello world!!")
	r := gin.Default()

	routes.SetupProductRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.Run()
}
