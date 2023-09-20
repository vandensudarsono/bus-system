package models

type (
	BusStop struct {
		Id   string  `json:"id" bson:"id"`
		Lat  float64 `json:"lat" bson:"lat"`
		Lng  float64 `json:"lng" bson:"lng"`
		Name string  `json:"name" bson:"name"`
	}

	BusLine struct {
		BusStops  []BusStop   `json:"busStops" bson:"busStops"`
		FullName  string      `json:"fullName" bson:"fullName"`
		Id        string      `json:"id" bson:"id"`
		Origin    string      `json:"origin" bson:"origin"`
		Path      [][]float64 `json:"path" bson:"path"`
		ShortName string      `json:"shortName" bson:"shortName"`
	}
)
