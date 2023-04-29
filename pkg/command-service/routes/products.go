package routes

import (
	"context"

	"github.com/gin-gonic/gin"

	"api-gateway-service/pkg/command-service/pb"
)

type AddProductRequest struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	StartPrice  int64             `json:"start_price"`
	Category    string            `json:"category"`
	Params      map[string]string `json:"params"`
}

func AddProduct(c *gin.Context, client pb.CommandServiceClient) {
	ctx := context.TODO()

	req := AddProductRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(400)
	}

	auctionId := c.Param("id")

	res, err := client.AddProduct(ctx, &pb.AddProductRequest{
		AuctionId: auctionId,
		Product: &pb.Product{
			Title:       req.Title,
			Description: req.Description,
			StartPrice:  req.StartPrice,
			Category:    req.Category,
			Params:      req.Params,
		},
	})
	if err != nil {

	}

	c.JSON(200, res)
}

type UpdateProductRequest struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	StartPrice  int64             `json:"start_price"`
	Category    string            `json:"category"`
	Params      map[string]string `json:"params"`
}

func UpdateProduct(c *gin.Context, client pb.CommandServiceClient) {
	ctx := context.TODO()

	req := UpdateProductRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatus(400)
	}

	auctionId := c.Param("id")

	res, err := client.UpdateProduct(ctx, &pb.UpdateProductRequest{
		AuctionId: auctionId,
		Product: &pb.Product{
			Title:       req.Title,
			Description: req.Description,
			StartPrice:  req.StartPrice,
			Category:    req.Category,
			Params:      req.Params,
		},
	})
	if err != nil {

	}

	c.JSON(200, res)
}

func DeleteProduct(c *gin.Context, client pb.CommandServiceClient) {
	ctx := context.TODO()

	auctionId := c.Param("id")
	productId := c.Param("product_id")

	res, err := client.DeleteProduct(ctx, &pb.DeletePoductRequest{
		AuctionId: auctionId,
		ProductId: productId,
	})
	if err != nil {

	}

	c.JSON(200, res)
}
