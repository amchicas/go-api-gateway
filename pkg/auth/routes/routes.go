package routes

import (
	"github.com/amchicas/go-api-gateway/pkg/auth/handler"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterRoutes(app *fiber.App, handler handler.AuthHandler) {

	auth := app.Group("/auth", logger.New())
	auth.Post("/login", handler.Login)
	user := app.Group("/user")
	user.Post("/", handler.Register)

}
