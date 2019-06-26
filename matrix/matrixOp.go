/*	This file defines matrix operations
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package matrix

import (
	"fmt"
	"math"
)

// Add computes componentwise sum of matrices a and b.
// Computes c = a + b, if dimensions match
func (a *Matrix) Add(b *Matrix) (c *Matrix, e error) {
	// check if dimensions match
	if a.rows != b.rows || a.cols != b.cols {
		e = fmt.Errorf("Error: Dimension missmatch")
		return
	}

	// create new matrix
	c, e = ZeroMat(a.rows, a.cols)
	if e != nil {
		return
	}

	// iterate over matrices and add up
	for i := range a.entries {
		for j := range a.entries[i] {
			c.entries[i][j] = a.entries[i][j] + b.entries[i][j]
		}
	}
	return
}

// Sub computes componentwise differences of matrices a and b.
// Computes c = a - b, if dimensions match
func (a *Matrix) Sub(b *Matrix) (c *Matrix, e error) {
	// negate b
	negB := b.ApplyFunc(func(x float64) float64 { return (-1) * x })
	// compute diff
	c, e = a.Add(negB)
	return
}

// Scale returns a scaled matrix by factor f.
// Returns nil pointer if factor = inf, or NaN
func (a *Matrix) Scale(factor float64) (c *Matrix, e error) {
	// check if factor is valid
	if math.IsNaN(factor) {
		e = fmt.Errorf("Error: factor is NaN")
		return
	}
	if math.IsInf(factor, 1) {
		e = fmt.Errorf("Error: factor is Inf")
		return
	}
	if math.IsInf(factor, -1) {
		e = fmt.Errorf("Error: factor is -Inf")
		return
	}

	// multiply matrix a by factor
	c = a.ApplyFunc(func(x float64) float64 { return x * factor })
	return
}

// CWiseProd computes the compnent-wise product of matrices a and b.
// Returns nil if sizes don't match
func (a *Matrix) CWiseProd(b *Matrix) (c *Matrix, e error) {
	// check if sizes match
	if a.rows != b.rows || a.cols != b.cols {
		e = fmt.Errorf("Error: mismatching dimensions")
		return
	}
	// allocate new matrix
	c, e = ZeroMat(a.rows, a.cols)
	if e != nil {
		return
	}
	// compute cwise product
	for i := range a.entries {
		for j := range a.entries[i] {
			c.entries[i][j] = a.entries[i][j] * b.entries[i][j]
		}
	}
	return
}

// ApplyFunc returns a new matrix which contains the entries of a after applying func f.
func (a *Matrix) ApplyFunc(f func(float64) float64) (m *Matrix) {
	// create new matrix
	m, _ = ZeroMat(a.rows, a.cols)

	for i := range a.entries {
		for j := range a.entries[i] {
			m.entries[i][j] = f(a.entries[i][j])
		}
	}
	return
}
