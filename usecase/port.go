package usecase

import "context"

type InputPort interface {
	// GetBusStops: get list of all bus stops available
	GetBusStops(ctx context.Context, offset, limit int64) (interface{}, int)
	// GetBusLinesInBusStop: by given bus stop id, user can get all available busliness
	//GetBuslinesInBusStop(ctx context.Context, busStopID string) (interface{}, int)
	// GetBuslinesAvailable: get all buslines available at this time
	GetBuslinesAvailable(ctx context.Context, offset, limit int64) (interface{}, int)
	GetBuslineDetailById(ctx context.Context, busLineID string) (interface{}, int)
	GetBuslinesByBusStopName(ctx context.Context, busStopName string) (interface{}, int)
	GetBuslinesByBusStopId(ctx context.Context, busStopID string) (interface{}, int)
}

type OutputPort interface {
	Present(ctx context.Context, data, meta interface{}, err error) (interface{}, int)
}
