package httpserver

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/log"
)

type httpServer struct {
	server http.Server
}

type HttpServer interface {
	Start() error
}

func NewHTTPServer(handler http.Handler) HttpServer {
	svr := &httpServer{
		server: http.Server{
			Handler: handler,
			Addr:    config.HTTPServer.Addr,
		},
	}
	return svr
}

func (svr *httpServer) Start() error {
	if svr == nil {
		return errors.New("nil httpServer")
	}
	return svr.server.ListenAndServe()
}

func NewHandler() http.Handler {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)
	w := log.GetWriter()
	if w != nil {
		gin.DefaultErrorWriter = w
		gin.DefaultWriter = w
	}
	engin := gin.Default()
	return engin
}
