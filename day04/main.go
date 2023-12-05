// Advent of Code 2023, Day 04
//
// You are given a list of "cards", each with two lists of numbers separated
// by a vertical bar. Some of the numbers in the right list are also in the
// left list. For Part 1, calculate "points" for each card by adding up 1, 2, 4, ...
// for each match. For Part 2, make a copy of the next n cards, where n is the
// number of matches, and count up the cards at the end.
//
// AK, 4 Dec 2023 (1:20)

package main

import (
	"fmt"
	"math"
	"strings"
)

// A card, just keep track of number of matches and copies
type Card struct {
	matches int // number of matches
	copies  int // number of copies of this card
}

func main() {

	// Read and process cards from the input file
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	fname := "sample.txt"
	fname = "input.txt"
	//fname = "day04_big_input.txt"
	cards := []Card{} // list of cards
	for _, l := range readLines(fname) {

		// Extract numbers to left and right of bar
		parts := strings.Split(l, " ")
		var left, right []int
		doingRight := false
		for i := 1; i < len(parts); i++ {
			w := strings.TrimSpace(parts[i])
			if len(w) == 0 || w[len(w)-1] == ':' {
				continue
			}
			if w == "|" {
				doingRight = true
				continue
			} else if doingRight {
				right = append(right, atoi(w))
			} else {
				left = append(left, atoi(w))
			}
		}

		// Count up winning numbers on this card, i.e., those in right
		// list that are also in left list
		var matches int
		for _, n := range right {
			if in(n, left) {
				matches++
			}
		}

		// Create the card
		card := Card{matches: matches, copies: 1}
		cards = append(cards, card)
	}

	// Part 1: one point for first match, 2 for second, etc.
	var ans1, ans2 int
	for _, c := range cards {
		points := int(math.Pow(2, float64(c.matches-1)))
		ans1 += points
	}
	fmt.Println("Part 1:", ans1) // s/b 13 for sample, 22193 for input

	// Part 2: for each card, create copies of subsequent cards for each match,
	// then count up how many cards at the end
	for cn := 0; cn < len(cards); cn++ {
		card := cards[cn]
		for i := 1; i <= card.matches; i++ {
			if cn+i < len(cards) {
				cards[cn+i].copies += card.copies
			}
		}
	}
	for cn := 0; cn < len(cards); cn++ {
		ans2 += cards[cn].copies
	}
	fmt.Println("Part 2:", ans2) // s/b 30 for sample, 5625994 input
}
