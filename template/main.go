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
	rows := bytes.Split(data, []byte("\n"))

	// Remove last row if empty
	if len(rows[len(rows)-1]) == 0 {
		rows = rows[:len(rows)-1]
	}

	for _, l := range rows {
		fmt.Println(string(l))
	}

}
