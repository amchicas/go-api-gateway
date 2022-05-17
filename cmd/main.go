package main

import (
	"log"

	"github.com/amchicas/go-api-gateway/pkg/auth"
	"github.com/amchicas/go-api-gateway/pkg/config"
	"github.com/amchicas/go-api-gateway/pkg/profile"
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
	authMid := auth.Exec(app, &c)
	profile.Exec(app, &c, authMid)
	log.Fatal(app.Listen(c.Port))
}
