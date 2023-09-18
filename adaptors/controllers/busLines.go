package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vandensudarsono/bus-system/usecase"
)

type BusLinesController struct {
	uc usecase.InputPort
}

func NewBusLinesController(uc usecase.InputPort) *BusLinesController {
	return &BusLinesController{
		uc: uc,
	}
}

func (c *BusLinesController) GetAvailableLines(ctx *fiber.Ctx) error {
	//sanitize request

	result, code := c.uc.GetBuslinesAvailable(ctx.Context(), 0, 10)
	ctx.JSON(result)

	return ctx.SendStatus(code)
}
