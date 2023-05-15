package modules

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createDataBase(ctx *gin.Context) {
	data := struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"dbname": data.Name, "password": data.Password})
}

func mainRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world from modularized golang server",
	})
}
