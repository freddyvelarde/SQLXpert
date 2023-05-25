package routes

import (
	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {
	app.POST("/create", createDataBase)
	app.POST("/query", makeQueries)
	app.POST("/databases", getAllDatabases)
	app.POST("/connection", connection)
}
