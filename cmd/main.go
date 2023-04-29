package main

import (
	auctions_srvc "api-gateway-service/pkg/auctions-service"
	auth_srvc "api-gateway-service/pkg/auth-service"
	command_srvc "api-gateway-service/pkg/command-service"
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	auth_srvc.RegisterRoutes(router)


	services := router.Group("services")
	auctions_srvc.RegisterRoutes(services)
	command_srvc.RegisterRoutes(services)


	err := router.Run()
	if err != nil {
		log.Fatalf("ROUTER ERROR: %x", err)
	}
}