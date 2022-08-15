package auth

import (
	"api-gateway-service/pkg/auth/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct{
	service *Service
}

func InitAuthMiddleware(srvc *Service) *AuthMiddlewareConfig{
	return &AuthMiddlewareConfig{srvc}
}

func (a *AuthMiddlewareConfig) AuthRequired(c *gin.Context){
	token, err := c.Cookie("auth")
	if err != nil{
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if token == ""{
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := a.service.Client.Validate(context.TODO(), &pb.ValidateRequest{Token: token})
	if err != nil || res.Status != http.StatusOK{
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user_id", res.UserID)

	c.Next()
}