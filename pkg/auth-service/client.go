package auth_service

import(
	"google.golang.org/grpc"
	"log"
	"api-gateway-service/pkg/models"

	"api-gateway-service/pkg/auth-service/pb"
)

type Service struct{
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient{
	conn, err := grpc.Dial(models.ConfigGlobal.AuthService.URL, grpc.WithInsecure)
	if err != nil{
		log.Fatalf("Auth service connection error: %v", err)
	}

	return pb.NewAuthServiceClient(conn)
}

