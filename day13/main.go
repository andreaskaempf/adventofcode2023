// Advent of Code 2023, Day 10
//
//
//
// AK, 10 Dec 2023 (part 1 1:20)

package main

import (
	"fmt"
	//"strings"
	"bytes"
	"io/ioutil"
)

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	//fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	lines := bytes.Split(data, []byte("\n"))

	// Find each block and process it
	var ans1, ans2 int
	var b [][]byte // the current block
	for _, l := range lines {
		if len(l) == 0 { // blank line means end of a block
			ans1 += processBlock(b)
			ans2 += part2(b)
			b = [][]byte{}
		} else {
			b = append(b, l)
		}
	}
	if len(b) > 0 { // process the last block
		ans1 += processBlock(b)
		ans2 += part2(b)
	}
	fmt.Println("Part 1:", ans1) // 405 for sample, 35360 for input
	fmt.Println("Part 2:", ans2) // 400 for sample, ? for input

}

// Find the line of symmetry in a block, return the number of columns to left
// if vertical, or 100 X the number of rows above if horizontal
func processBlock(b [][]byte) int {

	//printBlock(b)

	// Look for vertical line, must go all the way to the edge
	var vc, nv int
	for c := 0; c < len(b[0])-1; c++ {
		found := true
		for d := 0; d < len(b[0]); d++ {
			if c-d < 0 || c+d+1 >= len(b[0]) {
				break
			}
			if !same(getCol(c-d, b), getCol(c+d+1, b)) {
				found = false
				break
			}
		}
		if found {
			nv++
			//fmt.Println(nv, "Vertical @ c =", c)
			if c+1 > vc {
				vc = c + 1
			}
		}
	}

	// Look for horizontal line
	var hr, nh int
	for r := 0; r < len(b)-1; r++ {
		found := true
		for d := 0; d < len(b); d++ {
			if r-d < 0 || r+d+1 >= len(b) {
				break
			}
			if !same(b[r-d], b[r+d+1]) {
				found = false
				break
			}
		}
		if found {
			nh++
			//fmt.Println(nh, "Horizontal @ r =", r)
			if r+1 > hr {
				hr = r + 1
			}
		}
	}

	if hr == 0 && vc == 0 {
		fmt.Println("*** No line(s) found:")
		printBlock(b)
	}
	if hr > 0 && vc > 0 {
		fmt.Println("*** Both horiz and vertical found:")
		printBlock(b)
	}
	return vc + 100*hr
}

// For part 2, alter characters to try to find another line
func part2(b [][]byte) int {

	// The initial score
	s0 := processBlock(b)

	// Try altering characters until you get a different non-zero score
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[0]); c++ {
			old := b[r][c]
			if old == '.' {
				b[r][c] = '#'
			} else {
				b[r][c] = '.'
			}
			s1 := processBlock(b)
			if s1 > 0 && s1 != s0 {
				return s1
			}
			b[r][c] = old
		}
	}
	fmt.Println("No answer found")
	return 0
}

// Get a column slice
func getCol(c int, b [][]byte) []byte {
	col := make([]byte, len(b), len(b))
	for r := 0; r < len(b); r++ {
		col[r] = b[r][c]
	}
	return col
}

// Print a block
func printBlock(b [][]byte) {
	for _, r := range b {
		fmt.Println(string(r))
	}
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
