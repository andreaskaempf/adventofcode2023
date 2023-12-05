// Advent of Code 2023, Day 05
//
// Given a list of "seeds" (numbers), go through a series of transformations
// based on tables. Each table is a tuple (dst, src, n). If the number is
// within src .. src+n,  it is assigned to dst. If no dst is found, the dst
// is the same as src for that transformation. For Part 1, find the lowest
// final transformation. Part 2 is the same, but treat each pair of "seeds"
// as start and length of a range. Brute force is okay in Go, but
// would be faster with concurrency, memoization, or a smarter approach.
//
// AK, 5 Dec 2023 (1:20)

package main

import (
	"fmt"
	"strings"
)

// Global variables for the inputs
var seeds []int
var maps map[string][][]int
var sections []string

// For memoization
/*type Key struct {
	section string
	seed    int
}*/

//var cache map[Key]int

func main() {

	// Variables for the inputs
	maps = map[string][][]int{}
	var reading string // current map being read
	sections = []string{"seed-to-soil", "soil-to-fertilizer",
		"fertilizer-to-water", "water-to-light", "light-to-temperature",
		"temperature-to-humidity", "humidity-to-location"}
	//cache = map[Key]int{}

	// Read and parse the input file
	fname := "sample.txt"
	fname = "input.txt"
	for _, l := range readLines(fname) {
		if len(l) == 0 {
			// skip blank lines
		} else if len(seeds) == 0 { // First line is list of seed numbers
			seeds = parseList(l[7:]) // chop off leading "seeds: "
		} else if l[len(l)-1] == ':' { // Start new section
			reading = l[:len(l)-5]
			maps[reading] = [][]int{}
		} else { // read numbers
			maps[reading] = append(maps[reading], parseList(l))
		}
	}

	// Part 1: process each seed, find the lowest location
	lowest := -1
	for _, seed := range seeds {
		loc := processSeed(seed)
		if lowest == -1 || loc < lowest {
			lowest = loc
		}

	}
	fmt.Println("Part 1:", lowest) // s/b 424490994

	// Part 2: same, but treat seed numbers as ranges
	// TODO: this would be faster with memoization, concurrency,
	// or something smarter
	lowest = -1
	for i := 0; i < len(seeds); i += 2 {
		s1 := seeds[i]  // start of range
		n := seeds[i+1] // length of range
		fmt.Println("Seed", s1, "for length", n)
		for i := 0; i < n; i++ {
			loc := processSeed(s1 + i)
			if lowest == -1 || loc < lowest {
				lowest = loc
			}
		}
	}
	fmt.Println("Part 2:", lowest) // s/b 15290096
}

// Process one seed, find its ending location
func processSeed(seed int) int {

	// Go through each section
	src := seed
	for _, sect := range sections {

		// Check the cache
		/*key := Key{sect, src}
		if val, ok := cache[key]; ok {
			src = val
			continue
		}*/

		// Find the destination corresponding to the current source
		dst := -1
		for _, ranges := range maps[sect] {
			if src >= ranges[1] && src <= ranges[1]+ranges[2] {
				dst = ranges[0] + src - ranges[1]
				break
			}
		}

		// If not found, use the source value
		if dst == -1 {
			dst = src
		}

		// Add to cache
		// TODO: disabled, because uses too much memory, need to
		// shorten the section names
		//cache[key] = dst

		// Source of next transformation is the destination of this one
		src = dst
	}

	// Return the final destination
	return src
}

// Parse a list of numbers
func parseList(s string) []int {
	nums := []int{}
	for _, n := range strings.Split(s, " ") {
		nums = append(nums, atoi(n))
	}
	return nums
}
