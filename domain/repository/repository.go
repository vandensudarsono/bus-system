package repository

import (
	"context"

	"github.com/vandensudarsono/bus-system/domain/models"
)

type RepositoryDB interface {
	GetBusStops(ctx context.Context, offset, limit int64) []models.BusStop
	GetBuslines(ctx context.Context, offset, limit int64) []models.BusLine
	GetBuslineById(ctx context.Context, buslineID string) models.BusLine
	InsertBusLine(ctx context.Context, input []models.BusLine) bool
	InsertBusStops(ctx context.Context, input []models.BusStop) bool
}
