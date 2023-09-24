package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupBuslineHandler(app *fiber.App, h BuslineHandlers) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	//buslines
	buslines := v1.Group("/buslines")
	buslines.Get("/get-buslines", h.GetAvailableLines)
	buslines.Get("/details", h.GetBuslineDetailById)
	buslines.Get("/by-busstop-name", h.GetBuslineByBusStopName)
	buslines.Get("/by-busstop-id", h.GetBuslinesDetailByBusStopId)
}
