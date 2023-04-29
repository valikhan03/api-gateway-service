gen-auth-proto:
	protoc -I . --go_out=. --go-grpc_out=. pkg/auth-service/protobuf/auth_service.proto

gen-search-proto:
	protoc -I . --go_out=. --go-grpc_out=. pkg/search-service/protobuf/search_service.proto

gen-command-proto:
	protoc -I . --go_out=. --go-grpc_out=. pkg/command-service/protobuf/command_service.proto

gen-auctions-proto:
	protoc -I . --go_out=. --go-grpc_out=. pkg/auctions-service/protobuf/auctions_management_service.proto
	protoc -I . --go_out=. --go-grpc_out=. pkg/auctions-service/protobuf/auctions_run_service.proto