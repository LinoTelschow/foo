/*	This file describes methods to construct matrices and vectors.
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package matrix

import (
	"fmt"
)

// ZeroMat creates a zero matrix with r rows and c column
func ZeroMat(r, c int) (m *Matrix, e error) {
	if r <= 0 || c <= 0 {
		e = fmt.Errorf("Invalid dimensions")
		return
	}
	// create Matrix
	m = new(Matrix)
	m.rows = r
	m.cols = c
	// allocate array for entries
	m.rowVectors = make([]*Vector, r)
	for i := range m.rowVectors {
		m.rowVectors[i] = ZeroVec(c)
	}
	return
}

// IdMat creates a identity matrix with r rows and c column
// The diagonal entries are set to 1
func IdMat(r, c int) (m *Matrix, e error) {
	// create empty matrix
	m, e = ZeroMat(r, c)
	if e != nil {
		return
	}
	// set diagonal to 1
	for i := range m.rowVectors {
		if i < m.cols {
			m.rowVectors[i].Set(i, 1.0)
		}
	}
	return
}

// CreateFromSlice creates a matrix from a 2d slice of type: [][]float64
// Method assumes that all slices of []float64 have the same length
func MatrixFromSlice(slice [][]float64) (m *Matrix, e error) {
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
	m, _ = ZeroMat(len(slice), col)
	for i := range m.rowVectors {
		// copy values from input slice
		copy(m.rowVectors[i].entries, slice[i])
	}
	return
}

// CopyMatrix copies the whole content of src to a new matrix
func CopyMat(src *Matrix) (m *Matrix) {
	m, _ = ZeroMat(src.rows, src.cols)

	for i := range m.rowVectors {
		m.rowVectors[i] = CopyVector(src.rowVectors[i])
	}
	return
}

// ZeroVec creates a zero vector with r rows and c column.
// Return nil if invalid size parameter
func ZeroVec(size int) (v *Vector) {
	// check size
	if size < 1 {
		return
	}
	// create vector
	v = new(Vector)
	v.entries = make([]float64, size)
	return v
}

// VecFromSlice creates a vector from a slice.
// Returns nil, if zero sized slice
func VecFromSlice(slice []float64) (v *Vector) {
	// check slice
	if len(slice) < 1 {
		return
	}
	// create vector
	v = ZeroVec(len(slice))
	copy(v.entries, slice)
	return v
}

// CopyVector copies the whole content of src to a new vector
func CopyVector(src *Vector) (v *Vector) {
	v = VecFromSlice(src.entries)
	return
}
