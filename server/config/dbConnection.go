package config

import (
	"database/sql"
	"fmt"
	"log"

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

func Queries(query string, config DBConfig) (interface{}, error) {
	if !isQueryExpectedToReturnRows(query) {
		db, err := connection(config)
		// TODO: those errors instead of printing it in the terminal, send to the user
		// as a http response
		if err != nil {
			return "", fmt.Errorf("failed to connect to the database: %w", err)
		}
		defer db.Close()

		result, err := db.Exec(query)
		if err != nil {
			return "", fmt.Errorf("failed to execute the query: %w", err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return "", fmt.Errorf("failed to retrieve the number of rows affected: %w", err)
		}

		response := fmt.Sprintf("%s %d", query, rowsAffected)
		return response, nil
	}

	db, err := connection(config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve column names: %w", err)
	}

	data := []map[string]interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			log.Println("Failed to scan row:", err)
			continue
		}

		entry := make(map[string]interface{})
		for i, column := range columns {
			entry[column] = values[i]
		}

		data = append(data, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %w", err)
	}

	return data, nil
}
