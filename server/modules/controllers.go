package modules

import (
	"net/http"

	"github.com/freddyvelarde/SQLXpert/config"
	"github.com/gin-gonic/gin"
)

func createDataBase(ctx *gin.Context) {
	data := struct {
		DbName   string `json:"dbName"`
		Password string `json:"password"`
		User     string `json:"user"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		NewDB    string `json:"newDB"`
	}{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbConfig config.DBConfig

	dbConfig.Port = data.Port
	dbConfig.User = data.User
	dbConfig.Password = data.Password
	dbConfig.DbName = data.DbName
	dbConfig.Host = data.Host

	res := config.CreateNewDatabase(data.NewDB, dbConfig)

	ctx.JSON(http.StatusAccepted, gin.H{"message": res})
}

func mainRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world from modularized golang server",
	})
}
