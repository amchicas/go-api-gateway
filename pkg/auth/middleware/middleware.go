package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/amchicas/go-auth-srv/pkg/pb"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authService pb.AuthServiceClient
}

func InitAuthMiddleware(authService pb.AuthServiceClient) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}

}
func (a *AuthMiddleware) Validate(ctx *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	auth := ctx.Get(fiber.HeaderAuthorization)
	if auth == "" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Error Token",
		})
	}
	token := strings.Split(auth, "Bearer ")
	if len(token) != 2 {

		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Error Token",
		})

	}
	req := &pb.ValidateReq{Token: token[1]}
	res, err := a.authService.Validate(customContext, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  fmt.Sprint(res.Status),
			"message": fmt.Sprint(res.Error),
		})
	}
	ctx.Locals("userID", res.UserId)
	ctx.Locals("username", res.Username)
	ctx.Locals("role", res.Role)
	return ctx.Next()
}
