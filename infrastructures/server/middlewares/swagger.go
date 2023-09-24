package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupSwagger(server *fiber.App) {
	server.Use("/swagger/*", swagger.New(swagger.Config{
		//URL:         "http://localhost:8081/docs/swagger.yaml",
		//URL:         "http://localhost:8081/swagger/doc.yaml",
		URL:         "/doc.json",
		Title:       "Busline API",
		DeepLinking: false,
	}))
}
