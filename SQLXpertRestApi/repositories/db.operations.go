package repositories

import (
	"fmt"

	"github.com/freddyvelarde/SQLXpert/database"
	"github.com/freddyvelarde/SQLXpert/structs"
	"github.com/freddyvelarde/SQLXpert/utils"
	_ "github.com/lib/pq"
)

func CreateNewDatabase(newDb string, config structs.DBConfig) (interface{}, error) {
	db, err := database.Connection(config)
	if err != nil {
		return nil, err
	}

	createDBStatement := fmt.Sprintf("CREATE DATABASE %s;", newDb)

	_, err = db.Exec(createDBStatement)

	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("Database: '%s' was created successfully!", newDb), nil
}

func Queries(query string, config structs.DBConfig) (interface{}, error) {
	if utils.IsQueryExpectedToReturnRows(query) {
		data, err := utils.GetDataFromDB(query, config)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	res, err := utils.ExecuteQuery(query, config)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetAllDatabases(config structs.DBConfig) (interface{}, error) {
	query := "SELECT datname::text FROM pg_catalog.pg_database WHERE datistemplate = false;"

	data, err := utils.GetDataFromDB(query, config)
	if err != nil {
		return nil, err
	}

	databases := []string{}

	for _, value := range data.([]map[string]interface{}) {
		db, ok := value["datname"].(string)
		if !ok {
			continue
		}
		databases = append(databases, db)
	}
	return databases, nil
}

func GetTableNames(config structs.DBConfig) (interface{}, error) {
	query := "SELECT tablename::text FROM pg_catalog.pg_tables WHERE schemaname = 'public';"

	data, err := utils.GetDataFromDB(query, config)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, value := range data.([]map[string]interface{}) {
		tablename, ok := value["tablename"].(string)
		if !ok {
			continue
		}
		res = append(res, tablename)
	}

	return res, nil
}

func GetAllColumnNamesfromTable(config structs.DBConfig) (interface{}, error) {
	query := "SELECT column_name FROM information_schema.columns WHERE table_name = 'your_table';"

	data, err := utils.GetDataFromDB(query, config)
	if err != nil {
		return nil, err
	}

	columns := []string{}

	for _, value := range data.([]map[string]interface{}) {
		column, ok := value["tablename"].(string)
		if !ok {
			continue
		}
		columns = append(columns, column)
	}
	return columns, nil
}
