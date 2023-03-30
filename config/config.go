package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/koding/multiconfig"
	"github.com/r1cebucket/gopkg/log"
)

type logger struct {
	Level string `json:"level"`
}
type database struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type email struct {
}

type httpServer struct {
	Addr string `json:"addr"`
	Mode string `json:"mode"`
}

type kafka struct {
	Servers []string `json:"servers"`
}

// TODO add new conf struct here

type configure struct {
	Logger     logger     `json:"logger"`
	Database   database   `json:"database"`
	Redis      redis      `json:"redis"`
	Email      email      `json:"email"`
	HTTPServer httpServer `json:"http"`
	Kafka      kafka      `json:"kafka"`
	// TODO add new conf here
}

var conf configure

// user pointer to save space
var Logger *logger
var Database *database
var Redis *redis
var Email *email
var HTTPServer *httpServer
var Kafka *kafka

// TODO add new conf var here

func Parse(confPath string) error {
	if strings.HasSuffix(confPath, ".json") {
		parseJson(confPath)
		return nil
	} else if strings.HasSuffix(confPath, ".toml") {
		parseToml(confPath)
		return nil
	} else {
		return errors.New("config type not supported: " + confPath)
	}
}

func parseJson(confPath string) {
	// open and read config file
	confFile, err := os.Open(confPath)
	if err != nil {
		log.Fatal().Msg("Cannot open config file:" + err.Error())
		return
	}
	defer confFile.Close()
	data, err := io.ReadAll(confFile)
	if err != nil {
		log.Fatal().Msg("Cannot read config file:" + err.Error())
		return
	}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		return
	}

	register()
}

func parseToml(confPath string) {
	m := multiconfig.NewWithPath(confPath)
	err := m.Load(&conf)
	if err != nil {
		log.Fatal().Msg("Faile ot " + err.Error())
		return
	}

	register()
}

func register() {
	Database = &conf.Database
	Logger = &conf.Logger
	Redis = &conf.Redis
	HTTPServer = &conf.HTTPServer
	Kafka = &conf.Kafka

	// TODO add new conf here
}
