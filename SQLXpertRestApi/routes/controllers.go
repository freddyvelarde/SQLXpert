package routes

import (
	"net/http"

	"github.com/freddyvelarde/SQLXpert/repositories"
	"github.com/freddyvelarde/SQLXpert/structs"
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

	var dbConfig structs.DBConfig

	dbConfig.Port = data.Port
	dbConfig.User = data.User
	dbConfig.Password = data.Password
	dbConfig.DbName = data.DbName
	dbConfig.Host = data.Host

	response, err := repositories.CreateNewDatabase(data.NewDB, dbConfig)
	if err != nil {
		ctx.JSON(http.StatusConflict, err)
	}

	ctx.JSON(http.StatusAccepted, response)
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

	dbConfig := structs.DBConfig{
		Port:     body.Port,
		User:     body.User,
		Password: body.Password,
		DbName:   body.DbName,
		Host:     body.Host,
	}

	tableNames, _ := repositories.GetTableNames(dbConfig)
	// fmt.Println(d)
	data, err := repositories.Queries(utils.NormalizeQuery(body.Query, tableNames), dbConfig)
	// data, err := repositories.Queries(body.Query, dbConfig)
	if err != nil {
		ctx.JSON(http.StatusAccepted, err)
		return
	}

	columns, err := repositories.GetAllColumnNamesfromTable(dbConfig)
	if err != nil {
		ctx.JSON(http.StatusAccepted, err)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"data": data, "columns": columns})
	// ctx.JSON(http.StatusAccepted, data)
}

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

	dbConfig := structs.DBConfig{
		Port:     body.Port,
		User:     body.User,
		Password: body.Password,
		DbName:   body.DbName,
		Host:     body.Host,
	}

	res, err := repositories.GetAllDatabases(dbConfig)
	if err != nil {
		ctx.JSON(http.StatusAccepted, err)
		return
	}

	ctx.JSON(http.StatusAccepted, res)
}

func mainRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world from modularized golang server",
	})
}
