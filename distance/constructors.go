/*	This file defines the constructors of the distance package
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package distance

import (
	"fmt"
)

// This function returns a new coordinate variable
func NewCoordinate(city string, country string, isS bool, latDeg int, latMin int, isW bool, longDeg int, longMin int) (c *Coordinate, e error) {
	// check if invalid inputs
	if latDeg < 0 || latDeg >= 90 || latMin < 0 || latMin >= 60 ||
		longDeg < 0 || longDeg >= 180 || longMin < 0 || longMin >= 60 {
		e = fmt.Errorf("Invalid Coordinates")
		return
	}
	// init fields
	c = new(Coordinate)
	c.City = city
	c.Country = country
	c.IsSouth = isS
	c.Latitude = float64(latDeg) + float64(latMin)/60
	c.IsWest = isW
	c.Longitude = float64(longDeg) + float64(longMin)/60
	return
}

// This function returns a new normalized coordinate variable
func NewNormCoordinate(c *Coordinate) (nc *NormCoordinate, e error) {
	// check if invalid coordinate
	if c == nil {
		e = fmt.Errorf("invalid coordinate")
	}
	nc = new(NormCoordinate)
	// check latitude
	if c.IsSouth {
		nc.Latitude = -1.0 * c.Latitude
	} else {
		nc.Latitude = c.Latitude
	}
	// check longitude
	if !c.IsWest {
		nc.Longitude = 360 - c.Longitude
	} else {
		nc.Longitude = c.Longitude
	}
	return
}

// This function constructs a new CoordinateSet
func NewCoordinateSet(coord []*Coordinate) (cs *CoordinateSet, e error) {
	// check if not empty
	if len(coord) == 0 {
		e = fmt.Errorf("Error, empty input slice")
		return
	}
	cs = new(CoordinateSet)
	cs.Points = make([]*Coordinate, len(coord))
	copy(cs.Points, coord)
	return
}
