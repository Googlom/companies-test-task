package httpserver

import (
	"companies-test-task/internal/service"

	"github.com/gin-gonic/gin"
)

type Server interface {
	service.Companies
}

type Config struct {
	HmacSecret string
	ListenAddr string
}

// httpServer manages the internal state of Server
type httpServer struct {
	service.Companies

	cfg Config
}

func New(cfg Config, svc service.Companies) *httpServer {
	srv := new(httpServer)
	srv.cfg = cfg
	srv.Companies = svc
	return srv
}

func (srv *httpServer) Start(binder func(Config, Server, *gin.Engine)) error {
	eng := gin.New()
	binder(srv.cfg, srv, eng)

	return eng.Run(srv.cfg.ListenAddr)
}
