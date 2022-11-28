package main

import (
	"companies-test-task/internal/httpserver"
	"companies-test-task/internal/httpserver/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := httpserver.Config{Port: 8080} // TODO:

	srv, err := httpserver.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize server: %s\n", err)
	}

	err = srv.Start(httpBinder)
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
