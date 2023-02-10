package db

import (
	"database/sql"
	"fmt"

	_ "github.com/bmizerany/pq" // driver
	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
)

// manage connection object
var db *sql.DB

func DBConn() *sql.DB {
	return db
}

func Setup() {
	driverName := config.Database.Type
	source := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName)

	conn, err := sql.Open(driverName, source)
	db = conn
	db.SetMaxOpenConns(1000)
	if err != nil {
		log.Fatal().Msg("Cannot open postgres DB: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal().Msg("Cannot connect to postgres DB: " + err.Error())
	}
}

func Query(db *sql.DB, query string, params ...interface{}) ([]map[string]interface{}, error) {
	// create stmt to do the query
	stmt, err := DBConn().Prepare(query)
	if err != nil {
		log.Err(err).Msg("Faile to create stat")
		return nil, err
	}
	defer stmt.Close()
	// do query
	rows, err := stmt.Query(params...)
	if err != nil {
		log.Err(err).Msg("Faile to do query")
	}
	defer rows.Close()
	// read field name
	columns, err := rows.Columns()
	if err != nil {
		log.Err(err).Msg("Faile to read field name")
		return nil, err
	}
	// write values to map
	data := []map[string]interface{}{}
	row := make([]interface{}, len(columns))
	for i := range row {
		// convert to *interface
		var tmp interface{}
		row[i] = &tmp
	}
	for rows.Next() {
		err = rows.Scan(row...)
		if err != nil {
			log.Err(err).Msg("Faile to read row")
			return nil, err
		}

		item := make(map[string]interface{})
		for i, val := range row {
			// val is interface{} which store *interface{} value
			// convert val to *interface{} and then get the value
			item[columns[i]] = *(val.(*interface{}))
		}
		data = append(data, item)
	}
	return data, nil
}

// return rows affected
func Exec(db *sql.DB, exec string, params ...interface{}) (int, error) {
	// create stmt
	stmt, err := DBConn().Prepare(exec)
	if err != nil {
		log.Err(err).Msg("Faile to create stat")
		return 0, err
	}
	defer stmt.Close()
	// exec
	result, err := stmt.Exec(params...)
	if err != nil {
		log.Err(err).Msg("Faile to exec")
		return 0, err
	}
	affected, _ := result.RowsAffected()

	return int(affected), nil
}
