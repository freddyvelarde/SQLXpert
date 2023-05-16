package main

import (
	"github.com/freddyvelarde/SQLXpert/modules"
	"github.com/gin-gonic/gin"
)

func app() *gin.Engine {
	app := gin.Default()

	modules.Router(app)

	return app
}
