package routes

import (
	"api-gateway-service/pkg/search-service/pb"
	"context"

	"github.com/gin-gonic/gin"
)

type SearchRequestBody struct {
	SearchReq string `json:"search_request"`
	Page      int32  `json:"page"`
}

func Search(c *gin.Context, client pb.SearchServiceClient) {
	var req SearchRequestBody
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	products_chan := make(chan *pb.SearchProductsResponse, 1)
	auctions_chan := make(chan *pb.SearchAuctionsResponse, 1)
	err_chan := make(chan error)

	go func() {
		res_products, err := client.SearchProducts(context.TODO(), &pb.SearchProductsRequest{SearchRequest: req.SearchReq, Page: req.Page})
		if err != nil {
			err_chan <- err
		}
		products_chan <- res_products
	}()

	go func() {
		res_auctions, err := client.SearchAuctions(context.TODO(), &pb.SearchAuctionsRequest{SearchRequest: req.SearchReq, Page: req.Page})
		if err != nil {
			err_chan <- err
		}
		auctions_chan <- res_auctions
	}()

	res := form_resp(products_chan, auctions_chan, err_chan)

	c.JSON(int(res.Status), res)
}

type response struct {
	Match    int64
	Status   int32
	Error    string
	Auctions []byte
	Products []byte
}

func form_resp(
	products_chan chan *pb.SearchProductsResponse,
	auctions_chan chan *pb.SearchAuctionsResponse,
	err_chan chan error) *response {

	r := response{
		Match:  0,
		Error:  "",
		Status: 200,
	}

	cnt := 0

	// make it return value
	for {
		select {
		case products := <-products_chan:
			r.Match += products.Match
			r.Error += products.Error + "\n"
			r.Products = products.Products
			cnt++
			if cnt == 2 {
				break
			}
		case auctions := <-auctions_chan:
			r.Match += auctions.Match
			r.Error += auctions.Error + "\n"
			r.Auctions = auctions.Auctions
			cnt++
			if cnt == 2 {
				break
			}
		case err := <-err_chan:
			r.Error = err.Error()
			r.Status = 500
			break
		default:
		}
	}

	return &r
}
