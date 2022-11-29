package errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// APIError struct
type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func RespondWithError(code int, message string, c *gin.Context) {
	c.JSON(code, &APIError{Code: code, Message: message})
	c.Abort()
}
