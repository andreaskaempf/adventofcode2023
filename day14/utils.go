// Utility functions for Advent of Code

package main

// Is element in a list?
func in[T int | float64 | byte | string](c T, s []T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

// Reverse a byte slice
func reverse(b []byte) []byte {
	r := make([]byte, len(b), len(b))
	for i := 0; i < len(b); i++ {
		r[i] = b[len(b)-i-1]
	}
	return r
}

// Get a column slice
func getCol(c int, b [][]byte) []byte {
	col := make([]byte, len(b), len(b))
	for r := 0; r < len(b); r++ {
		col[r] = b[r][c]
	}
	return col
}

// Transpose an array of arrays
func transpose(b [][]byte) [][]byte {
	t := make([][]byte, len(b[0]), len(b[0]))
	for c := 0; c < len(b[0]); c++ {
		t[c] = getCol(c, b)
	}
	return t
}
