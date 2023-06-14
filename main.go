package main

import (
	auctions_service "api-gateway-service/pkg/auctions-service"
	auth_service "api-gateway-service/pkg/auth-service"
	command_service "api-gateway-service/pkg/command-service"
	query_service "api-gateway-service/pkg/query-service"
	payment_service "api-gateway-service/pkg/payment-service"
	"log"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"api-gateway-service/pkg/models"
	"api-gateway-service/pkg/redis"

)


func main() {
	models.InitConfigs()
	redis.InitRedisConn()
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authService := auth_service.RegisterRoutes(router)

	api := router.Group("/api/v1")
	api.Use(auth_service.InitAuthMiddleware(authService).AuthRequired)

	auctions_service.RegisterRoutes(api)
	command_service.RegisterRoutes(api)
	query_service.RegisterRoutes(api)
	payment_service.RegisterRoutes(api)

	err := router.Run()
	if err != nil {
		log.Fatalf("ROUTER ERROR: %v", err)
	}
}