package main

import (
	"github.com/madmaxieee/dcard-2024-backend-intern-assignment/initializers"
	"github.com/madmaxieee/dcard-2024-backend-intern-assignment/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Advertisement{})
	initializers.DB.AutoMigrate(&models.Condition{})
}
