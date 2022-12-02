package routes

import (
	"api-gateway-service/pkg/search-service/pb"
	"context"

	"github.com/gin-gonic/gin"
)

type SearchAuctionsRequestBody struct {
	SearchReq string `json:"search_request"`
	Page      int32  `json:"page"`
}

func SearchAuctions(c *gin.Context, client pb.SearchServiceClient) {
	var req SearchAuctionsRequestBody
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	res, err := client.SearchAuctions(context.TODO(), &pb.SearchAuctionsRequest{SearchRequest: req.SearchReq, Page: req.Page})
	if err != nil{
		c.AbortWithError(500, err)
		return
	}
	

	c.JSON(int(res.Status), res.Auctions)
}
