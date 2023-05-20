package main

import (
	"github.com/freddyvelarde/SQLXpert/modules"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func app() *gin.Engine {
	app := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://172.19.0.1:7677", "http://localhost:7677", "http://192.168.0.9:7677"}
	app.Use(cors.New(corsConfig))

	modules.Router(app)

	return app
}
