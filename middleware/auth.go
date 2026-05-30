package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader:=c.GetHeader("Authorization")

		if authHeader ==""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			c.Abort()
			return
		}
		authHeader = strings.TrimPrefix(authHeader, "Bearer")

		claims,err := ValidateToken(authHeader)

		if err != nil{
			log.printf("token validation error %v", err)
			c.JSON(http.StatusUnauthorized.gin.H("error":"invalid token"))
			c.Abort()
			return
		}

	}
}