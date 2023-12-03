// Advent of Code 2023, Day 03
//
// Given a surface of digits and special characters, find any numbers that
// are adjacent to special characters. For Part 2, find any pairs of numbers
// that are adjacent to the same asterisk.
//
// AK, 3 Dec 2023 (60 mins)

package main

import (
	"fmt"
)

// A row & column position (required because of compound map key)
type Position struct {
	r, c int
}

// Global dictionary of part number lists, mapped by the position of the "gear"
// the are adjacent to (required for Part 2)
var Gears map[Position][]int

func main() {

	// Read the input file into a list of strings
	fname := "sample.txt"
	fname = "input.txt"
	lines := readLines(fname)

	// Process each line, by looking for numbers (consecutive digits), and
	// the characters in their bounding box. For Part 1, just see if there
	// are any adjacent special characters. For Part 2, keep track of the
	// numbers that are found adjacent to any asterisks.
	var ans1, ans2 int           // the answers for parts 1 & 2, initialized to zero
	Gears = map[Position][]int{} // initialize global map for Part 2
	for ri, l := range lines {

		// Look for start of the next number. c0 will be the starting position
		// of a number, c1 one past the end of the number.
		var c0, c1 int    // both initialized to zero, thanks Go!
		for c0 < len(l) { // continue to past end of line

			// Skip if not a digit
			if !isDigit(l[c0]) {
				c0++
				continue
			}

			// Ok, start of number found, now find the end of the number
			for c1 = c0; c1 < len(l) && isDigit(l[c1]); c1++ {
			}

			// Extract the number found
			partNo := atoi(string(l[c0:c1]))

			// Part 1: add up part number if it is adjacent to any 'symbols'
			// (this also builds up a map of numbers found adjacent to each
			// asterisk, for Part 2)
			if adjacentToSymbols(lines, ri, c0, c1) {
				ans1 += partNo
			}

			// Move starting position past this number
			c0 = c1
		}

	}
	fmt.Println("Part 1:", ans1) // s/b 560670

	// For Part 2, look at each "gear" that was found, and if there were
	// two part numbers adjacent to it, add their product to the answer
	for _, partNos := range Gears {
		if len(partNos) == 2 {
			ans2 += partNos[0] * partNos[1]
		}
	}
	fmt.Println("Part 2:", ans2) // s/b 91622824
}

// Determine if number is adjacent to any symbols
func adjacentToSymbols(lines []string, ri, c0, c1 int) bool {
	var isAdjacent bool                     // initialized to false
	for r := max(ri-1, 0); r <= ri+1; r++ { // row above and below
		if r >= len(lines) {
			break
		}
		for c := max(c0-1, 0); c <= c1 && c < len(lines[r]); c++ {

			// If character at this position is a symbol, we're true for Part 1
			ch := lines[r][c]
			if isSymbol(ch) {
				isAdjacent = true
			}

			// If the symbol is a star, add this number to the list of
			// numbers adjacent this 'gear' (for Part 2)
			if ch == '*' {
				l := lines[ri]
				partNo := atoi(string(l[c0:c1]))
				pos := Position{r, c}
				Gears[pos] = append(Gears[pos], partNo)
			}
		}
	}
	return isAdjacent // return value is used for Part 1
}

// A character is a symbol if it's not a digit or period
func isSymbol(c byte) bool {
	return c != '.' && !isDigit(c)
}

// Is this character a digit?
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
