// Advent of Code 2023, Day 13
//
// Given a set of 2D fields consisting of '.' and '#' characters, find
// the row or column in each that gives a mirror reflection, i.e., left/right
// sides are mirror images, or top/bottom sides. For Part 2, flip every
// character on each field, to find a different reflection (ignoring the
// previous one).
//
// AK, 13 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Global variables with last vertical and horiz split, and flag
//
//	to ignore the last horiz or vertical split (for part 2)
var lastV, lastH int     // default zero
var ignoreLastSplit bool // default false

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	lines := bytes.Split(data, []byte("\n"))

	// Find each block and process it
	var ans1, ans2 int
	var b [][]byte // the current block
	for _, l := range lines {
		if len(l) == 0 { // blank line means end of a block
			ans1 += processBlock2(b)
			ans2 += part2(b)
			b = [][]byte{}
		} else {
			b = append(b, l)
		}
	}
	if len(b) > 0 { // process the last block
		ans1 += processBlock2(b)
		ans2 += part2(b)
	}

	// Show the answers
	fmt.Println("Part 1:", ans1) // 405 for sample, 35360 for input
	fmt.Println("Part 2:", ans2) // 400 for sample, 36755 for input
}

// For Part 2, alter each character on the shape, and find a different
// mirror line
func part2(b [][]byte) int {

	// Get the initial score, so that we can ignore this split when finding
	// the one with the character flipped
	ignoreLastSplit = false
	processBlock2(b)
	lv := lastV
	lh := lastH
	ignoreLastSplit = true

	// Try altering each character until you get a different vertical
	// or horizontal split
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[0]); c++ {

			// Flip this character between '.' and '#'
			old := b[r][c]  // so we can set it back
			if old == '.' { // alternate the character
				b[r][c] = '#'
			} else {
				b[r][c] = '.'
			}
			lastV = lv             // set the vertical or horizontal split position
			lastH = lh             // of the original split
			s1 := processBlock2(b) // find the new split
			b[r][c] = old          //set the flipped character back to what it was
			if s1 > 0 {            // found it
				return s1
			}
		}
	}

	// Should never happen
	fmt.Println("No answer found")
	return 0
}

// Find horizontal or vertical line of symmetry in the shape, and return
// the number of columns to left of it (if vertical) or 100 times the
// number of rows above it (if horizontal)
func processBlock2(b [][]byte) int {

	// Look at each col
	for c := 1; c <= len(b[0])-1; c++ {
		if ignoreLastSplit && c == lastV {
			continue
		}
		if isMirror(b, c) {
			lastV = c // remember position so we can ignore it in part 2
			lastH = -1
			return c
		}
	}

	// Transpose and do the same, to do all cols
	b = transpose(b)
	for c := 1; c <= len(b[0])-1; c++ {
		if ignoreLastSplit && c == lastH {
			continue
		}
		if isMirror(b, c) {
			lastV = -1
			lastH = c
			return c * 100
		}
	}

	// No mirror line found
	return 0
}

// Given a list of rows, and a column number, are the left
// and right sides of the rows mirror images?
func isMirror(b [][]byte, c int) bool {

	// Should never happen
	if len(b[0]) < 2 {
		fmt.Println("Too narrow")
		return false
	}

	// Check every row
	for r := 0; r < len(b); r++ {

		// Get left and right sides of the row around c, so c = 1
		// means the left side will be 1 char wide
		row := b[r]
		L := row[:c]
		R := row[c:]
		if len(L) == 0 || len(R) == 0 { // should never happen
			fmt.Println("Empty side:", L, R)
		}

		// Reverse the left side, so they are left justified toward the split
		L = reverse(L)

		// Trim to same length
		if len(L) > len(R) {
			L = L[:len(R)]
		} else if len(R) > len(L) {
			R = R[:len(L)]
		}

		// If left and right are not the same, this is not a mirror line
		if !same(L, R) {
			return false
		}
	}

	// If we get to here, it's a mirror line
	return true
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

// Print a block, for debugging
func printBlock(b [][]byte) {
	for _, r := range b {
		fmt.Println(string(r))
	}
}

// Compare two lists element-by-element, and report
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
