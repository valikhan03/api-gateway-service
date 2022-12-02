package search

import (
	"api-gateway-service/pkg/search-service/pb"

	"google.golang.org/grpc"
)

type Service struct{
	Client pb.SearchServiceClient
}

func InitServiceClient() pb.SearchServiceClient {
	conn, err := grpc.Dial("", grpc.WithInsecure())
	if err != nil{

	}

	return pb.NewSearchServiceClient(conn)
}