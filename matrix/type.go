/*	This file defines the types of the matrix package
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package matrix

// definition of matrix type
type Matrix struct {
	rows    int
	cols    int
	entries [][]float64
}
