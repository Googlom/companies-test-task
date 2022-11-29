package routes

import (
	"companies-test-task/internal/httpserver"
	"companies-test-task/internal/httpserver/errors"
	"companies-test-task/pkg/dto"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompany(srv httpserver.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dtoComp dto.Company
		if err := c.BindJSON(&dtoComp); err != nil {
			_ = c.Error(err)
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
	return func(c *gin.Context) {
		reqBodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			_ = c.Error(&errors.APIError{Code: http.StatusBadRequest, Message: "invalid request body: " + err.Error()})
			return
		}

		if err = srv.EditCompany(c.Param("id"), reqBodyBytes); err != nil {
			_ = c.Error(err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func DeleteCompany(srv httpserver.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := srv.DeleteCompany(c.Param("id"))
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.Status(http.StatusOK)
	}
}
