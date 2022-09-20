package routes

import (
	"api-gateway-service/pkg/search/pb"
	"context"

	"github.com/gin-gonic/gin"
)

type SearchProductsRequestBody struct {
	SearchReq string `json:"search_request"`
	Page      int32  `json:"page"`
}

func SearchProducts(c *gin.Context, client pb.SearchServiceClient) {
	var req SearchProductsRequestBody
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	res, err := client.SearchProduct(context.TODO(), &pb.SearchProductsRequest{SearchRequest: req.SearchReq, Page: req.Page})
	if err != nil{
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, gin.H{
		"products":string(res.Products),
	})
}