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

func CreateNewDatabase(newDb string, config DBConfig) string {
	db, err := connection(config)
	if err != nil {
		return "Failed to create db"
	}

	// Create the new db
	createDBStatement := fmt.Sprintf("CREATE DATABASE %s;", newDb)
	_, err = db.Exec(createDBStatement)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Database: '%s' was created successfully!", newDb)
}

func Querie(query string, config DBConfig) []User {
	// TODO: refactor this code, cause this code it was only for test database connection
	db, err := connection(config)
	if err != nil {
		log.Fatal("Failed to connect to the db", err)
	}

	rows, err := db.Query(query)
	db.QueryRow(query).Scan()
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
