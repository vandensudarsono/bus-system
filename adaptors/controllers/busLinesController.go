package controllers

import (
	"github.com/gofiber/fiber/v2"
	dto "github.com/vandensudarsono/bus-system/domain/DTO"
	errorcode "github.com/vandensudarsono/bus-system/internal/errorCode"
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

func (c *BusLinesController) GetBuslineDetailById(ctx *fiber.Ctx) error {
	//sanitize request
	var (
		result interface{}
		code   int
	)

	queries := ctx.Query("busLineID")
	if queries == "" {
		result = dto.Response{
			Status: dto.Status{
				Code:    errorcode.ErrParams,
				Message: errorcode.ErrParams.String(),
			},
		}

		code = fiber.StatusOK

	} else {
		result, code = c.uc.GetBuslineDetailById(ctx.Context(), queries)
	}

	ctx.JSON(result)

	return ctx.SendStatus(code)
}

func (c *BusLinesController) GetBuslineByBusStopName(ctx *fiber.Ctx) error {
	var (
		result interface{}
		code   int
	)

	queries := ctx.Query("name")
	if queries == "" {
		result = dto.Response{
			Status: dto.Status{
				Code:    errorcode.ErrParams,
				Message: errorcode.ErrParams.String(),
			},
		}

		code = fiber.StatusOK

	} else {
		result, code = c.uc.GetBuslinesByBusStopName(ctx.Context(), queries)
	}

	ctx.JSON(result)

	return ctx.SendStatus(code)
}

func (c *BusLinesController) GetBuslinesDetailByBusStopId(ctx *fiber.Ctx) error {
	var (
		result interface{}
		code   int
	)

	queries := ctx.Query("busStopID")
	if queries == "" {
		result = dto.Response{
			Status: dto.Status{
				Code:    errorcode.ErrParams,
				Message: errorcode.ErrParams.String(),
			},
		}

		code = fiber.StatusOK

	} else {
		result, code = c.uc.GetBuslinesByBusStopId(ctx.Context(), queries)
	}

	ctx.JSON(result)

	return ctx.SendStatus(code)
}
