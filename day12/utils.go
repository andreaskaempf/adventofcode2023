// Utility functions for Advent of Code

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse a list of integers
func parseList(s, delim string) []int {
	nums := []int{}
	for _, n := range strings.Split(s, delim) {
		nums = append(nums, atoi(n))
	}
	return nums
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

// Is element in a list?
func in[T int | float64 | byte | string](c T, s []T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

// Sum of a list
func sum[T int | float64](l []T) T {
	var y T
	for i := 0; i < len(l); i++ {
		y += l[i]
	}
	return y
}

// Make copy of a string of bytes
func makeCopy(s []byte) []byte {
	c := make([]byte, len(s), len(s))
	copy(c, s)
	return c
}
