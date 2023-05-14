package main

import (
	"github.com/freddyvelarde/SQLXpert/modules/config"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": config.DbConnection("localhost", "admin", "password", 5432),
		})
	})

	app.Run()
}
