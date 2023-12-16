// Advent of Code 2023, Day 16
//
// Given square a field of characters, simulate the movement of a "beam
// of light" entering the field. Slashes and backslashes change the
// direction of travel 90 degrees, and dashes or vertical bars "split" the
// beam in two. "Energy" is defined as the number of cells in the field that
// are eventually  touched by the beam. For Part 1, simulate a single beam
// entering top left, and report the energy. For Part 2, try entering from
// every position on every edge, and report the highest energy found.
// Quite easy, but need to know when to stop, ignore beams that leave the field
// and avoid creating duplicates of beams that are identical to any that
// already exist.
//
// AK, 16 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Direction
const (
	up    = iota // 0
	right        // 1
	down         // 2
	left         // 3
)

// A beam of light
type Beam struct {
	p   Point // x,y position
	dir int   // direction
}

// A 2D point (required to use as map key)
type Point struct {
	x, y int
}

// The rows of data
var rows [][]byte
var nr, nc int

func main() {

	// Read the input file into a matrix, i.e., a list of byte vectors
	fname := "sample.txt"
	//fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte("\n"))
	nr = len(rows)
	nc = len(rows[0])

	// Part 1: start with a single beam entering top left, get how many
	// cells are eventually "energized"
	beam0 := Beam{Point{-1, 0}, right}
	e := findEnergy(beam0)
	fmt.Println("Part 1:", e) // 46, 8539

	// Part 2: try entering from every cell on each side, find the
	// highest "energy"
	best := 0
	for r := 0; r < nr; r++ { // each row

		// Enter from left
		e := findEnergy(Beam{Point{-1, r}, right})
		if e > best {
			best = e
		}

		// Enter from right
		e = findEnergy(Beam{Point{nc, r}, left})
		if e > best {
			best = e
		}
	}
	for c := 0; c < nc; c++ { // each column

		// Enter from above
		e := findEnergy(Beam{Point{c, -1}, down})
		if e > best {
			best = e
		}

		// Enter from below
		e = findEnergy(Beam{Point{c, nr}, up})
		if e > best {
			best = e
		}
	}

	fmt.Println("Part 2:", best) // 51, 8674
}

// Simulate the movement of one (later more as beams are split) light beam(s),
// each reflecting off / and \, and splitting when encountering - and |.
// Keep track of locations where light has passed through, ("energized").
func findEnergy(beam0 Beam) int {

	beams := []Beam{beam0}
	energized := map[Point]int{}
	var iter, prevEnergized, lastChange int
	for {

		// Run until number of energized cells stops changing
		iter++
		//fmt.Printf("Iteration %d: %d beams, %d energized\n", iter, len(beams), len(energized))

		// Move each beam, only keep if still in range
		var validBeams []Beam
		for i, b := range beams {
			p := move(b)
			invalid := p.x < 0 || p.x >= nc || p.y < 0 || p.y >= nr
			if !invalid {
				beams[i].p = p
				energized[p] = 1 // Update locations energized
				validBeams = append(validBeams, beams[i])
			}
		}
		beams = validBeams

		// Stop when no more beams energized for a long time
		if len(energized) != prevEnergized {
			lastChange = iter
		}
		prevEnergized = len(energized)
		if iter-lastChange > 100 { // no change for 100 steps
			break
		}

		// Split or reflect beams
		var newBeams []Beam // for new beams that get created
		for i, b := range beams {

			// Get character at this location (checks for out-of-bounds)
			c := rows[b.p.y][b.p.x]

			// Check for mirrors (change direction), or splitters (can create
			// another beam)
			if c == '/' {
				if b.dir == right {
					beams[i].dir = up
				} else if b.dir == left {
					beams[i].dir = down
				} else if b.dir == up {
					beams[i].dir = right
				} else { // dir == down
					beams[i].dir = left
				}
			} else if c == '\\' {
				if b.dir == right {
					beams[i].dir = down
				} else if b.dir == left {
					beams[i].dir = up
				} else if b.dir == up {
					beams[i].dir = left
				} else { // dir == down
					beams[i].dir = right
				}
			} else if c == '-' {
				if b.dir == up || b.dir == down {
					beams[i].dir = left
					newBeams = append(newBeams, Beam{b.p, right})
				}
			} else if c == '|' {
				if b.dir == left || b.dir == right {
					beams[i].dir = up
					newBeams = append(newBeams, Beam{b.p, down})
				}
			}
		}

		// Add any new beams to the main list
		for _, b := range newBeams {
			if !beamExists(b, beams) {
				beams = append(beams, b)
			}
		}
	}

	// Answer is the number of "energized" cells
	return len(energized)
}

// Move a beam in its current direction, return new location
func move(b Beam) Point {
	p := b.p
	if b.dir == up {
		p.y--
	} else if b.dir == down {
		p.y++
	} else if b.dir == left {
		p.x--
	} else {
		p.x++
	}
	return p
}

// Check if identical beam already exists in list
func beamExists(beam Beam, beams []Beam) bool {
	for _, b := range beams {
		if b == beam {
			return true
		}
	}
	return false
}
