package routes

import (
	"companies-test-task/internal/httpserver"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateCompany(srv httpserver.Server) gin.HandlerFunc {
	if srv == nil {
		log.Fatalln("a nil srv was passed to CreateCompany handler")
	}

	return func(c *gin.Context) {
		// TODO
		err := srv.CreateCompany(nil)
		if err != nil {
			return
		}

		fmt.Println("POST")
	}
}

func GetCompany(srv httpserver.Server) gin.HandlerFunc {
	if srv == nil {
		log.Fatalln("a nil srv was passed to GetCompany handler")
	}

	return func(c *gin.Context) {
		// TODO
		_, err := srv.GetCompany("")
		if err != nil {
			return
		}

		fmt.Println("GET")
	}
}

func EditCompany(srv httpserver.Server) gin.HandlerFunc {
	if srv == nil {
		log.Fatalln("a nil srv was passed to EditCompany handler")
	}

	return func(c *gin.Context) {
		// TODO
		err := srv.EditCompany()
		if err != nil {
			return
		}

		fmt.Println("PATCH")
	}
}

func DeleteCompany(srv httpserver.Server) gin.HandlerFunc {
	if srv == nil {
		log.Fatalln("a nil srv was passed to DeleteCompany handler")
	}

	return func(c *gin.Context) {
		// TODO
		err := srv.DeleteCompany("")
		if err != nil {
			return
		}

		fmt.Println("DELETE")
	}
}
