// Package nyamath includes some functions for mathmatics.
package nyamath

// IsNumber checks if a rune is in ['0'..'9']
func IsNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

//Qpow calculates x^y and returns the result with the quick pow algorithm.
func Qpow(x, y int) int {
	if y < 0 {
		return 0
	}
	ret := 1
	for y > 0 {
		if y&1 != 0 {
			ret *= x
		}
		x *= x
		y >>= 1
	}
	return ret
}

// Min returns the minimize integer in the slice.
func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	res := a[0]
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

// Max returns the maximum integer in the slice.
func Max(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
