package usecase

import (
	"context"

	"github.com/vandensudarsono/bus-system/domain/repository"
)

type BusInteractor struct {
	repo repository.RepositoryDB
	out  OutputPort
}

func NewBusInteractor(repo repository.RepositoryDB, out OutputPort) *BusInteractor {
	return &BusInteractor{
		repo: repo,
		out:  out,
	}
}

func (bi *BusInteractor) GetBusStops(ctx context.Context, offset, limit int64) (interface{}, int) {
	result := bi.repo.GetBusStops(ctx, offset, limit)
	return bi.out.Present(ctx, result, nil, nil)
}

func (bi *BusInteractor) GetBuslinesAvailable(ctx context.Context, offset, limit int64) (interface{}, int) {

	result := bi.repo.GetBuslines(ctx, offset, limit)

	return bi.out.Present(ctx, result, nil, nil)
}
