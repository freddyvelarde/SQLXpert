package repositories

import (
	"fmt"
	"strings"

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
			return nil, fmt.Errorf("table name missing")
		}
		res = append(res, tablename)
	}

	return res, nil
}

func GetAllColumnNamesfromTable(config structs.DBConfig, query string) (interface{}, error) {
	querySplited := strings.Split(strings.ToLower(query), " ")

	var table string
	for i, word := range querySplited {
		if word == "from" {
			tb := querySplited[i+1]
			if tb[len(tb)-1] == ';' {
				table = tb[:len(tb)-1]
			} else {
				table = tb
			}
			break
		}
	}

	q := "SELECT column_name::text FROM information_schema.columns WHERE table_name = '" + table + "' AND table_schema = 'public';"

	data, err := utils.GetDataFromDB(q, config)
	if err != nil {
		return nil, err
	}

	columns := []string{}

	for _, value := range data.([]map[string]interface{}) {
		column, ok := value["column_name"].(string)
		if !ok {
			return nil, fmt.Errorf("column name missing")
		}
		columns = append(columns, column)
	}
	return columns, nil
}
