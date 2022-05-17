package profile

import (
	"github.com/amchicas/go-api-gateway/pkg/auth/middleware"
	"github.com/amchicas/go-api-gateway/pkg/config"
	"github.com/amchicas/go-api-gateway/pkg/profile/handler"
	"github.com/amchicas/go-api-gateway/pkg/profile/routes"
	"github.com/amchicas/go-api-gateway/pkg/profile/service"
	fiber "github.com/gofiber/fiber/v2"
)

func Exec(app *fiber.App, c *config.Config, authMid *middleware.AuthMiddleware) {

	profileSrv := service.InitServiceClient(c)
	profileHandler := handler.NewProfileHandler(profileSrv)
	routes.RegisterRoutes(app, profileHandler, authMid)
}
