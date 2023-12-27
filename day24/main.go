// Advent of Code 2023, Day 24
//
// Given a list of stones, each with a position and velocity, find
// the number stones whose trajectories will intersect. For Part 2,
// find the position and velocity of a new stone, whose trajectory
// will intersect with that of every other stone. Part 2 done using
// constraint solvers (Centipede in Go and Z3 from Python).
//
// AK, 24 and 26 Dec 2023

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// One hailstone's position and velocity
type Stone struct {
	x, y, z    float64 // position
	vx, vy, vz float64 // velocity, basically the slope
}

// List of stones
var stones []Stone

func main() {

	// Read the input file into a list of "hailstones"
	fname := "sample.txt"
	//fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	for _, l := range strings.Split(string(data), "\n") {
		parts := strings.Split(l, "@")
		pp := strings.Split(parts[0], ",")
		vv := strings.Split(parts[1], ",")
		s := Stone{}
		s.x = atof(pp[0])
		s.y = atof(pp[1])
		s.z = atof(pp[2])
		s.vx = atof(vv[0])
		s.vy = atof(vv[1])
		s.vz = atof(vv[2])
		stones = append(stones, s)
	}

	fmt.Println("Part 1:", part1()) // 2, 27732
	fmt.Println("Part 2:", part2())
}

// Part 1: Look for intersections that will happen with an X and Y
// position each at least 200000000000000 and at most 400000000000000.
// Disregard the Z axis entirely.
func part1() int {

	var ans int
	var lower float64 = 7 // values for sample data
	var upper float64 = 27
	if len(stones) > 5 { // i.e., input data
		lower = 200000000000000 // values for input data
		upper = 400000000000000
	}
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(stones); j++ {
			if i < j {

				// Get points and calculate their intersection
				A := stones[i]
				B := stones[j]
				x, y := intersect(A, B)
				//fmt.Println(A, B, x, y)

				// No intersection found (TODO: more robust)
				if x == 0 && y == 0 {
					continue
				}

				// Skip intersection happened in the past, i.e. velocity
				// and slope to intersection are not same sign
				if pos(A.vx) != pos(x-A.x) || pos(B.vx) != pos(x-B.x) {
					continue
				}

				// Check that intersection points are within range
				if x >= lower && x <= upper && y >= lower && y <= upper {
					ans++
				}
			}
		}
	}
	return ans
}

// Find intersection of two X,Y lines, adapted from
// https://www.geeksforgeeks.org/program-for-point-of-intersection-of-two-lines/
func intersect(A, B Stone) (float64, float64) {

	// Get velocity (slope) ratios, return if parallel
	mA := A.vy / A.vx
	mB := B.vy / B.vx
	if abs(mA-mB) < .000001 {
		return 0, 0
	}

	// Calculate intersection points
	x := (mA*A.x - mB*B.x + B.y - A.y) / (mA - mB)
	y := (mA*mB*(B.x-A.x) + mB*A.y - mA*B.y) / (mB - mA)

	return x, y
}

// Part 2: find the position and velocity of another stone, in three
// dimensions, so that it collides with all the other stones. This
// is a constraint satisfaction problem, solved using a constraint
// solver (Centipede in Go and Z3 from Python).
func part2() int {

	A := stones[0]
	X := Stone{24, 13, 10, -3, 1, 2} // new stone
	var t float64 = 5                // arbitrary time

	// Position of first stone at this time
	xA := A.x + t*A.vx
	yA := A.y + t*A.vy
	zA := A.z + t*A.vz

	// Position of new stone at this time
	xX := X.x + t*X.vx
	yX := X.y + t*X.vy
	zX := X.z + t*X.vz

	// Show match
	fmt.Println(xA, yA, zA)
	fmt.Println(xX, yX, zX)

	// Find parameters of stone X (6 vals), such that
	// pos(X) == pos(A) at some time, for every stone A
	// Optimization/constraint formulation:
	// - x,y,z and vx,vy,vz for new stone are variables
	// - t values for each stone are variables (not used)
	// - constraints are
	//       x1 + vx1 * t1 == xX + vxX * t1
	//       y1 + vy1 * t1 == yX + vyX * t1
	//       z1 + vz1 * t1 == zX + vzX * t1
	// - Not really an objective function, except to make
	//   sure there is a t for every stone

	// Use Centipede constraint solver to solve this, also
	// see part2.py for Z3 version
	//part2a()

	return 0
}

// Is number positive?
func pos(n float64) bool {
	return n > 0
}

// Parse a float, show message and return -1 if error
func atof(s string) float64 {
	n, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		fmt.Println("Could not parse float:", s)
		n = -1
	}
	return float64(n)
}

// Flexible version of math.Abs
func abs[T int | int64 | float64](x T) T {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
