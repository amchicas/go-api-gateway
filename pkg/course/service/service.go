package service

import (
	"fmt"

	"github.com/amchicas/go-api-gateway/pkg/config"
	"github.com/amchicas/go-course-srv/pkg/pb"
	"google.golang.org/grpc"
)

func InitServiceClient(c *config.Config) pb.CourseServiceClient {

	conn, err := grpc.Dial(c.CourseUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewCourseServiceClient(conn)
}
