package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func SetupMiddleware(server *fiber.App) {
	//use logger for each request
	server.Use(logger.New(logger.Config{
		Format:     "${pi} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))

	//use fiber pprof
	server.Use(pprof.New(pprof.Config{Prefix: "/api"}))
}
