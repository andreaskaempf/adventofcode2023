// Brute force solution used for Part 1 (actually a revised version that uses
// less memory)

package main

import (
	"fmt"
	"math"
)

// Process this line
func bruteForce(patt []byte, nums []int) int {

	// Count up the number of question marks, get positions
	var qq []int
	for i := 0; i < len(patt); i++ {
		if patt[i] == '?' {
			qq = append(qq, i)
		}
	}
	n := len(qq)                          // number of 1/0 bits
	iters := int(math.Pow(2, float64(n))) // number of iterations
	//fmt.Println("  ", n, "bits =>", iters, "iterations")

	// Count binary to consider all possibilities of q 1/0 bits
	bits := make([]int, n, n) // all digits initialized to zero
	done := false
	first := true
	p := make([]byte, len(patt), len(patt)) // buffer for copy of pattern
	ans := 0
	iter := 0
	for !done {

		// Show progress
		iter++
		if iter%10000000 == 0 {
			pcnt := float64(iter) / float64(iters) * 100.0
			fmt.Printf("\r   %.2f%%", pcnt)
		}

		// Add one to last digit, carry forward
		if first { // don't increment first time, so we include all zeros
			first = false
		} else {
			bits[n-1]++                  // increment last digit
			for i := n - 1; i > 0; i-- { // carry forward
				if bits[i] > 1 {
					bits[i] = 0
					bits[i-1]++
				}
			}
		}

		// Replace question marks with #/. corresponding to bits, see if
		// they match
		copy(p, patt) // make copy of pattern
		for i := 0; i < len(qq); i++ {
			q := qq[i]   // position of question mark
			b := bits[i] // corresponding bit
			if b == 1 {
				p[q] = '#'
			} else {
				p[q] = '.'
			}
		}

		// Check if matches
		if matches(p, nums) {
			ans++
		}

		// Not done if there are any zeros left
		done = true
		for i := 0; i < n; i++ {
			if bits[i] == 0 {
				done = false
				break
			}
		}
	}

	// Return the number of replacements that matched
	return ans
}

// Does this pattern match the numbers? I.e., do the number and lengths
// of contiguous blocks of '#' match the list of numbers?
func matches(patt []byte, nums []int) bool {

	// Summarize the sequences of '#'
	var blocks []int
	var b int // length of this block
	for i := 0; i < len(patt); i++ {
		c := patt[i]
		if c == '#' {
			b++ // exntend this block
		} else {
			if b > 0 {
				blocks = append(blocks, b)
			}
			b = 0
		}
	}
	if b > 0 {
		blocks = append(blocks, b)
	}

	// Matches if same
	if len(blocks) != len(nums) {
		return false
	}
	for i := 0; i < len(blocks); i++ {
		if blocks[i] != nums[i] {
			return false
		}
	}
	return true
}
