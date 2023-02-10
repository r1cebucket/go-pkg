package pkg_test

import (
	"testing"

	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
)

func init() {
	log.Setup("debug")
}

func TestParse(t *testing.T) {
	var err error
	err = config.Parse("../configs/conf.json")
	if err != nil {
		log.Err(err)
		t.Error()
	}
	err = config.Parse("../configs/conf.somethingelse")
	log.Info().Msg(err.Error())
	if err == nil {
		t.Error()
	}
}
