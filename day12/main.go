// Advent of Code 2023, Day 12
//
// AK, 12 Dec 2023 (Part 1: 6:05-8:12)

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {

	// Read the input file
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := bytes.Split(data, []byte("\n"))

	// Process each row
	var ans1 int
	for _, r := range rows {

		// Get pattern and list of numbers from row
		// e.g., ?###???????? 3,2,1
		parts := bytes.Split(r, []byte(" "))
		patt := parts[0]
		nums := parseList(string(parts[1]), ",")

		// For Part 2, pattern is five copies of itself concatenated,
		// and same for numbers
		// Tempting to use brute force, but consider another way ...

		// Part 1: How many ways can you set '?' to # or . so that the groupings of
		// contiguous hashes match the list of numbers?
		// Need to split up the pattern to avoid running out of memory
		fmt.Println(string(patt), nums)
		middle := len(patt) / 2
		c1 := getCombos(patt[:middle])
		c2 := getCombos(patt[middle:])
		for _, h1 := range c1 {
			for _, h2 := range c2 {
				both := append(h1, h2...)
				if matches(both, nums) {
					ans1++
				}
			}
		}
	}
	fmt.Println("Part 1:", ans1)
}

// Does this pattern match the numbers? I.e., do the number and lengths
// of contiguous blocks of '#' match the list of numbers?
func matches(patt []byte, nums []int) bool {

	// Summarize the sequences of '#'
	parts := bytes.Split(patt, []byte("."))
	blocks := []int{}
	for i := 0; i < len(parts); i++ {
		n := len(parts[i])
		if n > 0 {
			blocks = append(blocks, n)
		}
	}

	// Matches if same
	return same(blocks, nums)
}

// Get all possible combinations of a pattern, replacing
// '?' with '.' or '#'
func getCombos(patt []byte) [][]byte {

	res := [][]byte{}
	res = append(res, patt)
	for i := 0; i < len(patt); i++ {

		// Only look at positions that have '?'
		if patt[i] != '?' {
			continue
		}

		// Make list of copies, replacing this position
		res1 := [][]byte{}
		for r := 0; r < len(res); r++ {

			// Make a copy, assuming this ? is a hash
			p1 := make([]byte, len(patt), len(patt))
			copy(p1, res[r])
			p1[i] = '#'
			res1 = append(res1, p1)

			// Make a copy, assuming this ? is a hash
			p1 = make([]byte, len(patt), len(patt)) // reusable buffer
			copy(p1, res[r])
			p1[i] = '.' // Assume this ? is a period
			res1 = append(res1, p1)
		}

		// Add replacements to original list
		res = append(res, res1...)
	}

	// Copy the list, removing any patterns that have ?
	res1 := [][]byte{}
	for i := 0; i < len(res); i++ {
		if !in('?', res[i]) {
			res1 = append(res1, res[i])
		}
	}

	// Return the pruned list
	return res1
}
