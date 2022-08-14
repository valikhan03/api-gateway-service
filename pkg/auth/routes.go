package auth

import(
	"github.com/gin-gonic/gin"

	"api-gateway-service/pkg/auth/routes"

)

func RegisterRoutes(r *gin.Engine) *Service{
	service := &Service{
		Client: InitServiceClient(),
	}

	routes := r.Group("auth")
	routes.POST("/sign-up", service.SignUp)
	routes.POST("/sign-in", service.SignIn)

	return service
}

func (s *Service) SignUp(c *gin.Context) {
	routes.SignUp(c, s.Client)
}

func (s *Service) SignIn(c *gin.Context) {
	routes.SignIn(c, s.Client)
}