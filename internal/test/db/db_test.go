package db_test

import (
	"testing"

	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/db"
	"github.com/r1cebucket/gopkg/log"
)

func init() {
	log.Setup("debug")
	config.Parse("../../configs/conf.json")
	db.Setup()
}

// create table users(
// 	id serial,
// 	username varchar(64) not null unique,
// 	passwd varchar(256) not null
// );
// insert into users(username, passwd) values(testsuer, testpwd);

func TestQuery(t *testing.T) {
	query := `select * from users;`
	_, err := db.Query(db.DBConn(), query)
	if err != nil {
		t.Error()
	}
}

func TestExec(t *testing.T) {
	var exec string
	var err error
	var affected int
	exec = `insert into users(username, passwd) values($1, $2)`
	affected, err = db.Exec(db.DBConn(), exec, "this_is_a_test_username", "this_is_a_test_pwd")
	if err != nil || affected != 1 {
		t.Error()
		return
	}
	exec = `delete from users where username=$1`
	affected, err = db.Exec(db.DBConn(), exec, "this_is_a_test_username")
	if err != nil || affected != 1 {
		t.Error()
		return
	}
}
