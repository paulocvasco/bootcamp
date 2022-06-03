package authentication

import (
	customErrors "bootcamp/aula_11/go-web/internal/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var token string

func genToken() {
	token = uuid.NewString()
}

func GetToken() string {
	if token == "" {
		genToken()
	}
	return token
}

func ValidateToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("token_authentication")
		if authToken == "" || authToken != token {
			c.AbortWithError(http.StatusUnauthorized, customErrors.ErrorInvalidToken)
			return
		}
		c.Next()
	}
}
