package kafka

import (
	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
)

func init() {
	config.Parse("../../configs/conf.json")
	log.Setup("debug")
}
