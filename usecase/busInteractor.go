package usecase

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/vandensudarsono/bus-system/adaptors/clients"
	dto "github.com/vandensudarsono/bus-system/domain/DTO"
	"github.com/vandensudarsono/bus-system/domain/models"
	"github.com/vandensudarsono/bus-system/domain/repository"
	errorcode "github.com/vandensudarsono/bus-system/internal/errorCode"
	"github.com/vandensudarsono/bus-system/internal/geospacial"
	"github.com/vandensudarsono/bus-system/internal/helpers"
)

type BusInteractor struct {
	repo   repository.RepositoryDB
	client clients.Clients
	out    OutputPort
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

func (bi *BusInteractor) GetBuslineDetailById(ctx context.Context, busLineID string) (interface{}, int) {

	result := bi.repo.GetBuslineById(ctx, busLineID)

	if result.Id == "" {
		// Busline Datail is not found
		return bi.out.Present(ctx, nil, nil, fmt.Errorf(errorcode.ErrNotFound.String()))
	}

	return bi.out.Present(ctx, result, nil, nil)
}

func (bi *BusInteractor) GetBuslinesByBusStopId(ctx context.Context, busStopID string) (interface{}, int) {
	var response dto.BuslinesByBusStopID

	type (
		// busLineWithVehiclesMutex struct {
		// 	sync.Mutex
		// 	blwv map[string][]dto.BuslineWithVehichles
		// }

		availableVehiclesMutex struct {
			sync.Mutex
			av map[string][]dto.BusPositionWithTime
		}
	)

	result, err := bi.repo.GetBuslineByBusStopId(ctx, busStopID)
	if err != nil || len(result) == 0 {
		fmt.Printf("ERROR: at get buslines by id: %v. or there is no buslines: %d", err, len(result))

		return bi.out.Present(ctx, nil, nil, fmt.Errorf(errorcode.ErrNotFound.String()))
	}

	//get bus id data in a busline
	busStop := helpers.GetLatitudeAndLongitudeOfABusStopID(busStopID, result[0].BusStops)
	response.BusStop = busStop

	var (
		wg, wg2 sync.WaitGroup
		//blwv    busLineWithVehiclesMutex
		av availableVehiclesMutex
	)

	//request all available buses in those buslines result
	for i := 0; i < len(result); i++ {
		wg.Add(1)

		go func(busLine models.BusLine) {
			defer wg.Done()

			busPositions, err := bi.client.GetRunningBusInABusLines(ctx, busLine.Id)
			if err != nil || len(busPositions) == 0 {
				fmt.Println("ERROR: there an error for getting running bus in bus lines: ", err, len(busPositions))
			} else {

				for j := 0; j < len(busPositions); j++ {

					wg2.Add(1)

					go func(bp dto.BusPosition) {
						defer wg2.Done()

						busTime := geospacial.CountTimeRemaining(
							[]float64{float64(bp.Lat), float64(bp.Lng)},
							[]float64{busStop.Lat, busStop.Lng},
							busLine.Path,
						)

						av.Lock()

						av.av[busLine.Id] = append(av.av[busLine.Id], dto.BusPositionWithTime{
							Bearing:       bp.Bearing,
							CrowdLevel:    bp.CrowdLevel,
							Lat:           bp.Lat,
							Lng:           bp.Lng,
							VehiclePlate:  bp.VehiclePlate,
							TimeRemaining: strconv.Itoa(int(busTime*60)) + "m", // forced convert to minute
						})

						av.Unlock()

					}(busPositions[j])
				}

				wg2.Wait()
				//wait get time

			}

		}(result[i])

	}

	//wait all result
	wg.Wait()
	//get bus available and its time

	//initate its  result element into response model
	for _, data := range result {
		response.BuslineWithVehichles = append(
			response.BuslineWithVehichles,
			dto.BuslineWithVehichles{
				BusStops:         data.BusStops,
				FullName:         data.FullName,
				Id:               data.Id,
				Origin:           data.Origin,
				Path:             data.Path,
				ShortName:        data.ShortName,
				AvailableVehicle: av.av[data.Id],
			})
	}

	return bi.out.Present(ctx, response, nil, err)
}

func (bi *BusInteractor) GetBuslinesByBusStopName(ctx context.Context, busStopName string) (interface{}, int) {
	//var wg, wg2 sync.WaitGroup

	result := bi.repo.GetBuslineByBusStopName(ctx, busStopName, 0, 10)
	lenResult := len(result)

	if lenResult == 0 {
		return bi.out.Present(ctx, nil, nil, fmt.Errorf(errorcode.ErrNotFound.String()))
	}

	//request all available buses in those buslines result
	// for i := 0; i < lenResult; i++ {
	// 	wg.Add(1)

	// 	go func(busLine models.BusLine) {
	// 		defer wg.Done()

	// 		busPositions, err := bi.client.GetRunningBusInABusLines(ctx, busLine.Id)
	// 		if err != nil || len(busPositions) == 0 {

	// 		} else {
	// 			for j := 0; j < len(busPositions); j++ {
	// 				wg2.Add(1)
	// 				go func(bp models.BusPosition) {
	// 					geospacial.CountTimeRemaining(
	// 						[]float64{float64(bp.Lat), float64(bp.Lng)},
	// 						[]float64{}
	// 					)
	// 				}(busPositions[j])
	// 			}

	// 		}

	// 	}(result[i])

	// }

	// //wait all result
	// wg.Done()

	return bi.out.Present(ctx, result, nil, nil)

}
