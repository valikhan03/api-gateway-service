package auth_service

import(
	"github.com/gin-gonic/gin"

	"api-gateway-service/pkg/auth-service/routes"

)

func RegisterRoutes(r *gin.Engine) *Service{
	service := &Service{
		Client: InitServiceClient(),
	}

	auth := r.Group("auth")
	auth.POST("/sign-up", service.SignUp)
	auth.POST("/sign-in", service.SignIn)

	return service
}

func (s *Service) SignUp(c *gin.Context) {
	routes.SignUp(c, s.Client)
}

func (s *Service) SignIn(c *gin.Context) {
	routes.SignIn(c, s.Client)
}