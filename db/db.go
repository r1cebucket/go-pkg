package db

import (
	"database/sql"
	"fmt"

	_ "github.com/bmizerany/pq" // driver
	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
	"github.com/r1cebucket/gopkg/utils"
)

// manage connection object
var sqlDB *sql.DB

func DBConn() *sql.DB {
	if db != nil {
		sqlDB, _ = db.DB()
	}
	return sqlDB
}

// setup using raw SQL
func Setup() {
	driverName := config.Database.Driver
	source := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.TimeZone)

	conn, err := sql.Open(driverName, source)
	sqlDB = conn
	sqlDB.SetMaxOpenConns(1000)
	if err != nil {
		log.Fatal().Msg("Cannot open postgres DB: " + err.Error())
	}
	err = sqlDB.Ping()
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

	return utils.RowsToMap(rows)
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
