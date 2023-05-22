package utils

import (
	"fmt"
	"strings"

	"github.com/freddyvelarde/SQLXpert/database"
	"github.com/freddyvelarde/SQLXpert/structs"
)

func IsQueryExpectedToReturnRows(query string) bool {
	lowerQuery := strings.ToLower(query)

	return strings.HasPrefix(lowerQuery, "select")
}

func ExecuteQuery(query string, config structs.DBConfig) (interface{}, error) {
	db, err := database.Connection(config)
	if err != nil {
		return fmt.Sprintf("failed to connect to the Database: %s", err), err
	}
	defer db.Close()

	result, err := db.Exec(query)
	if err != nil {
		return fmt.Sprintf("failed to execute the query: %s", err), err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Sprintf("failed to retrieve the number of rows affected: %s", err), err
	}

	return fmt.Sprintf("%s %d", query, rowsAffected), nil
}

func GetDataFromDB(query string, config structs.DBConfig) (interface{}, error) {
	db, err := database.Connection(config)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	data := []map[string]interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return fmt.Sprintf("Failed to scan row: %s", err), err
		}

		entry := make(map[string]interface{})
		for i, column := range columns {
			entry[column] = values[i]
		}

		data = append(data, entry)
	}

	if err := rows.Err(); err != nil {
		return fmt.Sprintf("failed to iterate over rows: %s", err), nil
	}
	return data, nil
}

func GetColumnNamesFromTable(query string) string {
	words := strings.Fields(strings.ToLower(query))

	var tableName string
	for i, word := range words {
		if word == "from" && i+1 < len(words) {
			tableName = words[i+1]
			break
		}
	}

	var res strings.Builder
	for _, char := range tableName {
		if char != '"' && char != ';' && char != '\'' {
			res.WriteRune(char)
		}
	}

	return fmt.Sprintf("SELECT column_name::text FROM information_schema.columns WHERE table_schema = 'public' AND table_name = '%s'", res.String())
}
