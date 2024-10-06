// Advent of Code 2023, Day 07
//
// Given some hands of 5 cards, sort them by poker rank. For part 2, replace
// J (joker) with option that yields the best score. Answer for both parts is
// the sumproduct of the rank (sorted sequence) times a given bid.
//
// AK, Rust version 6 Oct 2024

use std::fs;
//use std::collections::HashMap;
use std::cmp::Ordering;
use itertools::Itertools;

// A hand of cards and its bid (from input file)
struct Hand {
    cards: String,
    bid: usize,
}

fn main() {

    // Read input file
    let fname = "input.txt";
    let data = fs::read_to_string(fname).expect("Read error");

    // Convert data to list of hands
    let mut hands = Vec::<Hand>::new();
    for l in data.lines() {
        let parts: Vec<_> = l.split(" ").collect();
        if parts.len() == 2 {
            let c = String::from(parts[0]);
            let b = parts[1].parse::<usize>().unwrap();
            hands.push(Hand{cards: c, bid: b});
        }
    }
    println!("{} hands", hands.len());

    // Sort hands by ascending strength of hand
	hands.sort_by(compare_hands);

	// Answer is sum product of rank * bid
	let mut ans = 0;
	let mut rank = 0;
	for h in hands {
		rank += 1;
		ans += rank * h.bid;
		//println!("{}: bid {} * rank {} = {}", h.cards, h.bid, rank, h.bid * rank);
	}
	println!("Part 1 (s/b 253205868): {}", ans); // s/b 253907829 part 2
}

// Compare two hands
fn compare_hands(a: &Hand, b: &Hand) -> Ordering {

		// Compare by type of hand, using Joker replacement for Part 2
		let ta = hand_type(a.cards.as_str());
		let tb = hand_type(b.cards.as_str());
		if ta < tb {
			return Ordering::Less;
		} else if ta > tb {
            return Ordering::Greater;
        }

		// If same type, compare card-by-card
		for i in 0..a.cards.len() {
            let ca = a.cards.chars().nth(i).unwrap();
            let cb = b.cards.chars().nth(i).unwrap();
            let sa = strength(ca);
            let sb = strength(cb);
			if sa < sb { // TODO: why is this reversed, should not sa < sb be Less?
				return Ordering::Greater;
			} else if sa > sb {
				return Ordering::Less;
		    }
        }

		// Should never happen
		Ordering::Equal 
}

// Calculate the poker hand "type" score for a hand
// 7 = Five of a kind (e.g., AAAAA)
// 6 = Four of a kind (AA8AA)
// 5 = Full house (23332)
// 4 = Three of a kind (TTT98)
// 3 = Two pair (23432)
// 2 = One pair (A23A4)
// 1 = All different (23456)
fn hand_type(hand: &str) -> u32 {

	// Count up how many of each letter, create list of frequencies in descending order
    let freq = hand.chars().counts();
    let mut counts: Vec<_> = freq.values().collect();
    counts.sort();
    counts.reverse();

	// Return type, inferring the type of hand from the counts
	if *counts[0] == 5 {
		return 7; // five of a kind
	} else if *counts[0] == 4 {
		return 6; // four of a kind
	} else if *counts[0] == 3 && *counts[1] == 2 {
		return 5; // full house
	} else if *counts[0] == 3 {
		return 4; // three of a kind
	} else if *counts[0] == 2 && *counts[1] == 2 {
		return 3; // two pairs
	} else if *counts[0] == 2 {
		return 2; // one pair
	} else {
		return 1; // "high card", all different
	}
}

// Strength of a card
// TODO: for Part 2, 'J' should return maximum score
fn strength(card: char) -> usize {
	String::from("AKQJT98765432").find(card).unwrap()+1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_hand_types() {
		assert!(hand_type("AAAAA") == 7);
		assert!(hand_type("AA8AA") == 6);
		assert!(hand_type("23332") == 5);
		assert!(hand_type("TTT98") == 4);
		assert!(hand_type("23432") == 3);
		assert!(hand_type("A23A4") == 2);
		assert!(hand_type("23456") == 1);
    }

	#[test]
	fn test_strength() {
		assert!(strength('A') == 1);
		assert!(strength('K') == 2);
		assert!(strength('2') == 13);
	}
}