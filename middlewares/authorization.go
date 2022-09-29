package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/herokh/golang-api/configs"
)

func AuthorizationMiddleware(c *gin.Context, config *configs.AppConfiguration) {
	auth := c.Request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(auth, "Bearer ")

	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.AuthSecret), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
