// Utility functions for Advent of Code

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Read lines from the input file, remove any blank lines at end
func readLines(filename string) []string {

	// Read data
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Split into lines
	lines := strings.Split(string(data), "\n")

	// Remove any trailing blank lines
	for len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	return lines
}

// Parse an integer, show message and return -1 if error
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse integer:", s)
		n = -1
	}
	return int(n)
}

// Parse a 64-bit integer, show message and return -1 if error
func atoi64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse integer:", s)
		n = -1
	}
	return n
}

// Parse a float, show message and return -1 if error
func atof(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Could not parse float:", s)
		n = -1
	}
	return float64(n)
}

// Parse a list of integers
func parseList(s string) []int {
	nums := []int{}
	for _, n := range strings.Split(s, " ") {
		nums = append(nums, atoi(n))
	}
	return nums
}

// Maximum of a list (of ints, floats, or strings, using generics)
func maxList[T int | float64 | string](l []T) T {
	var y T
	for i := 0; i < len(l); i++ {
		if i == 0 || l[i] > y {
			y = l[i]
		}
	}
	return y
}

// Minimum of a list (of ints, floats, or strings, using generics)
func minList[T int | float64 | string](l []T) T {
	var y T
	for i := 0; i < len(l); i++ {
		if i == 0 || l[i] < y {
			y = l[i]
		}
	}
	return y
}

// Sum of a list
func sum[T int | float64](l []T) T {
	var y T
	for i := 0; i < len(l); i++ {
		y += l[i]
	}
	return y
}

// Intersection of two lists
func intersection[T int | float64 | byte | string](a, b []T) []T {
	res := []T{}
	for i := 0; i < len(a); i++ {
		if in(a[i], b) {
			res = append(res, a[i])
		}
	}
	return res
}

// Union of two lists
func union[T int | float64 | byte | string](a, b []T) []T {
	res := []T{}
	copy(res, a) // warning: this will include duplicates in list a
	for i := 0; i < len(b); i++ {
		if !in(b[i], res) {
			res = append(res, b[i])
		}
	}
	return res
}

// Is element in a list?
func in[T int | float64 | byte | string](c T, s []T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

// Shallow compare two lists element-by-element, and report
// if they are the same
func same[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Flexible version of math.Abs
func abs[T int | int64 | float64](x T) T {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Simple inline if-then-else
func ifElse[T int | float64 | byte | string](cond bool, a, b T) T {
	if cond {
		return a
	} else {
		return b
	}
}

// Panic if a test condition is not true
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

// Make copy of a string of bytes
func makeCopy(s []byte) []byte {
	c := make([]byte, len(s), len(s))
	copy(c, s)
	return c
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

// Rotate a matrix 90 degrees right
func rotate(bb [][]byte) [][]byte {

	// Transpose
	bb = transpose(bb)

	// Reverse each row
	for r := 0; r < len(bb); r++ {
		bb[r] = reverse(bb[r])
	}
	return bb
}
