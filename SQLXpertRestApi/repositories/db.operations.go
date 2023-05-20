package repositories

import (
	"fmt"

	"github.com/freddyvelarde/SQLXpert/database"
	"github.com/freddyvelarde/SQLXpert/utils"
	_ "github.com/lib/pq"
)

func CreateNewDatabase(newDb string, config utils.DBConfig) (string, error) {
	db, err := database.Connection(config)
	if err != nil {
		return "", nil
	}

	createDBStatement := fmt.Sprintf("CREATE DATABASE %s;", newDb)

	_, err = db.Exec(createDBStatement)

	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("Database: '%s' was created successfully!", newDb), nil
}

func Queries(query string, config utils.DBConfig) QueryResponse {
	if isQueryExpectedToReturnRows(query) {
		data := getDataFromDB(query, config)
		return data
	}

	return executeQuery(query, config)
}

func GetAllDatabases(config utils.DBConfig) DatabaseNames {
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

func GetAllColumnNamesfromTable(config utils.DBConfig) ColumnNames {
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
