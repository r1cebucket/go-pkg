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
	Level string `json:"level" toml:"level"`
}
type database struct {
	Driver   string `json:"driver" toml:"driver"`
	Host     string `json:"host" toml:"host"`
	Port     string `json:"port" toml:"port"`
	User     string `json:"user" toml:"user"`
	Password string `json:"password" toml:"password"`
	DBName   string `json:"dbname" toml:"dbname"`
	TimeZone string `json:"timezone" toml:"timezone"`
}

type influxdb struct {
	Token  string `json:"token" toml:"token"`
	URL    string `json:"url" toml:"url"`
	Org    string `json:"org" toml:"org"`
	Bucket string `json:"bucket" toml:"bucket"`
}

type redis struct {
	Host     string `json:"host" toml:"host"`
	Port     string `json:"port" toml:"port"`
	Password string `json:"password" toml:"password"`
}

type email struct {
}

type httpServer struct {
	Addr string `json:"addr" toml:"addr"`
	Mode string `json:"mode" toml:"mode"`
}

type kafka struct {
	Servers []string `json:"servers" toml:"servers"`
}

// TODO add new conf struct here

type configure struct {
	Logger     logger     `json:"logger" toml:"logger"`
	Database   database   `json:"database" toml:"database"`
	Influxdb   influxdb   `json:"influxdb" toml:"influxdb"`
	Redis      redis      `json:"redis" toml:"redis"`
	Email      email      `json:"email" toml:"email"`
	HTTPServer httpServer `json:"http" toml:"http"`
	Kafka      kafka      `json:"kafka" toml:"kafka"`
	// TODO add new conf here
}

var conf configure

// use pointer to save space
var Logger *logger
var Database *database
var Influxdb *influxdb
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
	Logger = &conf.Logger
	Database = &conf.Database
	Influxdb = &conf.Influxdb
	Redis = &conf.Redis
	HTTPServer = &conf.HTTPServer
	Kafka = &conf.Kafka

	// TODO add new conf here
}
