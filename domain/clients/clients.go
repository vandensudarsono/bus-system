package clients

import (
	"context"

	dto "github.com/vandensudarsono/bus-system/domain/DTO"
	"github.com/vandensudarsono/bus-system/domain/models"
)

type Clients interface {
	// GetRunningBusInABusLines: call uWave API to get running bus in a busline
	GetRunningBusInABusLines(ctx context.Context, busLineID string) ([]dto.BusPosition, error)
	// GetAvailabeBuslines: call uWave API to get available buslines
	GetAvailableBuslines(ctx context.Context) ([]models.BusLine, error)
}
