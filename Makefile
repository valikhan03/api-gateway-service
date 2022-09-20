gen-auth-proto:
	protoc -I . --go_out=. --go-grpc_out=. pkg/auth/protobuf/auth_service.proto

gen-search-proto:
	protoc -I . --go_out=. --go-grpc_out=. pkg/search/protobuf/search_service.proto