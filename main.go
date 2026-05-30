package authenticationgo

import (
	"authentication/routes"
	"fmt"
	"log"
)
package main

import "github.com/gin-gonic/gin"

func main(){
	key:=GenerateRandomKey()
	SetJwtKey(key)

	r:=gin.Default()
	routes.SetupRoutes(r)
	r.Run(":"+port)
	log.Println("server is running in", port)
}