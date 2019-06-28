/*	This file defines matrix operations
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package matrix

import (
	"math"
)

// CopyMatrix returns a new matrix with the same content
func (a *Matrix) CopyMat() (m *Matrix) {
	m, _ = ZeroMat(a.rows, a.cols)

	for i := range m.rowVectors {
		m.rowVectors[i] = a.rowVectors[i].CopyVec()
	}
	return
}

// Add computes componentwise sum of matrices a and b.
// Computes c = a + b, if dimensions match.
// Dimension mismatch returns nil
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
