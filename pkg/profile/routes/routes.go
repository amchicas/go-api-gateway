package routes

import (
	"github.com/amchicas/go-api-gateway/pkg/auth/middleware"
	"github.com/amchicas/go-api-gateway/pkg/profile/handler"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterRoutes(app *fiber.App, handler handler.ProfileHandler, authMid *middleware.AuthMiddleware) {

	profile := app.Group("/profile", logger.New())

	profile.Post("/:id", authMid.Validate, handler.Add)
	profile.Patch("/:id", authMid.Validate, handler.Update)
	profile.Delete("/:id", authMid.Validate, handler.Delete)
	profile.Get("/:id", handler.Find)

}
