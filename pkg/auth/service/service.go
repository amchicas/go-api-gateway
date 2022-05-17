package service

import (
	"fmt"

	"github.com/amchicas/go-api-gateway/pkg/config"
	"github.com/amchicas/go-auth-srv/pkg/pb"
	"google.golang.org/grpc"
)

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	conn, err := grpc.Dial(c.AuthUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewAuthServiceClient(conn)
}
