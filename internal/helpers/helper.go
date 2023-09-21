package helpers

import "github.com/vandensudarsono/bus-system/domain/models"

func GetLatitudeAndLongitudeOfABusStopID(busStopID string, busStops []models.BusStop) models.BusStop {
	var busStop models.BusStop

	for i := 0; i < len(busStops); i++ {
		if busStops[i].Id == busStopID {
			busStop = busStops[i]
			break
		}
	}

	return busStop
}
