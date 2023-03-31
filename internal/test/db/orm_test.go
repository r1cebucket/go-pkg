package db_test

import (
	"fmt"
	"testing"

	"github.com/r1cebucket/gopkg/db"
	"github.com/r1cebucket/gopkg/log"
	"gorm.io/gorm"
)

func init() {
	db.SetupSession()
}

type UserORM struct {
	gorm.Model
	Username string `gorm:"unique"`
	Passwd   string `gorm:"not null"`
}

func TestSetupEngin(t *testing.T) {
	db.SetupSession()
}

func TestAutoMigrate(t *testing.T) {
	err := db.Session().AutoMigrate(&UserORM{})
	if err != nil {
		log.Err(err).Msg("migration faild")
		t.Error()
	}
}

func TestInsert(t *testing.T) {
	err := db.Session().Create(&UserORM{
		Username: "testuser3",
		Passwd:   "testuser",
	}).Error
	if err != nil {
		log.Err(err).Msg("insert faild")
		t.Error()
	}
}

func TestSelect(t *testing.T) {
	user := UserORM{}
	// db.Session().Take(&user, "id=?", 2)
	// db.Session().Where("id=?", 2).Take(&user)
	// db.Session().Take(&user, "username=?", "testuser2")

	log.Info().Msg(fmt.Sprint(user))
}
