package routes

import (
	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {
	app.GET("/", mainRoute)
	app.POST("/create", createDataBase)
	app.POST("/query", makeQueries)
	app.POST("/databases", getAllDatabases)
}
