package auth

import (
	"github.com/amchicas/go-api-gateway/pkg/auth/handler"
	"github.com/amchicas/go-api-gateway/pkg/auth/middleware"
	"github.com/amchicas/go-api-gateway/pkg/auth/routes"
	"github.com/amchicas/go-api-gateway/pkg/auth/service"
	"github.com/amchicas/go-api-gateway/pkg/config"
	fiber "github.com/gofiber/fiber/v2"
)

func Exec(app *fiber.App, c *config.Config) *middleware.AuthMiddleware {

	authSrv := service.InitServiceClient(c)
	authHandler := handler.NewAuthHandler(authSrv)
	routes.RegisterRoutes(app, authHandler)

	return middleware.InitAuthMiddleware(authSrv)
}
