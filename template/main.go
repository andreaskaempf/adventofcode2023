// Advent of Code 2023, Day 10
//
//
//
// AK, 10 Dec 2023

package main

import (
	"fmt"
	//"strings"
	"bytes"
	"io/ioutil"
)

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	//fname = "input.txt"
	// lines := readLines(fname)  // for strings
	data, _ := ioutil.ReadFile(fname)
	lines := bytes.Split(data, []byte("\n"))
	for _, l := range lines {
		fmt.Println(string(l))
	}

}
