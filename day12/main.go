// Advent of Code 2023, Day 12
//
// Given a pattern of '#' and '.', and a list of numbers representing
// the lengths of blocks of '#', find the number of possible patterns
// that match the given list of numbers. For Part 2, the pattern is
// repeated five times, with '?' in between, and the list of numbers
// is also repeated five times, defeating a brute force solution.
//
// AK, 12 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Cached values, keys are strings to be immutable
var Cache map[string]int64

func main() {

	// Input file name
	fname := "sample.txt"
	fname = "input.txt" // uncomment to use input file

	// Set to true for part 2
	var part2 bool // default false
	part2 = true   // uncomment for part 2

	// Initialize the cache
	Cache = map[string]int64{}

	// Read the input file
	fmt.Println("Using", fname, ", part 2 =", part2)
	data, _ := ioutil.ReadFile(fname)
	rows := bytes.Split(data, []byte("\n"))

	// Process each row
	var ans int64
	var oldAns int
	for _, r := range rows {

		// Get pattern and list of numbers from row
		// e.g., ?###???????? 3,2,1
		parts := bytes.Split(r, []byte(" "))
		patt := parts[0]
		nums := parseList(string(parts[1]), ",")

		// For Part 2, pattern is five copies of itself concatenated,
		// with '?' in between, and list of numbers is five copies of itself
		if part2 {
			var patt1 []byte
			var nums1 []int
			for i := 0; i < 5; i++ {
				patt1 = append(patt1, patt...)
				if i < 4 {
					patt1 = append(patt1, '?')
				}
				nums1 = append(nums1, nums...)
			}
			patt = patt1
			nums = nums1
		}

		// Find the number of possible patterns (replacing each '?' with
		// either '.' or '#') that match the given list of '####' lengths
		ans += countMatches(patt, nums)

		// Brute force solution from Part 1, for old times' sake
		if !part2 {
			oldAns += bruteForce(patt, nums)
		}
	}

	// Sample: 21 for Part 1, 525152 for Part 2
	// Input: 7163 for Part 1, 17788038834112 for Part 2
	fmt.Println("Answer =", ans)
	fmt.Println("Brute force =", oldAns)
}

// Find the number of possible patterns (replacing each '?' with
// either '.' or '#') that match the given list of '####' lengths.
// Dynamic programming solution, since brute force did not work
// with the larger inputs for Part 2.
func countMatches(patt []byte, lengths []int) int64 {

	// Retrieve value for this pattern and list of lengths from cache if
	// available. Key needs to be a string to be immutable.
	k := fmt.Sprintf("%s %v", patt, lengths)
	v, ok := Cache[k]
	if ok {
		return v
	}

	// Special cases
	if len(patt) == 0 && len(lengths) == 0 { //If both pattern and list of lengths
		Cache[k] = 1 // are empty, there is one match
		return 1
	} else if len(patt) == 0 { // If just pattern is empty, but
		Cache[k] = 0 // there are still block lengths, no match
		return 0
	} else if len(lengths) == 0 { // If there are no more block lengths,
		if in('#', patt) { // fail if there are any more blocks
			Cache[k] = 0
			return 0
		}
		Cache[k] = 1 // Otherwise, there is one match
		return 1
	} else if len(patt) < sum(lengths) { // If remaining blocks
		Cache[k] = 0 // won't fit inside pattern, fail
		return 0
	}

	// Deal with start of pattern, either a '#', '?' or '.'
	if patt[0] == '#' { // start of a block

		// Block must not contain any periods
		block := lengths[0] // length of the next expected block
		if in('.', patt[:block]) {
			Cache[k] = 0
			return 0
		}

		// If the block extends beyond current block length, no match
		if block < len(patt) && patt[block] == '#' {
			Cache[k] = 0
			return 0
		}

		// Check the rest of the pattern, beyond the current block
		l := []byte{}            // empty string in case no more
		if block+1 < len(patt) { // rest of pattern
			l = patt[block+1:]
		}
		v := countMatches(l, lengths[1:])
		Cache[k] = v
		return v

	} else if patt[0] == '.' { // period, skip it
		v := countMatches(patt[1:], lengths)
		Cache[k] = v
		return v

	} else if patt[0] == '?' { // question mark: try replacing with both '.' and '#'
		l := makeCopy(patt)           // make a fresh copy
		l[0] = '#'                    // try replacing '?' with '#'
		v := countMatches(l, lengths) // matches assuming '?' is '#'
		l[0] = '.'                    // assume '?' is '.'
		v += countMatches(l, lengths) // matches assuming '?' is '.'
		Cache[k] = v
		return v
	} else {
		panic("Invalid character in pattern")
	}
}
