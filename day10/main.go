// Advent of Code 2023, Day 10
//
// AK, 10 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

var lines [][]byte

type Point struct {
	x, y int
}

// Shortest distance to any point on path
var dist map[Point]int

// For Part 2, keep track of any position that has ever been visited
var everVisited map[Point]bool
var seq map[Point]int

func main() {

	// Read the input file (must have no empty lines at end)
	fname := "sample.txt"
	fname = "sample2a.txt"
	//fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	lines = bytes.Split(data, []byte("\n"))

	// Find starting position
	var start Point
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x] == 'S' {
				start = Point{x, y}
				break
			}
		}
	}
	fmt.Println("Start at", start)

	// Follow all paths from start, collect distances
	dist = map[Point]int{}         // Initialize shortest distance to any point on path
	visited := []Point{}           // Points already visited on this traversal
	everVisited = map[Point]bool{} // For part 2
	seq = map[Point]int{}
	explore(start, 0, visited) // Explore from start, get distances

	// Part 1: Find the maximum distance from start
	ans1 := 0
	for _, v := range dist {
		if v > ans1 {
			ans1 = v
		}
	}
	fmt.Println("Part 1:", ans1) // 6927

	// Part 2: count up points that are enclosed by path
	ans2 := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			p := Point{x, y}
			if !everVisited[p] && enclosed2(p) {
				ans2++
			}
		}
	}
	fmt.Println("Part 2 for", fname, ":", ans2) // Samples s/b 4, 8, 10
	//fmt.Println(seq)
}

// Explore all paths from this point, mark points as visited, and update
// global dictionary with shortest distances to each point
func explore(here Point, steps int, visited []Point) {

	// Mark this point as visited
	visited = append(visited, here)
	everVisited[here] = true
	_, ok := seq[here]
	if !ok {
		seq[here] = steps + 1
	}

	// Update distance to this point if shorter
	if dist[here] == 0 || steps < dist[here] {
		dist[here] = steps
	}

	// Explore each point from here that has not been explored
	for _, next := range nextSteps(here) {
		if !in(next, visited) { //visited[next] == 0 {
			explore(next, steps+1, visited) // pointer to visited?
		}
	}
}

// Return a list of the viable next steps from this position.
// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
func nextSteps(p Point) []Point {
	res := []Point{}

	// Above: bar, F, 7
	p1 := Point{p.x, p.y - 1}
	c := at(p1)
	if c == '|' || c == 'F' || c == '7' {
		res = append(res, p1)
	}

	// Below: vert bar, L, J
	p1 = Point{p.x, p.y + 1}
	c = at(p1)
	if c == '|' || c == 'L' || c == 'J' {
		res = append(res, p1)
	}

	// Left: dash, F, L
	p1 = Point{p.x - 1, p.y}
	c = at(p1)
	if c == '-' || c == 'F' || c == 'L' {
		res = append(res, p1)
	}

	// Right: dash, 7, J
	p1 = Point{p.x + 1, p.y}
	c = at(p1)
	if c == '-' || c == '7' || c == 'J' {
		res = append(res, p1)
	}

	return res
}

// Ray Casting algorithm, which involves casting a ray from the point in
// question in any direction and counting the number of times the ray intersects
// with the edges of the shape. If the number of intersections is odd, the point
// is inside the shape; if it is even, the point is outside the shape.
func enclosed2(p Point) bool {

	var vcuts, hcuts int

	// Count intersections to left or right
	// MAY BE FAILING BECAUSE IT TREATS CONTIGUOUS POINTS AS DIFFERENT, EVEN
	// IF THEY ARE ON THE SAME HORIZONTAL PATH. BUT YOU CAN'T JUST LOOK AT
	// THE POINTS, NEED TO CONSIDER IF TWO CONTIGUOUS POINTS WERE DRAWN NEXT
	// TO EACH OTHER, OR DURING A LOOP. CAN DO THIS BY KEEPING TRACK OF SEQUENCE
	// POINTS ARE VISITED, THEY ARE CONTIGUOUS IF ABS(DIFF) == 1.
	if p.x > 0 {
		hcuts = countIntersectionsWithPath(p, -1, 0)
	} else {
		hcuts = countIntersectionsWithPath(p, 1, 0)
	}

	// Count intersections above or below
	if p.y > 0 {
		vcuts = countIntersectionsWithPath(p, 0, -1)
	} else {
		vcuts = countIntersectionsWithPath(p, 0, 1)
	}

	// Return true if odd number of intersections
	fmt.Println(p, string(at(p)), hcuts, vcuts, hcuts%2, vcuts%2)
	return hcuts%2 != 0 //&& vcuts%2 != 0
}

func countIntersectionsWithPath(p Point, dx, dy int) int {

	// Move to the next position in given direction
	prev := p //Point{-1, -1}
	p1 := Point{p.x + dx, p.y + dy}
	var cuts int
	for at(p1) != 0 { // continue while within range

		// If we are on the path, increment counter,
		// unless previous point was also on path, and was drawn just before or after it
		if everVisited[p1] {
			if everVisited[prev] && abs(seq[p1]-seq[prev]) != 1 {
				cuts++
			}
		}
		prev = p1

		// Move to the next position in given direction
		p1 = Point{p1.x + dx, p1.y + dy}
	}

	return cuts
}

// Character at a point, 0 if out of bounds
func at(p Point) byte {
	if p.y < 0 || p.y >= len(lines) || p.x < 0 || p.x >= len(lines[0]) {
		return 0
	} else {
		return lines[p.y][p.x]
	}
}

// Is point in a list of points? (would be faster to use a map)
func in(p Point, pp []Point) bool {
	for _, x := range pp {
		if x == p {
			return true
		}
	}
	return false
}

// Absolute value
func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
