package main

import (
	"github.com/freddyvelarde/SQLXpert/modules"
	"github.com/gin-gonic/gin"
)

func app() *gin.Engine {
	app := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	modules.Router(app)

	return app
}
