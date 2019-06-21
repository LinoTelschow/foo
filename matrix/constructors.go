/*	This file describes methods to construct matrices
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package matrix

import (
	"fmt"
)

// ZeroMat creates a zero matrix with r rows and c column
func ZeroMat(r,c int) (m *Matrix, e error) {
	if r <= 0 || c <= 0 {
		e = fmt.Errorf("Invalid dimensions")
		return
	}
	// create Matrix
	m = new(Matrix)
	m.rows = r
	m.cols = c

	// allocate array for entries
	m.entries = make([][]float64, r)
	for i := range m.entries {
		m.entries[i] = make([]float64, c)
	}
	return
}

// IdMat creates a identity matrix with r rows and c column
// The diagonal entries are set to 1
func IdMat(r,c int) (m *Matrix, e error) {
	// create empty matrix
	m, e = ZeroMat(r,c)
	if e != nil {
		return
	}

	// set diagonal to 1
	for i := range m.entries {
		if i < m.cols {
			m.entries[i][i] = 1.0
		}
	}
	return
}

// CreateFromSlice creates a matrix from a 2d slice of type: [][]float64
// Method assumes that all slices of []float64 have the same length
func CreateFromSlice(slice [][]float64) (m *Matrix, e error) {
	// check argument
	if len(slice) < 1 {
		e = fmt.Errorf("Error: empty slice")
		return
	}

	// reference value for col size
	col := len(slice[0])

	// check if empty col
	if col == 0 {
		e = fmt.Errorf("Error: empty column in slice")
		return
	}

	// inspect number of cols
	for i := range slice {
		if col != len(slice[i]) {
			e = fmt.Errorf("Error: mismatching column lengths in slice")
			return
		}
	}

	// Create matrix
	m = new(Matrix)
	m.rows = len(slice)
	m.cols = col
	m.entries = make([][]float64, len(slice))
	for i := range m.entries {
		// create rows
		m.entries[i] = make([]float64, col)
		// copy values from input slice
		copy(m.entries[i], slice[i])
	}
	return
}