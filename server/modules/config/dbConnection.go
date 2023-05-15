package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	Name string `json:"name"`
}

type DbConfig struct {
	Host     string
	User     string
	Dbname   string
	Password string
	Port     int
}

func (db DbConfig) connection() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s dbname=%s user=%s port=%d password=%s sslmode=disable",
		db.Host, db.Dbname, db.User, db.Port, db.Password,
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

func (db DbConfig) Querie(query string) []User {
	database, err := db.connection()
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	rows, err := database.Query(query)
	if err != nil {
		log.Fatal("Failed to execute the query:", err)
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Name); err != nil {
			log.Println("Failed to scan row:", err)
		}
		users = append(users, user)
	}
	return users
}
