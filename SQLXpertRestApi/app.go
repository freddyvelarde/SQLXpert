package main

import (
	"github.com/freddyvelarde/SQLXpert/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func app() *gin.Engine {
	app := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://172.19.0.1:7677", "http://localhost:7677"}
	app.Use(cors.New(corsConfig))

	routes.Router(app)

	return app
}
