// Advent of Code 2023, Day 14
//
//
//
// AK, 14 Dec 2023 (Part 1: 30 mins)

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

var rows [][]byte

// First occurrence of each pattern

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	//fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte("\n"))

	hist := map[string]int{}
	// Part 1: roll up once, calculate score
	//rollUp()
	//fmt.Println("Part 1:", score()) // 136 sample, 109098 input

	// Part 2:
	for i := 0; i < 1_000_000_000; i++ {

		sb := toString(rows)
		v, ok := hist[sb]
		if ok {
			fmt.Println("Cycle", i, "happened at", v)
		} else {
			hist[sb] = i
		}

		//fmt.Println("Cycle", i+1)
		doCycle()
		//printBlock(rows)
	}
	fmt.Println("Part 2:", score()) // 136 sample, 109098 input
}

// Create a string representation of the block
func toString(b [][]byte) string {
	s := ""
	for r := 0; r < len(b); r++ {
		s += string(b[r])
	}
	return s
}

// Each cycle tilts the platform four times so that the rounded rocks roll north,
// then west, then south, then east.
func doCycle() bool {

	moved := false

	// Roll north
	if rollUp() {
		moved = true
	}

	// Roll west (left)
	rows = rotate(rows)
	if rollUp() {
		moved = true
	}
	// Roll south (down)
	rows = rotate(rows)
	if rollUp() { // roll west
		moved = true
	}

	// Roll east (right)
	rows = rotate(rows)
	if rollUp() {
		moved = true
	}

	// Restore direction, return if anything was moved
	rows = rotate(rows)
	return moved
}

// Roll all stones vertically up, return true if anything was moved
func rollUp() bool {
	nc := len(rows[0])
	moved := false
	for r := 1; r < len(rows); r++ { // Start with top row, move down
		for c := 0; c < nc; c++ { // Each column
			if rows[r][c] == 'O' { // If it's an 'O', try to move it
				for y := r; y > 0; y-- {
					if rows[y-1][c] == '.' {
						rows[y-1][c] = 'O'
						rows[y][c] = '.'
						moved = true
					} else {
						break
					}
				}
			}
		}
	}
	return moved
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

// Rotate a matrix in-place 90 degrees right
func rotate(bb [][]byte) [][]byte {

	// Transpose
	bb = transpose(bb)

	// Reverse each row
	for r := 0; r < len(bb); r++ {
		bb[r] = reverse(bb[r])
	}
	return bb
}

func testRotate() {

	bb := [][]byte{[]byte("1..2"), []byte("...."), []byte("3..4")}
	printBlock(bb)

	// Rotate once
	bb = rotate(bb)
	fmt.Println("Rotated:")
	printBlock(bb)

	// Rotate three more times
	bb = rotate(bb)
	bb = rotate(bb)
	bb = rotate(bb)
	fmt.Println("Three more times:")
	printBlock(bb)
}
