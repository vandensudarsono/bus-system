package models

type BusPosition struct {
	Bearing      int     `json:"bearing"`
	CrowdLevel   string  `json:"crowdLevel"`
	Lat          float32 `json:"lat"`
	Lng          float32 `json:"lng"`
	VehiclePlate string  `json:"vehiclePlate"`
}
