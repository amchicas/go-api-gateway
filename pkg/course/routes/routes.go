package routes

import (
	"github.com/amchicas/go-api-gateway/pkg/auth/middleware"
	"github.com/amchicas/go-api-gateway/pkg/course/handler"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterRoutes(app *fiber.App, handler handler.CourseHandler, authMid *middleware.AuthMiddleware) {

	course := app.Group("/course", logger.New())

	course.Post("/", authMid.Validate, handler.Add)
	course.Patch("/:id", authMid.Validate, handler.Update)
	course.Delete("/:id", authMid.Validate, handler.Delete)
	course.Get("/:id", handler.Find)

}
