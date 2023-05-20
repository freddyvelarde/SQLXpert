package database

import (
	"fmt"
	"strings"
)

type QueryResponse struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}

type ColumnNames struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Columns interface{} `json:"columns"`
}

type DatabaseNames struct {
	Message   string   `json:"message"`
	Error     bool     `json:"error"`
	Databases []string `json:"databases"`
}

func isQueryExpectedToReturnRows(query string) bool {
	lowerQuery := strings.ToLower(query)

	return strings.HasPrefix(lowerQuery, "select")
}

// SELECT column_name::text FROM information_schema.columns WHERE table_schema = 'public'   AND table_name = 'post'
func executeQuery(query string, config DBConfig) QueryResponse {
	response := QueryResponse{
		Message: "",
		Error:   true,
	}

	db, err := Connection(config)
	if err != nil {
		response.Message = fmt.Sprintf("failed to connect to the Database: %s", err)
		return response
	}
	defer db.Close()

	result, err := db.Exec(query)
	if err != nil {
		response.Message = fmt.Sprintf("failed to execute the query: %s", err)
		return response
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		response.Message = fmt.Sprintf("failed to retrieve the number of rows affected: %s", err)
		return response
	}

	response.Message = fmt.Sprintf("%s %d", query, rowsAffected)
	response.Error = false
	return response
}

func getDataFromDB(query string, config DBConfig) QueryResponse {
	response := QueryResponse{
		Message: "",
		Error:   true,
	}

	db, err := Connection(config)
	if err != nil {
		response.Message = fmt.Sprintf("failed to connect to the Database: %s", err)
		return response
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		response.Message = fmt.Sprintf("failed to execute the query: %s", err)
		return response
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		response.Message = fmt.Sprintf("failed to retrieve column names: %s", err)
		return response
	}

	data := []map[string]interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			response.Message = fmt.Sprintf("Failed to scan row: %s", err)
			return response
		}

		entry := make(map[string]interface{})
		for i, column := range columns {
			entry[column] = values[i]
		}

		data = append(data, entry)
	}

	if err := rows.Err(); err != nil {
		response.Message = fmt.Sprintf("failed to iterate over rows: %s", err)
		return response
	}
	response.Data = data
	response.Error = false
	return response
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
