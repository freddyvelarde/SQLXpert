package routes

import (
	"net/http"

	"github.com/freddyvelarde/SQLXpert/repositories"
	"github.com/freddyvelarde/SQLXpert/structs"
	"github.com/freddyvelarde/SQLXpert/utils"
	"github.com/gin-gonic/gin"
)

func connection(c *gin.Context) {
	body := struct {
		DbName   string `json:"dbName"`
		Password string `json:"password"`
		User     string `json:"user"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		NewDB    string `json:"newDB"`
	}{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbConfig := structs.DBConfig{
		Port:     body.Port,
		User:     body.User,
		Password: body.Password,
		DbName:   body.DbName,
		Host:     body.Host,
	}

	tables, err := repositories.GetTableNames(dbConfig)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "bad request"})
	}

	c.JSON(http.StatusAccepted, gin.H{"tablenames": tables, "status": http.StatusAccepted})
}

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Failed request"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"response": response, "status": http.StatusAccepted})
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
	data, err := repositories.Queries(utils.NormalizeQuery(body.Query, tableNames), dbConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Failed request"})
		return
	}

	columns, err := repositories.GetAllColumnNamesfromTable(dbConfig, body.Query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Failed request"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"data": data, "columns": columns, "status": http.StatusAccepted})
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
		return
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Failed request"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"status": http.StatusAccepted, "databases": res})
}

func mainRoute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world from modularized golang server",
	})
}
