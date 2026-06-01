package main

import (
	"authentication/config"
	"authentication/helpers"
	"authentication/routes"
	"log"

	"github.com/gin-gonic/gin"
)




func main(){
	port := "8080"
	key:=config.GenerateRandomKey()
	helpers.SetJwtKey(key)

	r:=gin.Default()
	routes.SetupRoutes(r)
	r.Run(":"+port)
	log.Println("server is running in", port)
}