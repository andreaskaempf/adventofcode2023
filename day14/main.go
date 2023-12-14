// Advent of Code 2023, Day 14
//
// Given a square field of 'O' (round rocks), '#' (square rocks) and '.',
// move all the round rocks as much as possible to the top. For Part 2, do this
// four times (once in each direction) to make a cycle, and simulate 10e9
// cycles. For both parts, the answer is the score calculated as the sum of the
// number of round rocks below in each column, for each row. Part 2 cannot
// be calculated using brute force, but the pattern repeats after a while, so
// use this to quickly determine what the pattern would look like after many
// iterations.
//
// AK, 14 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Global variable with matrix of data
var rows [][]byte

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte("\n"))

	// Set this to true for part 2
	var part2 bool
	part2 = true

	// Part 1: roll up once, calculate score
	if !part2 {
		rollUp()
		fmt.Println("Part 1:", score()) // 136 sample, 109098 input
		return
	}

	// Part 2: first detect start and length of recurring cycles
	var hist []string
	var scores []int
	var cycleStart, cycleLen int
	var cycleFirst string
	for i := 0; cycleLen == 0; i++ {

		// Do the next cycle
		doCycle()

		// If this configuration has been encountered before,
		// detect start and length of cycle
		sb := toString(rows)
		if in(sb, hist) {
			if cycleStart == 0 { // beginning of cycle
				cycleStart = i
				cycleFirst = sb
			} else if sb == cycleFirst { // past end of first cycle
				cycleLen = i - cycleStart
			}
		}

		// Add block and its score to history
		hist = append(hist, sb)
		scores = append(scores, score())
	}

	// sample: 10, 7   input: 98, 17
	fmt.Println("Cycle starts at", cycleStart, ", length", cycleLen)

	// Use this to get the state after 1_000_000_000 cycles, and
	// the score for that state
	i := (1_000_000_000 - cycleStart - 1) % cycleLen
	ans := scores[cycleStart+i]
	fmt.Println("Part 2 (s/b 64, 100064):", ans) // 64 sample, 100064 input
}

// Create a string representation of the block
func toString(b [][]byte) string {
	s := ""
	for r := 0; r < len(b); r++ {
		if r > 0 {
			s += " "
		}
		s += string(b[r])
	}
	return s
}

// Each cycle tilts the platform four times so that the rounded rocks
// roll north, then west, then south, then east. So roll up four times,
// rotating each time.
func doCycle() {
	rollUp() // Roll north
	rows = rotate(rows)
	rollUp() // Roll west (left)
	rows = rotate(rows)
	rollUp() // Roll south (down)
	rows = rotate(rows)
	rollUp() // Roll east (right)
	rows = rotate(rows)
}

// Roll all stones vertically up, return true if anything was moved
func rollUp() {
	nc := len(rows[0])
	for r := 1; r < len(rows); r++ { // Start with top row, move down
		for c := 0; c < nc; c++ { // Each column
			if rows[r][c] == 'O' { // If it's an 'O', try to move it
				for y := r; y > 0; y-- {
					if rows[y-1][c] == '.' {
						rows[y-1][c] = 'O'
						rows[y][c] = '.'
					} else {
						break
					}
				}
			}
		}
	}
}

// Calculate score
func score() int {
	ans := 0
	nc := len(rows[0])
	for r := 0; r < len(rows); r++ { // each row
		for c := 0; c < nc; c++ { // each column
			if rows[r][c] == 'O' {
				ans += len(rows) - r
			}
		}
	}
	return ans
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
