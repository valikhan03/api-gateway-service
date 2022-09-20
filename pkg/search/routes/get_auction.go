package routes

import (
	"api-gateway-service/pkg/search/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAuction(c *gin.Context, client pb.SearchServiceClient) {
	id := c.Param("id")
	if len(id) == 0{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := client.GetAuction(context.TODO(), &pb.GetAuctionInfoRequest{AuctionId: id})
	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, res)
}