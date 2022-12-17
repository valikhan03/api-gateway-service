package routes

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	pb "api-gateway-service/pkg/auctions-service/pb/auctions_run_service"
)

type SuggestPriceReqBody struct {
	ProductID string `json:"product_id"`
	NewPrice  int64  `json:"new_price"`
}

func SuggestPrice(c *gin.Context, client pb.AuctionsRunServiceClient) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatus(401)
		return
	}


	var reqbody SuggestPriceReqBody
	err := c.BindJSON(&reqbody)
	if err != nil{
		c.AbortWithStatus(400)
		return
	}

	req := &pb.SuggestPriceRequest{UserID: id.(int64), ProductID: reqbody.ProductID, NewPrice: reqbody.NewPrice}
	res, err := client.SuggestPrice(context.TODO(), req)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	if res.Status >= 200 && res.Status < 300 {
		c.JSON(int(res.Status), gin.H{"error": res.Error})
	}

	c.Status(int(res.Status))
}

var wsupgrader = websocket.Upgrader {
	HandshakeTimeout: 3 * time.Second,
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func GetCurrentPrice(c *gin.Context, client pb.AuctionsRunServiceClient) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	for{
		stream, err := client.GetCurrentPrice(context.TODO())
		if err != nil{
			c.AbortWithError(500, err)
			return
		}

		res, err := stream.Recv()
		
		err = conn.WriteJSON(map[string]string{"user":res.Username, "price":res.Price})
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
	}
}
