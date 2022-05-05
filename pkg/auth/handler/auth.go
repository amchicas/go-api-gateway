package handler

import (
	"context"
	"fmt"

	"github.com/amchicas/go-api-gateway/pkg/auth/domain"
	"github.com/amchicas/go-auth-srv/pkg/pb"
	fiber "github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService pb.AuthServiceClient
}

func NewAuthHandler(authService pb.AuthServiceClient) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Index(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{"status": "ok", "messages": "index test"})
}
func (h *AuthHandler) Register(c *fiber.Ctx) error {

	var input domain.Auth
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := c.BodyParser(&input); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "fail", "message": err.Error()})

	}

	req := &pb.RegisterReq{Username: input.Username, Email: input.Email, Password: input.Password, Role: 1}
	res, err := h.authService.Register(customContext, req)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on login request", "data": err.Error()})
	}
	return c.JSON(&fiber.Map{"status": fmt.Sprint(res.Status)})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	/*type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	*/
	var input domain.Auth
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Error on login request", "data": err.Error()})
	}

	req := &pb.LoginReq{Email: input.Email, Password: input.Password}
	res, err := h.authService.Login(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "message": err.Error()})
	}
	return c.JSON(&fiber.Map{"status": "success", "result": fmt.Sprint(res.Token)})
}
