package routes

import (
	"net/http"
	"strings"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"connected": false,
			"message":   "bad request",
			"status":    http.StatusBadRequest,
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error":     err,
			"message":   "bad request",
			"connected": false,
			"status":    http.StatusBadRequest,
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"tablenames": tables,
		"status":     http.StatusAccepted,
		"connected":  true,
	})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"failed": true,
			"status": http.StatusBadRequest,
		})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed request",
			"failed":  true, "status": http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"response": response,
		"status":   http.StatusAccepted,
		"failed":   false,
	})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"failed": true,
			"data":   nil,
			"status": http.StatusBadRequest,
		})
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

	queryNormalized := utils.NormalizeQuery(strings.ReplaceAll(body.Query, "\n", "  "), tableNames)
	dataSelected, err := repositories.Queries(queryNormalized, dbConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"failed":  true,
			"data":    nil,
			"message": "Failed query",
			"status":  http.StatusBadRequest,
		})
		return
	}

	columns, err := repositories.GetAllColumnNamesfromTable(dbConfig, strings.ReplaceAll(body.Query, "\n", "  "))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed to get the column names",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data":    dataSelected,
		"columns": columns,
		"status":  http.StatusBadRequest,
		"failed":  false,
	})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"failed": true,
			"status": http.StatusBadRequest,
		})
		return
	}

	dbConfig := structs.DBConfig{
		Port:     body.Port,
		User:     body.User,
		Password: body.Password,
		DbName:   body.DbName,
		Host:     body.Host,
	}

	databases, err := repositories.GetAllDatabases(dbConfig)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed request", "failed": true, "status": http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"status":    http.StatusAccepted,
		"databases": databases, "failed": false,
	})
}
