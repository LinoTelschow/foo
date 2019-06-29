/*	This file defines vector operations
	Author: Lino Telschow, tlino@student.ethz.ch
*/

package matrix

import (
	"math"
)

// returns the size of vector a.
func (a *Vector) Size() int {
	return len(a.entries)
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

// CopyVector returns a new vector with the same content
func (a *Vector) CopyVec() (v *Vector) {
	v = VecFromSlice(a.entries)
	return
}

// Mat returns n x 1 matrix with the values of a.
// n = Size of vector.
func (a *Vector) Mat() (m *Matrix) {
	m, _ = ZeroMat(a.Size(), 1)
	m.SetCol(0, a)
	return
}

// Slice returns a slice of type []float64 with contents of a.
func (a *Vector) Slice() (s []float64) {
	s = make([]float64, a.Size())
	for i := range s {
		s[i] = a.entries[i]
	}
	return
}

// Merge returns a merged vector c.
// c is the vector b appended to a
func (a *Vector) Merge(b *Vector) (c *Vector) {
	length := a.Size() + b.Size()
	offset := a.Size()
	c = ZeroVec(length)
	// copy values from a
	for i := range a.entries {
		c.entries[i] = a.entries[i]
	}
	// copy values from b
	for i := range b.entries {
		c.entries[offset+i] = b.entries[i]
	}
	return
}

// GetSubVec returns the subvector from index s (inclusive)
// to index e (inclusive). (returns entries: a[s], ... , a[e]).
// Returns nil if invalid indices.
func (a *Vector) GetSubVec(s, e int) (v *Vector) {
	// check indices
	if s < 0 || s >= a.Size() || e < 0 || e >= a.Size() || s > e {
		return
	}
	// copy values in new vector
	length := e - s + 1
	v = ZeroVec(length)
	for i := range v.entries {
		v.entries[i] = a.entries[s+i]
	}
	return
}

// SetSubVec sets subvector from s (inclusive) to e (inclusive).
// in a to the values of b. (modifies entries: a[s], ... , a[e])
// No update if invalid indices or missmatching sizes.
func (a *Vector) SetSubVec(s int, e int, b *Vector) {
	// check indices
	if s < 0 || s >= a.Size() || e < 0 || e >= a.Size() || s > e {
		return
	}
	length := e - s + 1
	// check sizes
	if a.Size() < b.Size() || b.Size() != length {
		return
	}
	// set values in a to b
	for i := range b.entries {
		a.entries[s+i] = b.entries[i]
	}
	return
}

// Add returns c = a + b.
// Return nil, if sizes not match.
func (a *Vector) Add(b *Vector) (c *Vector) {
	// check sizes
	if a.Size() != b.Size() {
		return
	}
	// add vectors
	c = ZeroVec(a.Size())
	for i := range a.entries {
		c.entries[i] = a.entries[i] + b.entries[i]
	}
	return
}

// Sub returns c = a - b.
// Return nil, if sizes not match.
func (a *Vector) Sub(b *Vector) (c *Vector) {
	// check sizes
	if a.Size() != b.Size() {
		return
	}
	// substract vectors
	c = ZeroVec(a.Size())
	for i := range a.entries {
		c.entries[i] = a.entries[i] - b.entries[i]
	}
	return
}

// Scale returns a scaled vector by factor.
// Returns nil, if factor = inf, or NaN
func (a *Vector) Scale(factor float64) (v *Vector) {
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
	// create scaled matrix
	v = a.ApplyFunc(func(x float64) float64 { return x * factor })
	return
}

// CWiseProd computes the compnent-wise product of vectors a and b.
// Returns nil if sizes don't match.
func (a *Vector) CWiseProd(b *Vector) (c *Vector) {
	// check sizes
	if a.Size() != b.Size() {
		return
	}
	// create c
	c = ZeroVec(a.Size())
	for i := range a.entries {
		c.entries[i] = a.entries[i] * b.entries[i]
	}
	return
}

// Dot returns the dot product of a and b
// return NaN if sizes don't match
func (a *Vector) Dot(b *Vector) float64 {
	// check sizes
	if a.Size() != b.Size() {
		return math.NaN()
	}
	var result float64 = 0
	for i := range a.entries {
		result += a.entries[i] * b.entries[i]
	}
	return result
}

// ApplyFunc returns a new vector which contains the entries of a after applying func f.
func (a *Vector) ApplyFunc(f func(float64) float64) (v *Vector) {
	// create new vector
	v = ZeroVec(a.Size())
	for i := range a.entries {
		v.entries[i] = f(a.entries[i])
	}
	return
}
