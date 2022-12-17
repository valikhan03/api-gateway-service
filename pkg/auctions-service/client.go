package auctions_service

import(
	"google.golang.org/grpc"

	pb_mng_service "api-gateway-service/pkg/auctions-service/pb/auctions_management_service"
	pb_run_service "api-gateway-service/pkg/auctions-service/pb/auctions_run_service"
)


type Service struct{
	ManagementClient pb_mng_service.AuctionsManagementServiceClient
	RunClient pb_run_service.AuctionsRunServiceClient
}


func InitManagementServiceClient() pb_mng_service.AuctionsManagementServiceClient {
	conn, err := grpc.Dial("", grpc.WithInsecure())
	if err != nil{

	}

	return pb_mng_service.NewAuctionsManagementServiceClient(conn)
}


func InitRunServiceClient() pb_run_service.AuctionsRunServiceClient {
	conn, err := grpc.Dial("", grpc.WithInsecure())
	if err != nil{

	}

	return pb_run_service.NewAuctionsRunServiceClient(conn)
}