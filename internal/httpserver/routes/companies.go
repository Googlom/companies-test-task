package routes

import (
	"companies-test-task/internal/httpserver"
	"companies-test-task/pkg/dto"
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
		var dtoComp dto.Company
		if err := c.BindJSON(&dtoComp); err != nil {
			return
		}

		resultComp, err := srv.CreateCompany(&dtoComp)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusCreated, resultComp)
	}
}

func GetCompany(srv httpserver.Server) gin.HandlerFunc {
	if srv == nil {
		log.Fatalln("a nil srv was passed to GetCompany handler")
	}

	return func(c *gin.Context) {
		comp, err := srv.GetCompany(c.Param("id"))
		if err != nil {
			_ = c.Error(err)
			return
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
		err := srv.DeleteCompany(c.Param("id"))
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.Status(http.StatusOK)
	}
}
