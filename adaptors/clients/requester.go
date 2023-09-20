package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	dto "github.com/vandensudarsono/bus-system/domain/DTO"
)

type Clients struct {
	client *fiber.Client
}

func NewClients(c *fiber.Client) *Clients {
	return &Clients{client: c}
}

func (cl *Clients) GetRunningBusInABusLines(ctx context.Context, busLineID string) ([]dto.BusPosition, error) {
	var (
		result []dto.BusPosition
		uWave  dto.PostitionsDTOResponse
		err    error
	)

	url := fmt.Sprintf("%s/%s", viper.GetStringSlice("requester.urls")[1], busLineID)

	agent := fiber.AcquireAgent()

	request := agent.Request()
	request.SetRequestURI(url)
	request.Header.SetMethod(fiber.MethodGet)

	if err = agent.Parse(); err != nil {
		return result, err
	}

	response := fiber.AcquireResponse()

	err = agent.DoDeadline(request, response, time.Now().Add(time.Minute))
	if err != nil {
		return result, err
	}

	res := response.Body()
	err = json.Unmarshal(res, &uWave)
	if err != nil {
		return result, err
	}

	result = append(result, uWave.Payload...)

	return result, err

}
