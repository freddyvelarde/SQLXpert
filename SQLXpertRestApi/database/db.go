package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/freddyvelarde/SQLXpert/structs"
)

func Connection(config structs.DBConfig) (*sql.DB, error) {
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
