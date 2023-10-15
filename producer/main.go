package main

import (
	"fmt"

	"github.com/ankitdmon/producer/initializers"
	"github.com/gin-gonic/gin"
)

func init (){
	initializers.LoadENV()
}

func main() {
	fmt.Println("Hello world!!")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.Run()
}
