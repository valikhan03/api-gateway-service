package command_service

import(
	"github.com/gin-gonic/gin"
	
	"api-gateway-service/pkg/command-service/routes"
)


func RegisterRoutes(r *gin.Engine) *Service {
	service := &Service{
		Client: InitServiceClient(),
	}

	auctions := r.Group("/auctions")
	auctions.POST("/", service.CreateAuction)
	auctions.PUT("/:id", service.UpdateAuction)
	auctions.DELETE("/:id", service.DeleteAuction)

	auctions.POST("/:id/participants", service.AddParticipant)
	auctions.DELETE("/:id/participants/:participant_id", service.DeleteParticipant)
	
	auctions.POST("/:id/products", service.AddProduct)
	auctions.PUT("/:id/products/:product_id", service.UpdateProduct)
	auctions.DELETE("/:id/products/:product_id", service.DeleteProduct)

	return service
}



func (s *Service) CreateAuction(c *gin.Context) {
	routes.CreateAuction(c, s.Client)
}


func (s *Service) UpdateAuction(c *gin.Context) {
	routes.UpdateAuction(c, s.Client)
}

func (s *Service) DeleteAuction(c *gin.Context) {
	routes.DeleteAuction(c, s.Client)
}	

func (s *Service) AddParticipant(c *gin.Context) {
	routes.AddParticipant(c, s.Client)
}

func (s *Service) DeleteParticipant(c *gin.Context) {
	routes.DeleteParticipant(c, s.Client)
}

func (s *Service) AddProduct(c *gin.Context) {
	routes.AddProduct(c, s.Client)
}

func (s *Service) UpdateProduct(c *gin.Context) {
	routes.UpdateProduct(c, s.Client)
}

func (s *Service) DeleteProduct(c *gin.Context) {
	routes.DeleteProduct(c, s.Client)
}