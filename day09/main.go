// Advent of Code 2023, Day 09
//
// Given a list of numbers, repeatedly get the deltas between each successive
// number. Then, use the delta to extrapolate back upwards to the previous
// row. For Part 1, extrapolate at the end of each row, for Part 2, before
// the beginning. Answers are the sums of the extrapolations.
//
// AK, 9 Dec 2023

package main

import "fmt"

func main() {

	// Read the input file into lists of numbers
	fname := "sample.txt"
	fname = "input.txt"
	var data [][]int
	for _, l := range readLines(fname) {
		data = append(data, parseList(l))
	}

	// Process each row and extrapolate the next number on the first row,
	// at end for Part 1, before beginning for Part 2. Add up extrapolations
	// to get answers for parts 1 and 2.
	var ans1, ans2 int
	for _, row := range data {

		// Get successive lists of deltas for this row
		deltas := [][]int{}
		dd := getDeltas(row)
		for len(dd) > 0 {
			deltas = append(deltas, dd)
			dd = getDeltas(dd)
		}

		// Now extrapolate each row from the bottom up
		var x1, x2 int
		for i := len(deltas) - 1; i >= 0; i-- {
			dd := deltas[i]
			x1 += dd[len(dd)-1] // part 1
			x2 = dd[0] - x2     // part 2
		}

		// Extrapolate first row, add to answers
		ans1 += row[len(row)-1] + x1
		ans2 += row[0] - x2

	}
	fmt.Println("Part 1:", ans1) // 1782868781
	fmt.Println("Part 2:", ans2) // 1057
}

// Get list of deltas for a list of numbers
func getDeltas(nums []int) []int {
	deltas := make([]int, len(nums)-1, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		deltas[i-1] = nums[i] - nums[i-1]
	}
	return deltas
}
