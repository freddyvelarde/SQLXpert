package database

import (
	"database/sql"
	"fmt"
	"strings"

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

func Connection(config DBConfig) (*sql.DB, error) {
	if strings.ToLower(config.Host) == "localhost" {
		config.Host = "172.19.0.1"
	}

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

	db, err := Connection(config)
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
	if isQueryExpectedToReturnRows(query) {
		data := getDataFromDB(query, config)
		return data
	}

	return executeQuery(query, config)
}

func GetAllDatabases(config DBConfig) DatabaseNames {
	query := "SELECT datname::text FROM pg_catalog.pg_database WHERE datistemplate = false;"

	data := getDataFromDB(query, config)

	databases := []string{}

	for _, value := range data.Data.([]map[string]interface{}) {
		db, ok := value["datname"].(string)
		if !ok {
			continue
		}
		databases = append(databases, db)
	}
	return DatabaseNames{
		Error:     false,
		Databases: databases,
	}
}

func GetAllColumnNamesfromTable(config DBConfig) ColumnNames {
	query := "SELECT column_name FROM information_schema.columns WHERE table_name = 'your_table';"

	data := getDataFromDB(query, config)

	columns := []string{}

	for _, value := range data.Data.([]map[string]interface{}) {
		column, ok := value["tablename"].(string)
		if !ok {
			continue
		}
		columns = append(columns, column)
	}
	return ColumnNames{
		Error:   false,
		Columns: columns,
	}
}
