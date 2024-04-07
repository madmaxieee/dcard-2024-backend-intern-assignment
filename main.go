package main

import (
	"github.com/gin-gonic/gin"
	"github.com/madmaxieee/dcard-2024-backend-intern-assignment/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	app.Run()
}
