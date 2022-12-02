package command_service

import(
	"google.golang.org/grpc"
	
	"api-gateway-service/pkg/command-service/pb"
)


type Service struct {
	Client pb.CommandServiceClient
}

func InitServiceClient() pb.CommandServiceClient {
	conn, err := grpc.Dial("", grpc.WithInsecure())
	if err != nil{

	}

	return pb.NewCommandServiceClient(conn)
}