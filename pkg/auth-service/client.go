package auth_service

import(
	"google.golang.org/grpc"

	"api-gateway-service/pkg/auth-service/pb"
)

type Service struct{
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient{
	conn, err := grpc.Dial("", grpc.WithInsecure)
	if err != nil{

	}

	return pb.NewAuthServiceClient(conn)
}

