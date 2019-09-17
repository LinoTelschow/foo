/*	This file defines the types of the matrix package
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package matrix

// definiton of vector type
type Vector struct {
	entries []float64
}

// definition of matrix type
type Matrix struct {
	rows    int
	cols    int
	entries []float64
}
