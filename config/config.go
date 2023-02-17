package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/r1cebucket/gopkg/log"
)

type configure struct {
	Logger     logger     `json:"logger"`
	Database   database   `json:"database"`
	Redis      redis      `json:"redis"`
	Email      email      `json:"email"`
	HTTPServer httpServer `json:http`
	// TODO add new conf here
}

var conf configure

// user pointer to save space
var Logger *logger
var Database *database
var Redis *redis
var Email *email
var HTTPServer *httpServer

// TODO add new conf var here

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
}

// TODO add new conf struct here

func Parse(confPath string) error {
	if strings.HasSuffix(confPath, ".json") {
		ParseJson(confPath)
		return nil
	} else if strings.HasSuffix(confPath, ".toml") {
		return nil
	} else {
		return errors.New("config type not supported: " + confPath)
	}
}

func ParseJson(confPath string) {
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

	Database = &conf.Database
	Logger = &conf.Logger
	Redis = &conf.Redis
	HTTPServer = &conf.HTTPServer

	// TODO add new conf here
}
