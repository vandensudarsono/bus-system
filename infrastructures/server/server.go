package server

import "github.com/gofiber/fiber/v2"

func NewServer() *fiber.App {
	config := fiber.Config{
		ServerHeader: "bus-lines",
		AppName:      "bus-lines-api",
	}

	App := fiber.New(config)

	return App
}
