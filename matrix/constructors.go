/*	This file implements functions to construct matrices and vectors.
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
	m.entries = make([]float64, r*c)
	return
}

// IdMat creates a identity matrix with r rows and c columns
// The diagonal entries are set to 1
func IdMat(r, c int) (m *Matrix, e error) {
	// create empty matrix
	m, e = ZeroMat(r, c)
	if e != nil {
		return
	}
	// set diagonal to 1
	offset := 0
	nrOfEntries := m.rows * m.cols
	for i := 0; i < m.rows; i++ {
		if (i*c + offset) < nrOfEntries {
			m.entries[i*c+offset] = 1.0
			offset++
		} else {
			break
		}
	}
	return
}

// CreateFromSlice creates a matrix from a 2d slice of type: [][]float64
// Method assumes that all slices of []float64 have the same length
func MatrixFromSlice(slice [][]float64) (m *Matrix, e error) {
	// check size of slice
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
	// copy entries
	c := m.cols
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.entries[i*c+j] = slice[i][j]
		}
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
