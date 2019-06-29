/*	This file defines getter and setters of the matrix package
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package matrix

import (
	"fmt"
	"math"
)

// Rows returns the number of rows of the matrix
func (m *Matrix) Rows() int {
	return m.rows
}

// Cols returns the number of cols of the matrix
func (m *Matrix) Cols() int {
	return m.cols
}

// Get(i,j) returns the element in row i and column j.
// returns NaN if invalid index
func (m *Matrix) Get(i, j int) float64 {
	v, e := m.GetSafe(i, j)
	if e != nil {
		return math.NaN()
	}
	return v
}

// Set(i,j, v) sets the entry of row i and col j to value v.
// No update, if invalid indices
func (m *Matrix) Set(i, j int, v float64) {
	m.SetSafe(i, j, v)
	return
}

// GetSafe(i,j) returns the element in row i and column j and an error value
func (m *Matrix) GetSafe(i, j int) (elem float64, e error) {
	// check if valid indices
	if i < 0 || j < 0 {
		e = fmt.Errorf("Error: indices are negative")
		return
	}
	// check if out of bounds
	if i >= m.rows || j >= m.cols {
		e = fmt.Errorf("Error: indices are out of bound")
		return
	}
	// return elem
	elem = m.rowVectors[i].entries[j]
	return
}

// SetSafe makes the same as safe, but it doesn't ignore the error values
func (m *Matrix) SetSafe(i, j int, v float64) (e error) {
	// check if valid indices
	if i < 0 || j < 0 {
		e = fmt.Errorf("Error: indices are negative")
		return
	}
	// check if out of bounds
	if i >= m.rows || j >= m.cols {
		e = fmt.Errorf("Error: indices are out of bound")
		return
	}
	// set entry
	m.rowVectors[i].entries[j] = v
	return
}

// GetRow returns a new vector of the i-th row of matrix m.
// Returns nil if invalid index
func (m *Matrix) GetRow(i int) (v *Vector) {
	// check if valid idx
	if i < 0 || i >= m.rows {
		return
	}
	// return vector
	v = m.rowVectors[i].CopyVec()
	return
}

// SetRow sets i-th row of matrix m.
// Doesn't update, if invalid index or missmatching sizes,
func (m *Matrix) SetRow(i int, vec *Vector) {
	// check if valid i and vec
	if i < 0 || i >= m.rows || vec.Size() != m.cols {
		return
	}
	// update row
	m.rowVectors[i] = vec.CopyVec()
}

// GetCol returns a new vector with the contents of j-th column.
// Returns nil if invalid index
func (m *Matrix) GetCol(j int) (v *Vector) {
	// check if valid idx
	if j < 0 || j >= m.cols {
		return
	}
	// return vector
	v = ZeroVec(m.rows)
	for i := range m.rowVectors {
		v.entries[i] = m.rowVectors[i].Get(j)
	}
	return
}

// SetRow sets j-th column of matrix m.
// Doesn't update, if invalid index or missmatching sizes,
func (m *Matrix) SetCol(j int, vec *Vector) {
	// check sizes
	if j < 0 || j >= m.cols || vec.Size() != m.rows {
		return
	}
	// update
	for i := range m.rowVectors {
		m.rowVectors[i].Set(j, vec.Get(i))
	}
}

// Get returns the i-th element in vector v.
// Returns NaN if invalid index
func (v *Vector) Get(i int) float64 {
	// check if valid index
	if i < 0 || i >= v.Size() {
		return math.NaN()
	}
	// return element
	return v.entries[i]
}

// Set sets the i-th entry to value.
// If i is an ivalid index, then nothing is updated
func (v *Vector) Set(i int, value float64) {
	// check index
	if i < 0 || i >= v.Size() {
		return
	}
	// update i-th entry
	v.entries[i] = value
	return
}
