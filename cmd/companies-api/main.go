package main

import (
	"companies-test-task/internal/db"
	"companies-test-task/internal/httpserver"
	"companies-test-task/internal/httpserver/auth"
	"companies-test-task/internal/httpserver/routes"
	"companies-test-task/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	serverConfig := httpserver.Config{ // TODO: move to env
		HmacSecret: "233972089023bb1838ae877063b3080c7a4fefd57a1a8125a5ff29546b0ea1f1",
		Port:       8080,
	}
	serviceConfig := service.Config{} // TODO: move to env
	dbConfig := db.Config{            // TODO: move to env
		Host:           "localhost",
		Port:           5432,
		User:           "postgres",
		Password:       "12345",
		DbName:         "companies",
		MigrationsPath: "db_migration",
	}

	//////////////////////////////////////////////
	storage, err := db.New(dbConfig)
	if err != nil {
		log.Fatalf("database initialization failed: %s\n", err)
	}

	svc, err := service.New(serviceConfig, storage)
	if err != nil {
		log.Fatalf("service initialization failed: %s\n", err)
	}

	httpServer, err := httpserver.New(serverConfig, svc)
	if err != nil {
		log.Fatalf("failed to initialize server: %s\n", err)
	}

	err = httpServer.Start(routeBinder)
	if err != nil {
		log.Fatalf("server closed unexpectedly: %s", err)
	}
	os.Exit(0)
}

func routeBinder(cfg httpserver.Config, srv httpserver.Server, r *gin.Engine) {
	r.GET("/companies/:id", routes.GetCompany(srv))

	authMw := auth.JwtAuthMiddleware([]byte(cfg.HmacSecret))
	r.Use(authMw).POST("/companies", routes.CreateCompany(srv))
	r.Use(authMw).PATCH("/companies/:id", routes.EditCompany(srv))
	r.Use(authMw).DELETE("/companies/:id", routes.DeleteCompany(srv))
}

// TODO: handle invalid company type on create/update
