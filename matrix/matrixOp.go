/*	This file defines matrix methods
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

// CopyMatrix returns a new matrix with the same contents
func (a *Matrix) CopyMat() (m *Matrix) {
	m, _ = ZeroMat(a.rows, a.cols)

	for i := range m.rowVectors {
		m.rowVectors[i] = a.rowVectors[i].CopyVec()
	}
	return
}

// Add computes componentwise sum of matrices a and b.
// Computes c = a + b, if dimensions match.
// Dimension mismatch returns nil.
func (a *Matrix) Add(b *Matrix) (c *Matrix) {
	// check if dimensions match
	if a.rows != b.rows || a.cols != b.cols {
		return
	}
	// create new matrix
	c, _ = ZeroMat(a.rows, a.cols)
	// iterate over matrices and add up row vectors
	for i := range a.rowVectors {
		vecA := a.rowVectors[i]
		vecB := b.rowVectors[i]
		c.rowVectors[i] = vecA.Add(vecB)
	}
	return
}

// Sub computes componentwise differences of matrices a and b.
// Computes c = a - b, if dimensions match.
// Dimension mismatch returns nil
func (a *Matrix) Sub(b *Matrix) (c *Matrix) {
	// negate b
	negB := b.ApplyFunc(func(x float64) float64 { return (-1) * x })
	// compute diff
	c = a.Add(negB)
	return
}

// Scale returns a scaled matrix by factor f.
// Returns nil pointer if factor = inf, or NaN
func (a *Matrix) Scale(factor float64) (c *Matrix) {
	// check if factor is valid
	if math.IsNaN(factor) {
		return
	}
	if math.IsInf(factor, 1) {
		return
	}
	if math.IsInf(factor, -1) {
		return
	}
	// multiply matrix a by factor
	c = a.ApplyFunc(func(x float64) float64 { return x * factor })
	return
}

// CWiseProd computes the compnent-wise product of matrices a and b.
// Dimension mismatch returns nil
func (a *Matrix) CWiseProd(b *Matrix) (c *Matrix) {
	// check if sizes match
	if a.rows != b.rows || a.cols != b.cols {
		return
	}
	// allocate new matrix
	c, _ = ZeroMat(a.rows, a.cols)
	// compute cwise product
	for i := range a.rowVectors {
		vecA := a.rowVectors[i]
		vecB := b.rowVectors[i]
		c.rowVectors[i] = vecA.CWiseProd(vecB)
	}
	return
}

// ApplyFunc returns a new matrix which contains the entries of a after applying func f.
func (a *Matrix) ApplyFunc(f func(float64) float64) (m *Matrix) {
	// create new matrix
	m, _ = ZeroMat(a.rows, a.cols)

	for i := range a.rowVectors {
		vec := a.rowVectors[i]
		m.rowVectors[i] = vec.ApplyFunc(f)
	}
	return
}
