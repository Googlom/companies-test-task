package routes

import (
	"companies-test-task/internal/httpserver"
	"fmt"
	"log"
	"net/http"

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
		comp, err := srv.GetCompany(c.Param("id"))
		if err != nil {
			fmt.Println(err)
			return // TODO: log me
		}

		c.JSON(http.StatusOK, comp)
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
