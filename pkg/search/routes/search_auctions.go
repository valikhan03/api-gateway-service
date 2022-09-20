package routes

import (
	"api-gateway-service/pkg/search/pb"
	"context"
	"net/http"

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
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.SearchAuctions(context.TODO(), &pb.SearchAuctionsRequest{SearchRequest: req.SearchReq, Page: req.Page})
	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	

	c.JSON(200, gin.H{
		"auctions": string(res.Auctions),	
	})
}
