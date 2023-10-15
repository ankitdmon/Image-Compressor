package main

import (
	"github.com/ankitdmon/producer/initializers"
	"github.com/ankitdmon/producer/models"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Product{})
}
