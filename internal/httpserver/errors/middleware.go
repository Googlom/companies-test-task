package errors

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		var apiErr *APIError
		if errors.As(err.Err, &apiErr) {
			RespondWithError(apiErr.Code, apiErr.Message, c)
		} else if err.IsType(gin.ErrorTypeBind) { // Gin struct binding validation errors
			RespondWithError(http.StatusBadRequest, err.Error(), c)
		} else if errors.Is(err.Err, pgx.ErrNoRows) {
			RespondWithError(http.StatusNotFound, "Record not found", c)
		} else if strings.Contains(err.Error(), "duplicate key") {
			RespondWithError(http.StatusConflict, "Duplicate field value", c)
		} else {
			RespondWithError(http.StatusInternalServerError, "Internal server error", c)
		}
		
		log.Println(err.Error())
	}
}
