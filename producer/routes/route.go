package routes

import (
	"github.com/ankitdmon/producer/controllers"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.Engine) {
	r.POST("/product", controllers.CreateProduct)
	r.GET("/product/:id", controllers.GetProducts)
}
