package models

type PostitionsDTOResponse struct {
	Payload []BusPosition `json:"payload"`
	Status  int           `json:"status"`
}

type BusAvailableDTOResponse struct {
	Payload []BusLine `json:"payload"`
	Status  int       `json:"status"`
}
