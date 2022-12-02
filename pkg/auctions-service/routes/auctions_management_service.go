package routes

import (
	"context"

	"github.com/gin-gonic/gin"

	pb "api-gateway-service/pkg/auctions-service/pb/auctions_management_service"
)

func StartAuction(c *gin.Context, client pb.AuctionsManagementServiceClient) {
	id := c.Param(":id")

	res, err := client.StartAuction(context.TODO(), &pb.StartAuctionRequest{AuctionId: id})
	if err != nil {
		c.AbortWithStatus(500)
	}

	if res.Status >= 200 && res.Status < 300 {
		c.JSON(int(res.Status), gin.H{"error": res.Error})
	}

	c.Status(int(res.Status))
}

func CancelAuction(c *gin.Context, client pb.AuctionsManagementServiceClient) {
	id := c.Param(":id")

	res, err := client.CancelAuction(context.TODO(), &pb.CancelAuctionRequest{AuctionId: id})
	if err != nil {
		c.AbortWithStatus(500)
	}

	if res.Status >= 200 && res.Status < 300 {
		c.JSON(int(res.Status), gin.H{"error": res.Error})
	}

	c.Status(int(res.Status))
}

type SetStartTimeReqBody struct {
	AuctionId string `json:"id"`
	StartTime string `json:"start_time"`
	TZ        string `json:"tz"`
}

func SetStartTime(c *gin.Context, client pb.AuctionsManagementServiceClient) {
	var reqbody SetStartTimeReqBody
	err := c.BindJSON(&reqbody)
	if err != nil{
		c.AbortWithStatus(400)
		return
	}


	res, err := client.SetStartTime(context.TODO(), &pb.SetStartTimeRequest{})
	if err != nil{
		c.AbortWithStatus(500)
	}

	if res.Status >= 200 && res.Status < 300 {
		c.JSON(int(res.Status), gin.H{"error": res.Error})
	}

	c.Status(int(res.Status))
}
