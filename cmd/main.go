package main

import (
	auctions_srvc "api-gateway-service/pkg/auctions-service"
	auth_srvc "api-gateway-service/pkg/auth-service"
	command_srvc "api-gateway-service/pkg/command-service"
	search_srvc "api-gateway-service/pkg/search-service"
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	auctions_srvc.RegisterRoutes(router)
	auth_srvc.RegisterRoutes(router)
	command_srvc.RegisterRoutes(router)
	search_srvc.RegisterRoutes(router)


	err := router.Run()
	if err != nil {
		log.Fatalf("ROUTER ERROR: %x", err)
	}
}