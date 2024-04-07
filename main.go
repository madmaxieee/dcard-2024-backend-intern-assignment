package main

import (
	"github.com/gin-gonic/gin"
	"github.com/madmaxieee/dcard-2024-backend-intern-assignment/controllers"
	"github.com/madmaxieee/dcard-2024-backend-intern-assignment/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	app := gin.Default()
	app.GET("/api/v1/ad", controllers.CreateAdvertisement)

	app.Run()
}
