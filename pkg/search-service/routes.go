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

	routes := c.Group("search")
	routes.POST("")
}


func (s *Service) GetAuction(c *gin.Context) {
	routes.GetAuction(c, s.Client)
}

func (s *Service) GetProduct(c *gin.Context) {
	routes.GetProduct(c, s.Client)
}

func (s *Service) SearchAuctions(c *gin.Context) {
	routes.SearchAuctions(c, s.Client)
}

func (s *Service) SearchProducts(c *gin.Context) {
	routes.SearchProducts(c, s.Client)
}