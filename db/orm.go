package db

import (
	"fmt"

	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Session() *gorm.DB {
	return db
}

// setup engin of gorm
func SetupEngin() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.TimeZone)
	session, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("Could not open database")
	}
	db = session

	// sqlDB := DBConn()
	// if err != nil {
	// 	log.Fatal().Msg("Could not get sqlDB")
	// }
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Hour)
}
