// Advent of Code 2023, Day 01
//
// Description:
//
// AK, 1 Dec 2023

package main

import (
	"fmt"
	//"strings"
)

func main() {

	// Read the input file
	fname := "sample.txt"
	//fname = "input.txt"
	for _, l := range readLines(fname) {
		fmt.Println(l)
	}
}
