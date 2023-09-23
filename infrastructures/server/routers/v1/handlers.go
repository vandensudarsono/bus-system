package handlers

import "github.com/gofiber/fiber/v2"

type BuslineHandlers interface {
	GetAvailableLines(ctx *fiber.Ctx) error
	GetBuslineDetailById(ctx *fiber.Ctx) error
	GetBuslineByBusStopName(ctx *fiber.Ctx) error
	GetBuslinesDetailByBusStopId(ctx *fiber.Ctx) error
}
