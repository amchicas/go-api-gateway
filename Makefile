proto:
	protoc --proto_path=./pkg/auth/pb/proto ./pkg/auth/pb/proto/*.proto  --go-grpc_out=. --go_out=.
server:
	go run cmd/main.go
