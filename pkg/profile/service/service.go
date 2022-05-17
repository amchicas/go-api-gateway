package service

import (
	"fmt"

	"github.com/amchicas/go-api-gateway/pkg/config"
	"github.com/amchicas/go-profile-srv/pkg/pb"
	"google.golang.org/grpc"
)

func InitServiceClient(c *config.Config) pb.ProfileServiceClient {

	conn, err := grpc.Dial(c.ProfileUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewProfileServiceClient(conn)
}
