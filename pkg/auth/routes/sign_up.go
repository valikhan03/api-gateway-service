package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gateway-service/pkg/auth/pb"
)

type SignUpRequestBody struct {
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	PhoneNum  string `json:"phone_num"`
	BirthDate string `json:"birth_date"`
}

func SignUp(c *gin.Context, client pb.AuthServiceClient) {
	var req SignUpRequestBody
	err := c.BindJSON(&req)
	if err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.SignUp(context.TODO(), &pb.SignUpRequest{
		FullName: req.Fullname,
		Email: req.Email,
		Password: req.Password,
		PhoneNum: req.PhoneNum,
		BirthDate: req.BirthDate,
	})
	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(int(res.Status), res)
}