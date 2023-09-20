package presenter

import (
	"context"

	dto "github.com/vandensudarsono/bus-system/domain/DTO"
	errorcode "github.com/vandensudarsono/bus-system/internal/errorCode"
)

type Presenter struct{}

func NewPresenter() *Presenter {
	return &Presenter{}
}

func (p *Presenter) Present(ctx context.Context, data, meta interface{}, err error) (interface{}, int) {

	if err != nil {
		switch err.Error() {
		case errorcode.ErrNotFound.String():
			return dto.Response{
				Status: dto.Status{Code: errorcode.ErrNotFound, Message: errorcode.ErrNotFound.String()},
			}, 200
		default:
			return dto.Response{
				Status: dto.Status{Code: errorcode.ErrUnknown, Message: errorcode.ErrUnknown.String()},
			}, 200
		}
	}

	return dto.Response{
		Data: data,
		Meta: meta,
		Status: dto.Status{
			Code:    0,
			Message: "success",
		},
	}, 200
}
