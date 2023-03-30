package httpserver_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/httpserver"
)

func init() {
	config.Parse("../../configs/conf.json")
}

func TestStart(t *testing.T) {
	engin := httpserver.NewHandler().(*gin.Engine)
	engin.GET("/")
	httpserver.NewHTTPServer(engin).Start()
}
