// Advent of Code 2023, Day 18
//
// Given a set of instructions to draw a polygon consisting of just
// horizontal and vertical lines, count the number of points that are
// inside the polygon. For Part 2, instructions are revised to give a
// much bigger shape, that cannot be computed in memory. Did Part 1 using
// simple recursive flood fill, Part 2 using go-geom library.
//
// AK, 18 Dec 2023

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/twpayne/go-geom"
)

type Point struct {
	x, y int
}

var lines []string
var points map[Point]int

func main() {

	// Read the input file into a list of strings
	fname := "sample.txt"
	fname = "input.txt"
	lines = readLines(fname)

	// Part 1: Add up the number of points coloured
	fmt.Println("Part 1:", part1())

	// Part 2: Same, but use geometry
	fmt.Println("Part 2:", area(true))
}

// Part 1: interpret the instructions as direction, distance, color,
// draw & fill, return number of squares that are on
func part1() int {

	// Map the edge, starting from 0,0
	p := Point{} // current location 0,0
	points = map[Point]int{}
	for _, l := range lines {

		// Parse line
		parts := strings.Split(l, " ")
		dir := parts[0] // direction R/L/D/U
		n := atoi(parts[1])
		//color := parts[2]  // not used for Part 1

		// Draw edges of shape
		for i := 0; i < n; i++ {
			if dir == "R" {
				p.x++
			} else if dir == "L" {
				p.x--
			} else if dir == "U" {
				p.y--
			} else {
				p.y++
			}
			points[p] = 1
		}
	}

	// Fill the shape, using the simple recursive algorithm, and
	// return the number points filled including border
	fill1(4, 2)
	return len(points)
}

// Simple recursive algorithm to flood-fill the shape, starting with the
// given position, which must be inside shape (otherwise program will crash,
// so you know to try a different value).
// Works for part 1, but too memory intensive for part 2
func fill1(x, y int) {
	points[Point{x, y}] = 1 // fill current point
	if empty(x-1, y) {
		fill1(x-1, y)
	}
	if empty(x+1, y) {
		fill1(x+1, y)
	}
	if empty(x, y-1) {
		fill1(x, y-1)
	}
	if empty(x, y+1) {
		fill1(x, y+1)
	}
}

// Part 2: interpret the instructions from just the color,
// use different fill algorithm to avoid stack overflow
// NOT USED
func part2() int64 {

	// Map the edge, starting from 0,0
	fmt.Println("Creating map")
	p := Point{} // current location 0,0
	points = map[Point]int{}
	for _, l := range lines {

		// Parse direction and number from hex color, e.g., "(#1b58a2)"
		// The first five hexadecimal digits encode the distance in meters
		// as a five-digit hexadecimal number.
		// The last hexadecimal digit encodes the direction to dig:
		// 0 means R, 1 means D, 2 means L, and 3 means U.
		parts := strings.Split(l, " ")
		color := parts[2] // direction & number not used for Part 2
		ns := color[2:7]
		n, err := strconv.ParseInt(ns, 16, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		dir := int(color[7] - '0')

		// Draw edges of shape
		for i := 0; i < int(n); i++ {
			if dir == 0 { // right
				p.x++
			} else if dir == 2 { // left
				p.x--
			} else if dir == 3 { // up
				p.y--
			} else { // else 1 = down
				p.y++
			}
			points[p] = 1 // use 1 to designate edge (fill is 2)
		}
	}

	// Fill the shape
	//fill2(4, 2)

	// Return the number of points filled
	return fill4() //len(points)
}

// Get area of polygon outlined by instructions, using the Go-Geom library,
// used for Part 2
func area(part2 bool) int64 {

	// Collect the corner points that make up the polygon
	var x, y float64     // start at 0,0
	var boundary float64 // length of the border
	coords := []geom.Coord{}
	coords = append(coords, geom.Coord{0, 0})
	directions := []string{"R", "D", "L", "U"} // for Part 2
	for _, l := range lines {

		// Parse line, e.g., R 6 (#70c710)
		// For Part 1, direction and number are given (ignore colour).
		// For Part 2, the first five hex digits of the colour encode the
		// distance, last hex digit encodes the direction so that
		// 0 means R, 1 means D, 2 means L, and 3 means U.
		parts := strings.Split(l, " ")
		dir := parts[0] // direction R/L/D/U
		n := atof(parts[1])
		if part2 {
			color := parts[2] // not used for Part 1
			ns := color[2:7]
			n64, err := strconv.ParseInt(ns, 16, 64)
			if err != nil {
				fmt.Println("Hex converstion:", err.Error())
			}
			n = float64(n64)
			dir = directions[int(color[7]-'0')]
		}

		// Add a coordinate for this step
		if dir == "R" {
			x += n
		} else if dir == "L" {
			x -= n
		} else if dir == "U" {
			y -= n
		} else {
			y += n
		}
		coords = append(coords, geom.Coord{x, y})

		// Measure the length of the boundary around the shape, so we can
		// add 1/2 width padding to get the area including the outline
		boundary += n + .5
	}

	// Close the shape
	coords = append(coords, geom.Coord{0, 0})

	// Calculate the area
	sh := geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{coords})
	a := sh.Area()
	return int64(a + boundary/2 - 2)
}

// NOT USED
func fill4() int64 {

	// Get edge points into a list
	fmt.Println("Building list of points")
	pp := make([]Point, len(points))
	for p, _ := range points {
		pp = append(pp, p)
	}

	// Sort by y, x
	fmt.Println("Sorting")
	sort.Slice(pp, func(i, j int) bool {
		if pp[i].y == pp[j].y {
			return pp[i].x < pp[j].x
		} else {
			return pp[i].y < pp[j].y
		}
	})

	// Process each row: include any cells that are on, and treat
	// consecutive blocks of cells as toggling inside state
	fmt.Println("Counting")
	prevY := -9999999
	for _, p := range pp {
		if p.y != prevY {

		}
	}
	return 0
}

// NOT USED
func fill3() int64 {

	// Get min/max x and y
	var minX, minY, maxX, maxY int
	first := true
	fmt.Println("Map has", len(points), "points")
	for p, _ := range points {
		if first || p.x < minX {
			minX = p.x
		}
		if first || p.x > maxX {
			maxX = p.x
		}
		if first || p.y < minY {
			minY = p.y
		}
		if first || p.y > maxY {
			maxY = p.y
		}
		first = false
	}
	fmt.Printf("X = %d - %d, Y = %d - %d\n", minX, maxX, minY, maxY)

	// hit a wall:
	//   count it
	//   if !inside -> inside = true
	// else (not wall):
	//   if inside and prev char was a wall -> inside = false
	//   else if
	var ans int64
	for y := minY; y <= maxY; y++ {
		//if y%100 == 0 {
		pcnt := float64(y-minY) / float64(maxY-minY) * 100
		fmt.Printf("\r%.2f %% done", pcnt)
		//}
		inside := false
		for x := minX; x <= maxX; x++ {
			//if !empty(x, y) && (empty(x-1, y) || empty(x+1, y)) {
			if !empty(x, y) && empty(x-1, y) {
				inside = !inside
			}
			if inside || !empty(x, y) {
				//points[Point{x, y}] = 2
				ans++
			}
		}
	}
	return ans
}

// Is a position empty?
func empty(x, y int) bool {
	_, ok := points[Point{x, y}]
	return !ok
}

// For debugging
func draw() {
	for y := -1; y < 12; y++ {
		for x := -1; x < 12; x++ {
			fmt.Print(ifElse(empty(x, y), ".", "#"))
		}
		fmt.Print("\n")
	}
}
