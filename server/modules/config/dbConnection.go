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

type DbConfig struct {
	Host     string
	User     string
	Password string
	Port     int
}

func (db DbConfig) connection(dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, dbName,
	)
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = database.Ping()

	if err != nil {
		database.Close()
		return nil, err
	}

	return database, nil
}

func (db DbConfig) CreateDataBase(newDb string) string {
	database, err := db.connection("postgres")
	if err != nil {
		log.Fatal("Failed to create database")
	}

	// Create the new database
	createDBStatement := fmt.Sprintf("CREATE DATABASE %s", newDb)
	_, err = database.Exec(createDBStatement)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Database '%s' created successfully!\n", newDb)
}

func (db DbConfig) Querie(query string, dbName string) []User {
	database, err := db.connection(dbName)
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	rows, err := database.Query(query)
	database.QueryRow(query).Scan()
	if err != nil {
		log.Fatal("Failed to execute the query:", err)
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			log.Println("Failed to scan row:", err)
		}
		users = append(users, user)
	}
	return users
}
