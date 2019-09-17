package distance

import (
	"fmt"
	"math"
)

// constants
var EarthRadius int = 6371000

// Computes distance between coordinates s and t by using Haversine formula.
// Needs to use normalized coordinates to compute distances
func (s *Coordinate) Distance(t *Coordinate) float64 {
	normSrc, _ := NewNormCoordinate(s)
	normDst, _ := NewNormCoordinate(t)
	return normSrc.Distance(normDst)
}

// compute distance between normalized coordinates s and t by using Haversine formula
func (s *NormCoordinate) Distance(t *NormCoordinate) float64 {
	deltaLat := toRad(s.Latitude) - toRad(t.Latitude)
	deltaLong := toRad(s.Longitude) - toRad(t.Longitude)
	a := math.Sin(deltaLat/2) * math.Sin(deltaLat/2)
	b := math.Cos(toRad(s.Latitude)) * math.Cos(toRad(t.Latitude))
	c := math.Sin(deltaLong/2) * math.Sin(deltaLong/2)
	return 2 * float64(EarthRadius) * math.Asin(math.Sqrt(a+b*c))
}

func (cs *CoordinateSet) AllDistances() {
	size := len(cs.Points)
	counter := 1
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			fmt.Println("--------------------------------------------------------------------------------------------------")
			src := cs.Points[i]
			dst := cs.Points[j]
			fmt.Printf("Nr: %d \n", counter)
			fmt.Printf("Src: %v \n", src)
			fmt.Printf("Dst: %v \n", dst)
			fmt.Printf("Distance: %g \n", src.Distance(dst))
			counter++
		}
	}
}

// helper converts degree to rad
func toRad(deg float64) float64 {
	if deg >= 0 && deg < 360 {
		return (deg / 180) * math.Pi
	} else if deg > 360 {
		factor := math.Floor(deg / 360)
		newDeg := deg - float64(factor*360)
		return (newDeg / 180) * math.Pi
	} else if deg < 0 && deg >= 360 {
		newDeg := deg + 360
		return (newDeg / 180) * math.Pi
	} else {
		factor := -1.0 * math.Floor(deg/360)
		newDeg := deg + factor*360
		return (newDeg / 180) * math.Pi
	}
}
