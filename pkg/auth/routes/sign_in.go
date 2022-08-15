package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gateway-service/pkg/auth/pb"
)

type SignInRequestBody struct{
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func SignIn(c *gin.Context, client pb.AuthServiceClient) {
	var req SignInRequestBody
	err := c.BindJSON(&req)
	if err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.SignIn(context.TODO(), &pb.SignInRequest{})
	if err != nil{
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if ok, err := validateToken(res.Token); !ok {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.SetCookie("token_auth", res.Token, 1, "", "", true, true)

	
}

func validateToken(token string) (bool, error) {
	return true, nil
}