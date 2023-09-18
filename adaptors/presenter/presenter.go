package presenter

import (
	"context"

	"github.com/vandensudarsono/bus-system/domain/models"
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
			return models.Response{
				Status: models.Status{Code: errorcode.ErrNotFound, Message: errorcode.ErrNotFound.String()},
			}, 200
		default:
			return models.Response{
				Status: models.Status{Code: errorcode.ErrUnknown, Message: errorcode.ErrUnknown.String()},
			}, 200
		}
	}

	return models.Response{
		Data: data,
		Meta: meta,
		Status: models.Status{
			Code:    0,
			Message: "success",
		},
	}, 200
}
