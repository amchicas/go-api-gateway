package main

import (
	"log"

	"github.com/amchicas/go-api-gateway/pkg/auth/handler"
	"github.com/amchicas/go-api-gateway/pkg/auth/routes"
	auth "github.com/amchicas/go-api-gateway/pkg/auth/services"
	"github.com/amchicas/go-api-gateway/pkg/config"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	app := fiber.New()
	app.Use(logger.New())

	authSrv := auth.InitServiceClient(&c)
	authHandler := handler.NewAuthHandler(authSrv)
	routes.RegisterRoutes(app, authHandler)

	log.Fatal(app.Listen(c.Port))
}
