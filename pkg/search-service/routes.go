package search

import (
	"api-gateway-service/pkg/search-service/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(c *gin.Engine) {
	service := &Service{
		Client: InitServiceClient(),
	}

	c.GET("/auctions/:id", service.GetAuction)
	c.GET("/products/:id", service.GetProduct)

	c.POST("/search", service.Search)
}

func (s *Service) GetAuction(c *gin.Context) {
	routes.GetAuction(c, s.Client)
}

func (s *Service) GetProduct(c *gin.Context) {
	routes.GetProduct(c, s.Client)
}

func (s *Service) Search(c *gin.Context) {
	routes.Search(c, s.Client)
}
