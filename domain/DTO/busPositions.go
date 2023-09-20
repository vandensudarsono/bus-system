package dto

type (
	//uWaveDTO
	BusPosition struct {
		Bearing      float32 `json:"bearing"`
		CrowdLevel   string  `json:"crowdLevel"`
		Lat          float32 `json:"lat"`
		Lng          float32 `json:"lng"`
		VehiclePlate string  `json:"vehiclePlate"`
	}

	BusPositions map[string][]BusPosition
)
