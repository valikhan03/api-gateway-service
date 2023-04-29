package routes

import (
	"context"

	"github.com/gin-gonic/gin"

	"api-gateway-service/pkg/command-service/pb"
)

type CreateAuctionRequest struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	OrganizerID     int64  `json:"-"`
	MaxParticipants int32  `json:"max_participants"`
}

func CreateAuction(c *gin.Context, client pb.CommandServiceClient) {
	ctx := context.TODO()

	req := CreateAuctionRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(400)
	}

	res, err := client.CreateAuction(ctx, &pb.CreateAuctionRequest{
		Auction: &pb.Auction{
			Title:           req.Title,
			Description:     req.Description,
			MaxParticipants: req.MaxParticipants,
		},
	})
	if err != nil {

	}

	c.JSON(200, res)
}

type UpdateAuctionRequest struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	OrganizerID     int64  `json:"organizer_id"`
	MaxParticipants int32  `json:"max_participants"`
}

func UpdateAuction(c *gin.Context, client pb.CommandServiceClient) {
	ctx := context.TODO()

	req := UpdateAuctionRequest{}

	err := c.BindJSON(&req)
	if err != nil{
		c.AbortWithStatus(400)
	}

	res, err := client.UpdateAuction(ctx, &pb.UpdateAuctionRequest{
		Auction: &pb.Auction{
			Id: req.Id,
			Title: req.Title,
			Description: req.Description,
			OrganizerID: 0,
			MaxParticipants: req.MaxParticipants,
		},
	})

	if err != nil{

	}

	c.JSON(200, res)
}

// needed to add validation on rights of user to delete auction
func DeleteAuction(c *gin.Context, client pb.CommandServiceClient) {
	ctx := context.TODO()

	id := c.Param("id")
	if len(id) == 0 {

	}

	res, err := client.DeleteAuction(ctx, &pb.DeleteAuctionRequest{Id: id})
	if err != nil {

	}

	c.JSON(200, res)
}


func AddParticipant(c *gin.Context, client pb.CommandServiceClient) {
	ctx := context.TODO()

	auctionId := c.Param("id")
	participantId := c.Param("participant_id")

	res, err := client.AddParticipant(ctx, &pb.AddParticipantRequest{
		AuctionId: auctionId,
		ParticipantId: participantId,
	})
	if err != nil{

	}

	c.JSON(200, res)
}

func DeleteParticipant(c *gin.Context, client pb.CommandServiceClient) {
	ctx := context.TODO()
	
	auctionId := c.Param("id")
	participantId := c.Param("participant_id")

	res, err := client.DeleteParticipant(ctx, &pb.DeleteParticipantRequest{
		AuctionId: auctionId,
		ParticipantId: participantId,
	})
	if err != nil{

	}

	c.JSON(200, res)
}
