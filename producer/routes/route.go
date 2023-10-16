package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/ankitdmon/producer/controllers"
)

func SetupProductRoutes(r *gin.Engine) {
    r.POST("/product", controllers.CreateProduct)
    r.GET("/product/:id", controllers.GetProducts)
}
