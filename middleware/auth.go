package middleware

import (
	"log"
	"net/http"
	"strings"

	"authentication/helpers"

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
		authHeader = strings.TrimPrefix(authHeader, "Bearer ")

		_,err := helpers.ValidateToken(authHeader)

		if err != nil{
			log.Printf("token validation error %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{
	"error": "invalid token",
})
			c.Abort()
			return
		}

	}
}