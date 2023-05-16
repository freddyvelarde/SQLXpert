package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Port     int
	DbName   string
}

func connection(config DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

type CreateDBResponse struct {
	Error     error  `json:"error"`
	DBCreated bool   `json:"dbCreated"`
	Message   string `json:"message"`
}

func CreateNewDatabase(newDb string, config DBConfig) CreateDBResponse {
	response := CreateDBResponse{
		Error:     nil,
		DBCreated: false,
		Message:   "",
	}

	db, err := connection(config)
	if err != nil {
		response.Error = err
		response.Message = "Failed to create database"
		return response
	}

	// Create the new db
	createDBStatement := fmt.Sprintf("CREATE DATABASE %s;", newDb)
	_, err = db.Exec(createDBStatement)
	if err != nil {
		response.Error = err
		response.Message = "Failed to create database"
		return response
	}

	response.Message = fmt.Sprintf("Database: '%s' was created successfully!", newDb)
	response.DBCreated = true
	return response
}

func Queries(query string, config DBConfig) QueryResponse {
	if !isQueryExpectedToReturnRows(query) {
		return executeQuery(query, config)
	}

	// columns := getDataFromDB(getTheTableName(query), config)
	data := getDataFromDB(query, config)
	return data
}
