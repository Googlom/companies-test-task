package httpserver

import (
	"companies-test-task/internal/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server interface {
	service.Companies
}

// httpServer manages the internal state of Server
type httpServer struct {
	service.Companies

	cfg Config
}

func New(cfg Config, svc service.Companies) (*httpServer, error) {
	srv := new(httpServer)

	err := validateCfg(cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid http server configuration: %w", err)
	}
	srv.cfg = cfg
	srv.Companies = svc

	return srv, nil
}

func (srv *httpServer) Start(binder func(Config, Server, *gin.Engine)) error {
	eng := gin.New()
	binder(srv.cfg, srv, eng)

	err := eng.Run(":" + strconv.Itoa(srv.cfg.Port))
	if err != nil {
		return err
	}

	return nil
}
