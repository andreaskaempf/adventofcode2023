// Advent of Code 2023, Day 08
//
// AK, 8 Dec 2023

package main

import (
	"fmt"
)

// Node, with its left and right
type Node struct {
	id, L, R string
}

func main() {

	// Read the input file into path and dict of nodes
	var path string
	nodes := map[string]Node{}
	fname := "sample3.txt"
	fname = "input.txt"
	for _, l := range readLines(fname) {
		if len(path) == 0 {
			path = l // First line is LRLR* path
		} else if len(l) > 0 { // e.g., QRX = (XNN, TCJ)
			n := Node{id: l[:3], L: l[7:10], R: l[12:15]}
			nodes[n.id] = n
		}
	}

	// Part 1: navigate path from AAA to ZZZ, count steps
	loc := "AAA"        // current location, start at AAA
	steps := 0          // number of steps so far
	pos := 0            // position in path
	_, ok := nodes[loc] // Skip part 1 if AAA not found
	if ok {
		for loc != "ZZZ" { // continue until we reach ZZZ
			n, _ := nodes[loc]    // get current node
			if path[pos] == 'L' { // Get next location
				loc = n.L
			} else {
				loc = n.R
			}
			steps++
			pos++
			if pos >= len(path) {
				pos = 0
			}
		}
		fmt.Println("Part 1:", steps)
	}

	// Part 2: same, simulataneously navigate all paths from xxA to xxZ
	locs := []string{}       // starting/current locations
	ends := map[string]int{} // corresponding ending locations
	for k := range nodes {
		if k[2] == 'A' {
			locs = append(locs, k)
			end := k[:2] + "Z"
			ends[end] = 1
		}
	}
	fmt.Println(locs, ends)

	// Navigate to subsequent nodes, until all on ending node
	steps = 0
	pos = 0
	for {

		// Count up how many destinations we have reached
		arrived := 0
		for i := 0; i < len(locs); i++ {
			_, isEnd := ends[locs[i]]
			if isEnd { //locs[i][2] == 'Z' && in(locs[i], ends) {
				arrived++
			}
		}
		if arrived > 2 { // occasional progress
			fmt.Println(steps, arrived, locs)
		}
		if arrived == len(locs) { // stop when all arrived
			break
		}

		// Move each location one iteration according to current position in
		// the route
		for i := 0; i < len(locs); i++ {
			n, _ := nodes[locs[i]] // get current node
			//assert(ok, locs[i]+" not found")
			if path[pos] == 'L' { // Get next location
				locs[i] = n.L
			} else {
				//assert(path[pos] == 'R', "Invalid path")
				locs[i] = n.R
			}
		}

		// Next position
		steps++
		pos++
		if pos >= len(path) {
			pos = 0
		}
	}
	fmt.Println("Part 2:", steps)
}
