// docs: https://gorm.io/docs/
// tags: https://gorm.io/zh_CN/docs/models.html#%E5%AD%97%E6%AE%B5%E6%A0%87%E7%AD%BE
package db

import (
	"fmt"
	"log"
	"time"

	"github.com/r1cebucket/gopkg/config"
	zerolog "github.com/r1cebucket/gopkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Session() *gorm.DB {
	if db == nil {
		zerolog.Fatal().Msg("no session found, run db.SetupSession() first")
	}
	return db
}

// setup engin of gorm
func SetupSession() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.TimeZone,
	)
	session, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: logger.New(
				log.New(zerolog.GetWriter(), "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             200 * time.Millisecond,
					LogLevel:                  3,
					IgnoreRecordNotFoundError: false,
					Colorful:                  true,
				},
			),
			// Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		zerolog.Fatal().Msg("Could not open database")
		return err
	}
	db = session

	// sqlDB := DBConn()
	// if err != nil {
	// 	log.Fatal().Msg("Could not get sqlDB")
	// }
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}
