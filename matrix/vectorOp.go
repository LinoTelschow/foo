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

// CopyVector returns new vector with the same content
func (a *Vector) CopyVec() (v *Vector) {
	v = VecFromSlice(a.entries)
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
	// add vectors
	c = ZeroVec(a.Size())
	for i := range a.entries {
		c.entries[i] = a.entries[i] - b.entries[i]
	}
	return
}

// Scale returns a scaled vector by factor f.
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
