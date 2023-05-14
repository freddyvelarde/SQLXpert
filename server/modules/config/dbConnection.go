package config

import (
	// "database/sql"
	"fmt"
)

func DbConnection(host, user, password string, port int) string {
	connStr := fmt.Sprintf("host=%s user=%s port=%d password=%s sslmode=require", host, user, port, password)
	return connStr
	// sql.Open()
}
