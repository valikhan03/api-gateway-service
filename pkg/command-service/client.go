package command_service

import(
	"google.golang.org/grpc"
	"api-gateway-service/pkg/models"
	"log"

	"api-gateway-service/pkg/command-service/pb"
)


type Service struct {
	Client pb.CommandServiceClient
}

func InitServiceClient() pb.CommandServiceClient {
	conn, err := grpc.Dial(models.ConfigGlobal.CommandService.URL, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Command service connection error: %v", err)
	}

	return pb.NewCommandServiceClient(conn)
}