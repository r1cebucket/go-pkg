package db_test

import (
	"testing"

	"github.com/r1cebucket/gopkg/db"
)

type UserORM struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique"`
	passwd   string `gorm:"not null"`
}

func TestSetupEngin(t *testing.T) {
	db.SetupEngin()
}

func TestAutoMigrate(t *testing.T) {
	db.SetupEngin()
	db.Session().AutoMigrate(&UserORM{})
}
