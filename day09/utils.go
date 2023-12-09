// Utility functions for Advent of Code

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Read lines from the input file, remove any blank lines at end
func readLines(filename string) []string {

	// Read data
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Split into lines
	lines := strings.Split(string(data), "\n")

	// Remove any trailing blank lines
	for len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	return lines
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

// Parse a list of integers
func parseList(s string) []int {
	nums := []int{}
	for _, n := range strings.Split(s, " ") {
		nums = append(nums, atoi(n))
	}
	return nums
}
