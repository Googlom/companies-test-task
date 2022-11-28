package auth

import (
	"companies-test-task/internal/httpserver/errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthMiddleware(hmacSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errors.RespondWithError(http.StatusUnauthorized, "Authorization required", c)
			return
		}
		parts := strings.Fields(authHeader)

		if len(parts) != 2 {
			errors.RespondWithError(http.StatusUnauthorized, "Authorization header must be 'Bearer {token}'", c)
			return
		}
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSecret, nil
		})
		if err != nil {
			errors.RespondWithError(http.StatusUnauthorized, err.Error(), c)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			errors.RespondWithError(http.StatusUnauthorized, "Token invalid or expired", c)
			return
		}
	}
}
