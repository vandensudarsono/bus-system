package repository

import (
	"context"

	"github.com/vandensudarsono/bus-system/domain/models"
)

type RepositoryDB interface {
	GetBusStops(ctx context.Context, offset, limit int64) []models.BusStop
	GetBuslines(ctx context.Context, offset, limit int64) []models.BusLine
	GetBuslineById(ctx context.Context, buslineID string) models.BusLine
	// GetBuslineByBusStopId: give list of bus lines based given a busStop id. So in a bus stop there are some
	// Bus lines that related to some other bus stops.
	GetBuslineByBusStopId(ctx context.Context, busStopID string) ([]models.BusLine, error)
	GetBuslineByBusStopName(ctx context.Context, name string, offset, limit int64) []models.BusLine
	InsertBusLine(ctx context.Context, input []models.BusLine) bool
	InsertBusStops(ctx context.Context, input []models.BusStop) bool
}
