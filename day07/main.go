// Advent of Code 2023, Day 07
//
// Given some hands of 5 cards, sort them by poker rank. For part 2, replace
// J (joker) with option that yields the best score. Answer for both parts is
// the sumproduct of the rank (sorted sequence) times a given bid.
//
// AK, 7 Dec 2023

package main

import (
	"fmt"
	"sort"
)

// A "hand" in the card game
type Hand struct {
	cards      string
	bid        int
	replJokers string
}

func main() {

	// Set this to true for Part 2
	part2 := false

	// Read the input file into a list of "hands"
	fname := "sample.txt"
	fname = "input.txt"
	hands := []Hand{}
	for _, l := range readLines(fname) {
		h := Hand{cards: l[:5], bid: atoi(l[6:])}
		h.replJokers = h.cards
		hands = append(hands, h)
	}

	// For part 2, replace each hand that has jokers, with its highest-ranking
	// equivalent, by trying all possible combinations
	if part2 {
		for i := 0; i < len(hands); i++ {
			combos := allPossible(hands[i].cards)
			bestScore := 0
			var bestHand string
			for _, h := range combos { // try each combination
				score := handType(h)   // get type (score)
				if score > bestScore { // remember best found
					bestScore = score
					bestHand = h
				}
			}
			hands[i].replJokers = bestHand // need the best hand for Part 2
		}
	}

	// Assign strength to each card, in this order (highest to lowest):
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
	// Note that J score is zero for part 1, but included for part 2
	cards := "AKQJT98765432"
	if part2 { // why?
		cards += "J"
	}
	strength := map[byte]int{}
	for i := 0; i < len(cards); i++ {
		strength[cards[i]] = len(cards) - i
	}

	// Sort them by ascending strength of hand
	sort.Slice(hands, func(i, j int) bool {

		// Compare by type of hand, using Joker replacement for Part 2
		tA := handType(hands[i].replJokers)
		tB := handType(hands[j].replJokers)
		if tA != tB {
			return tA < tB
		}

		// If same type, compare card-by-card
		A := hands[i].cards
		B := hands[j].cards
		for x := 0; x < 5; x++ {
			if strength[A[x]] != strength[B[x]] {
				return strength[A[x]] < strength[B[x]]
			}
		}
		fmt.Println("Ambiguous:", A, B) // should never happen
		return true
	})

	// Answer is sum product of rank * bid
	var ans int
	for i := 0; i < len(hands); i++ {
		rank := i + 1
		ans += rank * hands[i].bid
	}
	fmt.Println(ans) // 253205868 part 1,  s/b 253907829 part 2
}

// Calculate the poker hand "type" score for a hand
// 7 = Five of a kind (e.g., AAAAA)
// 6 = Four of a kind (AA8AA)
// 5 = Full house (23332)
// 4 = Three of a kind (TTT98)
// 3 = Two pair (23432)
// 2 = One pair (A23A4)
// 1 = All different (23456)
func handType(hand string) int {

	// Count up how many of each letter
	res := map[byte]int{}
	for i := 0; i < len(hand); i++ {
		res[hand[i]]++
	}

	// Turn into list of counts
	counts := []int{}
	for _, v := range res {
		counts = append(counts, v)
	}

	// Sort counts in descending
	sort.Slice(counts, func(i, j int) bool {
		return counts[j] < counts[i]
	})

	// Return type, inferring the type of hand from the counts
	if counts[0] == 5 {
		return 7 // five of a kind
	} else if counts[0] == 4 {
		return 6 // four of a kind
	} else if counts[0] == 3 && counts[1] == 2 {
		return 5 // full house
	} else if counts[0] == 3 {
		return 4 // three of a kind
	} else if counts[0] == 2 && counts[1] == 2 {
		return 3 // two pairs
	} else if counts[0] == 2 {
		return 2 // one pair
	} else {
		return 1 // "high card", all different
	}
}

// Create all possible combinations of a hand, replacing Jokers
// with non-Joker characters from the hand
func allPossible(hand string) []string {

	// Special case
	if hand == "JJJJJ" {
		return []string{"AAAAA"}
	}

	// Find all the characters used in the hand
	chars := []byte{} // non-J characters in the hand
	jokers := []int{} // positions of the jokers
	for i := 0; i < len(hand); i++ {
		c := hand[i]
		if c == 'J' { // joker found, save position
			jokers = append(jokers, i)
		} else if !in(c, chars) { // otherwise add character to set
			chars = append(chars, c)
		}
	}

	// If no Jokers found, return hand unchanged
	if len(jokers) == 0 {
		return []string{hand}
	}

	// Otherwise, create new versions of the string, replacing each Joker
	// with each possible character from the original string; we assume
	// that each J should be the same character, which seems to work
	results := []string{}
	for _, c := range chars {
		bb := []byte(hand)
		for i := 0; i < len(bb); i++ {
			if bb[i] == 'J' {
				bb[i] = c
			}
		}
		results = append(results, string(bb))
	}
	return results
}
