package routes

import (
	"api-gateway-service/pkg/search/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)



func GetProduct(c *gin.Context, client pb.SearchServiceClient) {
	id := c.Param("id")
	if len(id) == 0{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := client.GetProduct(context.TODO(), &pb.GetProductInfoRequest{ProductId: id})
	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, res)
}	