package main

import (
	auctions_srvc "api-gateway-service/pkg/auctions-service"
	auth_srvc "api-gateway-service/pkg/auth-service"
	command_srvc "api-gateway-service/pkg/command-service"
	query_srvc "api-gateway-service/pkg/query-service"
	payment_srvc "api-gateway-service/pkg/payment-service"
	"log"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth_srvc.RegisterRoutes(router)


	api := router.Group("/api/v1")
	auctions_srvc.RegisterRoutes(api)
	command_srvc.RegisterRoutes(api)
	query_srvc.RegisterRoutes(api)
	payment_srvc.RegisterRoutes(api)

	err := router.Run()
	if err != nil {
		log.Fatalf("ROUTER ERROR: %v", err)
	}
}