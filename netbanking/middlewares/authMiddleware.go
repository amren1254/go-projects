package middleware

import (
	"net/http"

	"netbanking/auth"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func ExtractUsernameFromTokenClaims(c *gin.Context) (string, error) {
	token := c.Query("token")
	if token != "" {
		return token, nil
	}
	//c := u.(*jwt.Token).Claims.(jwt.MapClaims)
	return "", nil
}
