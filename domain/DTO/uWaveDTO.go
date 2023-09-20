package dto

import "github.com/vandensudarsono/bus-system/domain/models"

type PostitionsDTOResponse struct {
	Payload []BusPosition `json:"payload"`
	Status  int           `json:"status"`
}

type BusAvailableDTOResponse struct {
	Payload []models.BusLine `json:"payload"`
	Status  int              `json:"status"`
}
