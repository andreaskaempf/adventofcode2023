// Advent of Code 2023, Day 22
//
// Given a set of long bricks, each defined by two points in 3d space,
// determine which bricks can be removed without causing any other
// bricks to fall. For Part 2, determine how many bricks can be
// removed without causing any other bricks to fall. Fairly straightforward
// simulation, not very efficient.
//
// AK, 27 Dec 2023

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Positions of one brick
type Brick struct {
	start, end Point
	removed    bool // for each testing of removal
}

// A point in 3d space
type Point struct {
	x, y, z int
}

func main() {

	// Read the input file into a list of "bricks"
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	var bricks []Brick
	for _, l := range strings.Split(string(data), "\n") {
		parts := strings.Split(l, "~") // e.g., 1,0,1~1,2,1
		a := strings.Split(parts[0], ",")
		b := strings.Split(parts[1], ",")
		B := Brick{}
		B.start = Point{atoi(a[0]), atoi(a[1]), atoi(a[2])}
		B.end = Point{atoi(b[0]), atoi(b[1]), atoi(b[2])}
		bricks = append(bricks, B)
	}
	fmt.Println(len(bricks), "bricks read")

	// Part 1: try removing each brick, see if it affects any other bricks
	fmt.Println("Doing initial jiggle")
	bb, _ := jiggle(bricks) // let all bricks fall to bottom
	var part1, part2 int
	for i := 0; i < len(bb); i++ {
		fmt.Print("Brick ", i, " of ", len(bb))
		bb[i].removed = true     // remove this brick
		bb2, moved := jiggle(bb) // jiggle, see if anything moved
		if !moved {              // if nothing moved, okay to remove this
			fmt.Println(" can be safely removed")
			part1++
		} else { // Otherwise, cannot remov it
			fmt.Println(" cannot be removed")
		}
		numberMoved := diff(bb, bb2) // how many other bricks have moved?
		if numberMoved > 0 {
			fmt.Printf("  %d other bricks have moved\n", numberMoved)
			part2 += numberMoved
		}
		bb[i].removed = false // put the brick back
	}
	fmt.Println("Part 1 (5, 485):", part1)
	fmt.Println("Part 2 (7, 74594):", part2)
}

// "Jiggle" the space so all bricks fall to the lowest possible position,
// i.e., move down vertically, but don't intersect with anything else.
// Returns revised list of brick positions, and true if any bricks were moved.
// This is very slow, might be faster if we sort bricks by bottom location
// and start at the bottom, and allow bricks to fall by more than one step
// at a time.
func jiggle(bricks []Brick) ([]Brick, bool) {

	// Make a copy of the bricks
	bb := make([]Brick, len(bricks), len(bricks))
	copy(bb, bricks)

	// Iterate until no more bricks can be moved down
	anyMoved := false
	for {

		moved := false // any bricks moved in this iteration?
		for i := 0; i < len(bb); i++ {

			// Can't move brick if already at bottom, or removed
			b := bb[i]
			if b.removed || min(b.start.z, b.end.z) == 1 {
				continue
			}

			// Move the brick down one step
			bb[i].start.z--
			bb[i].end.z--

			// If any collisions, move brick back, otherwise leave it
			// and note that we have moved something
			space := getSpace(bb, true)
			if len(space) == 0 { // i.e., overlap found
				bb[i].start.z++
				bb[i].end.z++
			} else {
				moved = true
				anyMoved = true
			}
		}

		// Stop if no more bricks moved this iteration
		if !moved {
			break
		}
	}

	// Return updated list of bricks, and whether any bricks were moved
	return bb, anyMoved
}

// Return a map of space positions that are filled by a brick,
// numbers > 1 indicate overlap. If stopIfOverlaps is true,
// aborts if an overlap is found, and return empty map.
func getSpace(bricks []Brick, stopIfOverlaps bool) map[Point]int {

	points := map[Point]int{}
	for _, b := range bricks {

		// Ignore if flagged as removed
		if b.removed {
			continue
		}

		// Get min/max in each dimension
		x0 := min(b.start.x, b.end.x)
		x1 := max(b.start.x, b.end.x)
		y0 := min(b.start.y, b.end.y)
		y1 := max(b.start.y, b.end.y)
		z0 := min(b.start.z, b.end.z)
		z1 := max(b.start.z, b.end.z)

		// Increment the counter in each position
		for x := x0; x <= x1; x++ {
			for y := y0; y <= y1; y++ {
				for z := z0; z <= z1; z++ {
					points[Point{x, y, z}]++
					if stopIfOverlaps && points[Point{x, y, z}] > 1 {
						return map[Point]int{}
					}
				}
			}
		}
	}

	// Return map of points occouped
	return points
}

// For part 2, compare two stacks of bricks, and report
// how many have changed in position
func diff(bb1, bb2 []Brick) int {
	if len(bb1) != len(bb2) {
		panic("diff() received different length lists")
	}
	var ans int
	for i := 0; i < len(bb1); i++ {
		if bb1[i] != bb2[i] {
			ans++
		}
	}
	return ans
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
