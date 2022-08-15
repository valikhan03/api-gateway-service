package routes

import (
	"context"
	"net/http"
	"time"

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

	//add domain later
	cookie := http.Cookie{
		Name: "auth",
		Value: res.Token,
		Path: "/api/v1",
		Expires: time.Now().Add(24 * time.Hour),
		MaxAge: 60*60*24,
	}

	http.SetCookie(c.Writer, &cookie)
}
