package modules

import (
	"net/http"

	"github.com/freddyvelarde/SQLXpert/repositories"
	"github.com/freddyvelarde/SQLXpert/utils"
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

	var dbConfig utils.DBConfig

	dbConfig.Port = data.Port
	dbConfig.User = data.User
	dbConfig.Password = data.Password
	dbConfig.DbName = data.DbName
	dbConfig.Host = data.Host

	response, err := repositories.CreateNewDatabase(data.NewDB, dbConfig)
	if err != nil {
		ctx.JSON(http.StatusConflict, err)
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": response})
}

func makeQueries(ctx *gin.Context) {
	body := struct {
		DbName   string `json:"dbName"`
		Password string `json:"password"`
		User     string `json:"user"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Query    string `json:"query"`
	}{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbConfig := utils.DBConfig{
		Port:     body.Port,
		User:     body.User,
		Password: body.Password,
		DbName:   body.DbName,
		Host:     body.Host,
	}

	data := repositories.Queries(body.Query, dbConfig)
	columns := repositories.GetAllColumnNamesfromTable(dbConfig)

	ctx.JSON(http.StatusAccepted, gin.H{"data": data, "columns": columns})
	// ctx.JSON(http.StatusAccepted, data)
}

// func databaseConnection(ctx *gin.Context) {
//   body := struct {
//     DbName   string `json:"dbName"`
//     Password string `json:"password"`
//     User     string `json:"user"`
//     Host     string `json:"host"`
//     Port     int    `json:"port"`
//   }{}
//
//   dbConfig := utils.DBConfig{
//     Port:     body.Port,
//     User:     body.User,
//     Password: body.Password,
//     DbName:   body.DbName,
//     Host:     body.Host,
//   }
//
//   _, err := repositories.Connection(dbConfig)
//   if err != nil {
//     ctx.JSON(http.StatusBadRequest, err)
//     return
//   }
//
//   ctx.JSON(http.StatusAccepted, gin.H{"message": "Database connected successfully"})
// }

func getAllDatabases(ctx *gin.Context) {
	body := struct {
		DbName   string `json:"dbName"`
		Password string `json:"password"`
		User     string `json:"user"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
	}{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	dbConfig := utils.DBConfig{
		Port:     body.Port,
		User:     body.User,
		Password: body.Password,
		DbName:   body.DbName,
		Host:     body.Host,
	}

	res := repositories.GetAllDatabases(dbConfig)

	ctx.JSON(http.StatusAccepted, res)
}

func mainRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world from modularized golang server",
	})
}
