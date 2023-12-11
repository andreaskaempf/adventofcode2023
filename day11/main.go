// Advent of Code 2023, Day 11
//
// Given a 2D field of 'galaxies' surrounded by space, first 'expand' this by
// doubling each row/column that is empty, and add up the manhattan distance
// between every pair. For Part 2, expand by factor of a million instead of
// doubling. Go's default int size is 64, enough to handle this.
//
// AK, 11 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// A point in 2D space
type Point struct {
	x, y int
}

func main() {

	// Increment to be used for "expanding the universe"
	increment := 1 // part 1
	//increment = 100 - 1 // for testing sample in part 2
	//increment = 1000000 - 1 // uncomment for part 2

	// Read the input file into array of byte vectors
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := bytes.Split(data, []byte("\n"))

	// Convert to a list of x,y positions
	G := []Point{}
	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(rows[0]); c++ {
			if rows[r][c] == '#' {
				G = append(G, Point{x: c, y: r})
			}
		}
	}
	//fmt.Println("Galaxies:", G)

	// Collect lists of rows and columns that have no hashes, to be expanded
	var xRows, xCols []int
	for r := 0; r < len(rows); r++ {
		empty := true
		for _, g := range G {
			if g.y == r {
				empty = false
				break
			}
		}
		if empty {
			xRows = append(xRows, r)
		}
	}
	for c := 0; c < len(rows[0]); c++ {
		empty := true
		for _, g := range G {
			if g.x == c {
				empty = false
				break
			}
		}
		if empty {
			xCols = append(xCols, c)
		}
	}

	// Expand the empty rows and columns
	for i := len(xRows) - 1; i >= 0; i-- { // in reverse order
		xr := xRows[i] // the row to expand
		for j := 0; j < len(G); j++ {
			if G[j].y >= xr {
				G[j].y += increment
			}
		}
	}
	for i := len(xCols) - 1; i >= 0; i-- { // in reverse order
		xc := xCols[i] // the col to expand
		for j := 0; j < len(G); j++ {
			if G[j].x >= xc {
				G[j].x += increment
			}
		}
	}
	//fmt.Println("Expanded:", G)

	// Find manhattan distance between every pair
	var ans int
	for i := 0; i < len(G); i++ {
		for j := 0; j < len(G); j++ {
			if j > i {
				g1 := G[i]
				g2 := G[j]
				ans += abs(g1.x-g2.x) + abs(g1.y-g2.y)
			}
		}
	}
	fmt.Println("Answer:", ans) // 9521776 for Part1, 553224415344 Part 2
}

// Flexible version of math.Abs
func abs[T int | int64 | float64](x T) T {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
