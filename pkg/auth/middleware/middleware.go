package middleware

import (
	"fmt"

	"github.com/amchicas/go-auth-srv/pkg/pb"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authService pb.AuthServiceClient
}

func InitAuthMiddleware(authService pb.AuthServiceClient) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}

}
func (a *AuthMiddleware) AuthRequired(c *fiber.Ctx) {
	auth := c.Get(fiber.HeaderAuthorization)
	fmt.Println(auth)
}
