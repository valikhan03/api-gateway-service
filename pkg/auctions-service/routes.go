package auctions_service

import(
	// pb_mng_service "api-gateway-service/pkg/auctions-service/pb/auctions_management_service"
	// pb_run_service "api-gateway-service/pkg/auctions-service/pb/auctions_run_service"

	"api-gateway-service/pkg/auctions-service/routes"
	"github.com/gin-gonic/gin" 
)


func RegisterRoutes(r *gin.RouterGroup) *Service{
	service := &Service{
		ManagementClient: InitManagementServiceClient(),
		RunClient: InitRunServiceClient(),
	}

	auctions := r.Group("/auctions")

	
	auctions.POST("/:id/start", service.StartAuction)
	auctions.POST("/:id/cancel", service.CancelAuction)
	auctions.POST("/:id/set-time", service.SetStartTime)

	auctions.POST(":id/:product_id/suggest", service.SuggestPrice)
	auctions.GET(":id/:product_id/price", )

	return service
}


func (s *Service) StartAuction(c *gin.Context) {
	routes.StartAuction(c, s.ManagementClient)
}


func (s *Service) CancelAuction(c *gin.Context) {
	routes.CancelAuction(c, s.ManagementClient)
}


func (s *Service) SetStartTime(c *gin.Context) {
	routes.SetStartTime(c, s.ManagementClient)
}


func (s *Service) SuggestPrice(c *gin.Context) {
	routes.SuggestPrice(c, s.RunClient)
}


func (s *Service) GetCurrentPrice(c *gin.Context) {
	routes.GetCurrentPrice(c, s.RunClient)
}