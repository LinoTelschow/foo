/*	This file defines methods for important go interfaces
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package matrix

import (
	"fmt"
)

// implements the Stringer interface for matrix type
func (m Matrix) String() string {
	var s string
	// print dimension
	s = "---------------------------------------------------------------------------------------------------------------------\n"
	s = s + fmt.Sprintf("Dimension: Rows: %d \t Cols: %d \n", m.rows, m.cols)
	s = s + "Matrix: \n"

	// print matrix
	for i := range m.entries {
		localString := ""
		for j := range m.entries[i] {
			localString = localString + fmt.Sprintf("%10.4g ", m.entries[i][j])
		}
		s = s + localString + "\n"
	}
	s = s + "---------------------------------------------------------------------------------------------------------------------\n"
	return s
}

// implements the Stringer interface for vector type
func (v Vector) String() string {
	var s string
	// print size
	s = "---------------------------------------------------------------------------------------------------------------------\n"
	s = s + fmt.Sprintf("Size: %d\n", len(v.entries))
	s = s + "Vector: "
	s = s + fmt.Sprintf("%v\n", v.entries)
	s = s + "---------------------------------------------------------------------------------------------------------------------\n"
	return s
}
