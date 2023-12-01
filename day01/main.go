// Advent of Code 2023, Day 01
//
// Find digits in a string, combine first and last digit on each line to
// make a number, and add these up. For part 2, look for embedded names
// of numbers as well, complicated as last & first letters are shared
// if they are the same.
//
// AK, 1 Dec 2023

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Process each line
	fname := "sample.txt"
	fname = "sample2.txt"
	fname = "input.txt"
	var ans1, ans2, calib int
	for _, l := range readLines(fname) {

		// Part 1: combine first and last digits on each line, add them up
		digs := getDigits(l)
		if len(digs) > 0 { // sample2 fails on some lines
			calib = digs[0]*10 + digs[len(digs)-1]
			ans1 += calib
		}

		// Part 2: same, but first look for digit names embedded in strings
		l2 := replaceNums(l)
		digs = getDigits(l2)
		calib = digs[0]*10 + digs[len(digs)-1]
		ans2 += calib
		//fmt.Println(l, l2, digs, calib)
	}

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}

// Get all digits from a string, into a list of ints
func getDigits(s string) []int {
	res := []int{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			res = append(res, int(c-'0'))
		}
	}
	return res
}

// Replace embedded digit names with digits
func replaceNums(s string) string {

	// Replace some tricky combinations of numbers that share same first
	// and last letter
	s = strings.ReplaceAll(s, "oneight", "oneeight")
	s = strings.ReplaceAll(s, "twone", "twoone")
	s = strings.ReplaceAll(s, "threeight", "threeeight")
	s = strings.ReplaceAll(s, "fiveight", "fiveeight")
	s = strings.ReplaceAll(s, "sevenine", "sevennine")
	s = strings.ReplaceAll(s, "eightwo", "eighttwo")
	s = strings.ReplaceAll(s, "nineight", "nineeight")

	// Replace individual digit names with digits
	nums := []string{"one", "two", "three", "four", "five", "six", "seven",
		"eight", "nine"}
	for i, n := range nums {
		dig := fmt.Sprintf("%d", i+1)
		s = strings.ReplaceAll(s, n, dig)
	}
	return s
}
