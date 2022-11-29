package main

import (
	"companies-test-task/internal/db"
	"companies-test-task/internal/httpserver"
	"companies-test-task/internal/httpserver/auth"
	"companies-test-task/internal/httpserver/errors"
	"companies-test-task/internal/httpserver/routes"
	"companies-test-task/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := loadConfiguration()
	if err != nil {
		log.Fatalf("configration failed: %s", err)
	}

	storage, err := db.New(cfg.Database)
	if err != nil {
		log.Fatalf("database initialization failed: %s\n", err)
	}

	svc := service.New(cfg.Service, storage)
	httpServer := httpserver.New(cfg.HttpServer, svc)

	err = httpServer.Start(routeBinder)
	if err != nil {
		log.Fatalf("server closed unexpectedly: %s", err)
	}
	os.Exit(0)
}

func routeBinder(cfg httpserver.Config, srv httpserver.Server, eng *gin.Engine) {
	eng.Use(errors.Middleware(), gin.Recovery())
	eng.GET("/companies/:id", routes.GetCompany(srv))

	authMw := auth.JwtMiddleware([]byte(cfg.HmacSecret))
	eng.Use(authMw).POST("/companies", routes.CreateCompany(srv))
	eng.Use(authMw).PATCH("/companies/:id", routes.EditCompany(srv))
	eng.Use(authMw).DELETE("/companies/:id", routes.DeleteCompany(srv))
}

// TODO: handle invalid company type on create/update
