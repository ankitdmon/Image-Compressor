package main

import (
	"github.com/ankitdmon/producer/initializers"
	"github.com/ankitdmon/producer/models"
	"github.com/ankitdmon/producer/utils"
)

func init() {
	initializers.LoadENV()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Product{})
	utils.InsertInitialUserData(initializers.DB)
}
