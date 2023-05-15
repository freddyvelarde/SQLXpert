package main

import (
	"github.com/freddyvelarde/SQLXpert/modules/config"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	dbConfig := config.DbConfig{Host: "localhost", Password: "admin", Port: 5432, User: "admin"}
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			// "message": dbConfig.Querie("create table \"user\" (id serial primary key, name text, email text not null, password text);", "gotest"),
			"message": dbConfig.Querie("select * from \"user\";", "gotest"),
		})
	})

	app.GET("/create", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": dbConfig.CreateDataBase("eli"),
		})
	})

	app.Run()
}
