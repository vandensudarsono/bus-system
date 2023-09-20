package geospacial

import (
	"math"
	"sync"
)

func haversineDistance(point []float64, point2 []float64) float64 {
	// Earth radius in kilometers (you can adjust this value)
	const earthRadiusKm = 6371.0

	// Convert latitude and longitude from degrees to radians
	lat1 := point[0] * math.Pi / 180.0
	lon1 := point[1] * math.Pi / 180.0
	lat2 := point2[0] * math.Pi / 180.0
	lon2 := point2[1] * math.Pi / 180.0

	// Haversine formula
	dlat := lat2 - lat1
	dlon := lon2 - lon1
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadiusKm * c

	return distance
}

func findClosestPointIndex(currentPoint []float64, path [][]float64) (int, float64) {

	if len(path) == 0 {
		return -1, 0.0 // No points in the path
	}

	closestIndex := 0
	closestDistance := haversineDistance(currentPoint, path[0])

	for i, pathPoint := range path {
		distance := haversineDistance(currentPoint, pathPoint)
		if distance < closestDistance {
			closestDistance = distance
			closestIndex = i
		}
	}

	return closestIndex, closestDistance
}

//CalucateTotalDistance: sum from current distance till the bus stop positions

func CountTimeRemaining(CurrentPositionBus, CurrentPositionBusStop []float64, Point [][]float64) float64 {

	var (
		wg                                                                      sync.WaitGroup
		closestIndex, clossestIndexBusStop                                      int
		closestDistance, closestDistanceBusStop, distanceBusAndBusStop, busTime float64
	)

	//add paralel execution to cut 2 loop time
	wg.Add(2)

	go func([]float64, [][]float64) {
		defer wg.Done()

		// Find the index of the closest point in the path to the current point
		closestIndex, closestDistance = findClosestPointIndex(CurrentPositionBus, Point)

	}(CurrentPositionBus, Point)

	go func([]float64, [][]float64) {
		defer wg.Done()

		clossestIndexBusStop, closestDistanceBusStop = findClosestPointIndex(CurrentPositionBusStop, Point)

	}(CurrentPositionBusStop, Point)

	// TODO: calculate speed
	// this actually need a background processs to get current position of bus
	// respect to the time i.e. 5s or 10s or 1 minute. But it will cost in time
	// response so i will assumed this vehicle constant at 20 km/hour of speed.
	// it need more additional work in the future to separate this process in
	// another container.

	speed := 20.0

	//wait all result
	wg.Wait()

	// To get estimated time incoming bus reach the current bus stop, we will add
	// all of the distance from last position to the current bus stop position and
	// devide it with actual speed.

	//calculate total distance from a point to required bus stop
	distanceBusAndBusStop = closestDistance + closestDistanceBusStop

	for closestIndex := closestIndex; closestIndex < clossestIndexBusStop; closestIndex++ {
		distanceBusAndBusStop += haversineDistance(Point[closestIndex], Point[closestIndex+1])
	}

	// time is distance / speed
	// fmt.Println("distanceBusAndBusStop: ", distanceBusAndBusStop)

	busTime = distanceBusAndBusStop / speed

	// fmt.Println("busTime", math.Round(busTime*60))

	return busTime
}
