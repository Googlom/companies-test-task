package errors

import "github.com/gin-gonic/gin"

// APIError struct
type APIError struct {
	Code    int
	Message string
}

// RespondWithError function handler error
func RespondWithError(code int, message string, c *gin.Context) {
	c.JSON(code, &APIError{Code: code, Message: message})
	c.Abort()
}
