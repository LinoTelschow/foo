/*	This file defines methods for important go interfaces
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package distance

import (
	"fmt"
)

// implements the Stringer interface for matrix type
func (c Coordinate) String() string {
	var s string
	s = fmt.Sprintf("City: %s, Country: %s, ", c.City, c.Country)
	// case distinction on South/North
	if c.IsSouth {
		s = s + fmt.Sprintf("Latitude: %g° S ", c.Latitude)
	} else {
		s = s + fmt.Sprintf("Latitude: %g° N ", c.Latitude)
	}
	// case distinction on West/East
	if c.IsWest {
		s = s + fmt.Sprintf("Longitude: %g° W", c.Longitude)
	} else {
		s = s + fmt.Sprintf("Longitude: %g° E", c.Longitude)
	}
	return s
}
