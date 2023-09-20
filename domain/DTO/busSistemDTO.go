package dto

import (
	"github.com/vandensudarsono/bus-system/domain/models"
	errorcode "github.com/vandensudarsono/bus-system/internal/errorCode"
)

type (
	Response struct {
		Data   interface{} `json:"data"`
		Meta   interface{} `json:"meta"`
		Status Status      `json:"status"`
	}

	Status struct {
		Code    errorcode.ErrorCode `json:"code"`
		Message string              `json:"message"`
	}

	BusPositionWithTime struct {
		Bearing       float32 `json:"bearing"`
		CrowdLevel    string  `json:"crowdLevel"`
		Lat           float32 `json:"lat"`
		Lng           float32 `json:"lng"`
		VehiclePlate  string  `json:"vehiclePlate"`
		TimeRemaining string  `json:"timeRemaining"`
	}

	BuslineWithVehichles struct {
		BusStops         []models.BusStop      `json:"busStops" bson:"busStops"`
		FullName         string                `json:"fullName" bson:"fullName"`
		Id               string                `json:"id" bson:"id"`
		Origin           string                `json:"origin" bson:"origin"`
		Path             [][]float64           `json:"path" bson:"path"`
		ShortName        string                `json:"shortName" bson:"shortName"`
		AvailableVehicle []BusPositionWithTime `json:"availableVehicles"`
	}

	BuslinesByBusStopID struct {
		BusStop              models.BusStop         `json:"busStop"`
		BuslineWithVehichles []BuslineWithVehichles `json:"buslineWithVehicles"`
	}
)
