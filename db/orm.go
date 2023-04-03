// docs: https://gorm.io/docs/
// tags: https://gorm.io/zh_CN/docs/models.html#%E5%AD%97%E6%AE%B5%E6%A0%87%E7%AD%BE
package db

import (
	"fmt"

	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Session() *gorm.DB {
	if db == nil {
		log.Fatal().Msg("no session found, run db.SetupSession() first")
	}
	return db
}

// setup engin of gorm
func SetupSession() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.TimeZone)
	session, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
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
