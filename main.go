package main

import (
	"authentication/config"
	"authentication/helpers"
	"authentication/routes"
	"log"

	"github.com/gin-gonic/gin"
)




func main() {
	config.ConnectDb()

	port := "8080"

	key := config.GenerateRandomKey()
	helpers.SetJwtKey(key)

	r := gin.Default()
	routes.SetupRoutes(r)

	log.Println("server is running on port", port)

	r.Run(":" + port)
}