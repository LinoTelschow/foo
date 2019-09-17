/*	This file defines the types of the distance package
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package distance

// This type represents a real world coordinate
type Coordinate struct {
	City      string
	Country   string
	IsSouth   bool
	Latitude  float64
	IsWest    bool
	Longitude float64
}

// This type represents a normalized coordinate
// (used to compute distances)
type NormCoordinate struct {
	Latitude  float64
	Longitude float64
}

// This type represents a Set of Coordinates
type CoordinateSet struct {
	Points []*Coordinate
}
