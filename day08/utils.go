// Utility functions for Advent of Code

package main

import (
	"io/ioutil"
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

// Panic if a test condition is not true
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
