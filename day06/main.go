// Advent of Code 2023, Day 06
//
// Simulate a race, in which you hold down a button to recharge a motor,
// which lets the vehicle go faster. Distance is the speed muliplied by
// remaining time. Find the number of ways you can hold down the button, so
// as to exceed the previous maximum distance given. For part 2, just one
// race, but much bigger number.
//
// AK, 6 Dec 2023 (8 mins Part 1,17 mins total)

package main

import (
	"fmt"
)

func main() {

	// Sample data (hard coded, no input files)
	Time := []int{7, 15, 30}
	Distance := []int{9, 40, 200}

	// Use input data
	if true {
		Time = []int{60, 80, 86, 76}
		Distance = []int{601, 1163, 1559, 1300}
	}

	// Part 1: For each race, count up the number of ways you can hold down
	// the button, so as to exceed the given maximum distance, and multiply
	// these counts together
	ans1 := 1                        // answer is product, so start with 1
	for i := 0; i < len(Time); i++ { // each race
		T := Time[i]
		D := Distance[i]
		var better int
		for t := 1; t < T; t++ {
			speed := t
			dist := (T - t) * speed
			if dist > D {
				better++
			}
		}
		ans1 *= better
	}
	fmt.Println("Part 1:", ans1) // s/b 1155175

	// Part 2: same idea, but one very long race, time and disance obtained
	// by concatenating the times and distances of the sample values
	T := 71530 // sample values
	D := 940200
	if true { // input values
		T = 60808676
		D = 601116315591300
	}
	better := 0
	for t := 1; t < T; t++ {
		speed := t
		dist := (T - t) * speed
		if dist > D {
			better++
		}
	}
	fmt.Println("Part 2:", better) // s/b 35961505
}
