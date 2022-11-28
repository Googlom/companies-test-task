package main

import (
	"companies-test-task/internal/db"
	"companies-test-task/internal/httpserver"
	"companies-test-task/internal/httpserver/routes"
	"companies-test-task/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	serverConfig := httpserver.Config{Port: 8080} // TODO:
	serviceConfig := service.Config{}             // TODO:
	dbConfig := db.Config{                        // TODO:
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

	err = httpServer.Start(httpBinder)
	if err != nil {
		log.Fatalf("server closed unexpectedly: %s", err)
	}
	os.Exit(0)
}

func httpBinder(srv httpserver.Server, r *gin.Engine) {
	r.POST("/companies", routes.CreateCompany(srv))
	r.GET("/companies/:id", routes.GetCompany(srv))
	r.PATCH("/companies/:id", routes.EditCompany(srv))
	r.DELETE("/companies/:id", routes.DeleteCompany(srv))
}

// TODO: handle invalid company type on create/update
