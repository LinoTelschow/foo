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
	elem = m.getEntry(i,j)
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
	m.setEntry(i,j, v)
	return
}

// GetRow returns a new vector of the i-th row of matrix m.
// Returns nil if invalid index
func (m *Matrix) GetRow(i int) (v *Vector) {
	// check if valid idx
	if i < 0 || i >= m.rows {
		return
	}
	// create vector
	v = ZeroVec(m.cols)
	// copy values
	for j := 0; j < m.cols; j++ {
		v.entries[j] = m.getEntry(i,j)
	}
	return
}

// SetRow sets i-th row of matrix m.
// Doesn't update, if invalid index or missmatching sizes,
func (m *Matrix) SetRow(i int, vec *Vector) {
	// check if valid i and vec
	if i < 0 || i >= m.rows || vec.Size() != m.cols {
		return
	}
	// update i-th row
	for j := 0; j < m.cols; j++ {
		m.setEntry(i,j, vec.entries[j])
	}
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
	for i := 0; i < m.rows; i++ {
		v.entries[i] = m.getEntry(i,j)
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
	// update j-th column
	for i := 0; i < m.rows; i++ {
		m.setEntry(i,j, vec.entries[i])
	}
}

// CopyMatrix returns a new matrix with the same contents
func (a *Matrix) CopyMat() (m *Matrix) {
	m, _ = ZeroMat(a.rows, a.cols)

	// copy entries
	for i := range a.entries {
		m.entries[i] = a.entries[i]
	}
	return
}

// GetBlock returns a submatrix of a.
// urow = upper row, ucol = upper col
// lrow = lower row, lcol = lower col
// example:              |. . . . . . .|
//  x = (urow, ucol)     |. x-------- .|
//  y = (lrow, lcol)     |. | . . . | .|
//                       |. --------y .|
//                       |. . . . . . .|
// if invalid indices returns nil.
// The indices are always inclusive.
// (like in GetSubVec and SetSubVec)
func (a *Matrix) GetBlock(urow, ucol, lrow, lcol int) (b *Matrix) {
	// check range
	if urow < 0 || urow >= a.Rows() || lrow < 0 || lrow >= a.Rows() ||
		ucol < 0 || ucol >= a.Cols() || lcol < 0 || lcol >= a.Cols() {
		return
	}
	// check
	if urow > lrow || ucol > lcol {
		return
	}
	// define dimensions
	r := lrow - urow + 1
	c := lcol - ucol + 1
	// create matrix
	b, _ = ZeroMat(r, c)
	// copy entries
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.cols; j++ {
			b.setEntry(i,j, a.getEntry(urow+i, ucol + j))
		}
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
	// iterate over matrices and add up
	for i := range a.entries {
		c.entries[i] = a.entries[i] + b.entries[i]
	}
	return
}

// Sub computes componentwise differences of matrices a and b.
// Computes c = a - b, if dimensions match.
// Dimension mismatch returns nil
func (a *Matrix) Sub(b *Matrix) (c *Matrix) {
	// check if dimensions match
	if a.rows != b.rows || a.cols != b.cols {
		return
	}
	// create new matrix
	c, _ = ZeroMat(a.rows, a.cols)
	// iterate over matrices and substract
	for i := range a.entries {
		c.entries[i] = a.entries[i] - b.entries[i]
	}
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
	for i := range a.entries {
		c.entries[i] = a.entries[i] * b.entries[i]
	}
	return
}

// ApplyFunc returns a new matrix which contains the entries of a after applying func f.
func (a *Matrix) ApplyFunc(f func(float64) float64) (m *Matrix) {
	// create new matrix
	m, _ = ZeroMat(a.rows, a.cols)
	// apply function on elements
	for i := range a.entries {
		m.entries[i] = f(a.entries[i])
	}
	return
}

// getEntry returns the entry in the i row and j column.
// not exported, intended to use as building block for library
func (a *Matrix) getEntry(i,j int) float64 {
	return a.entries[a.cols * i + j]
}

// setEntry sets value v to the matrix element at row i and column j.
// not exported, intended to use as building block for library
func (a *Matrix) setEntry(i,j int, v float64) {
	a.entries[a.cols * i + j] = v
}
