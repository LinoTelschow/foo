/* This package implements Euclidean's Algorithm */

package euclid

// Euclids algo
// computes gcd(a,b)
func Euclid(a, b int) int {
	if a == 0 {
		return b
	} else {
		for b != 0 {
			if a > b {
				a = a - b
			} else {
				b = b - a
			}
		}
		return a
	}
}

// Extended Euclid algo
// Computes gcd(a,b) = d and s, t such that gcd(a,b) = s*a + t*b
func ExtendedEuclid(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	d1, s1, t1 := ExtendedEuclid(b, a%b)
	d, s, t := d1, t1, s1-(a/b)*t1
	return d, s, t
}
