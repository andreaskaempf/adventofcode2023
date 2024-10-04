// Advent of Code 2023, Day 04
//
// You are given a list of "cards", each with two lists of numbers separated
// by a vertical bar. Some of the numbers in the right list are also in the
// left list. For Part 1, calculate "points" for each card by adding up 1, 2, 4, ...
// for each match. For Part 2, make a copy of the next n cards, where n is the
// number of matches, and count up the cards at the end.
//
// AK, Rust version 3/4 Oct 2024

use std::fs;
use std::collections::HashSet;

// Structure for a card
struct Card {
    matches: u32,
    copies: u32,
}

fn main() {

    // Read file into a list of strings
    let fname = "input.txt";
    let data = fs::read_to_string(fname)
        .expect("Cannot read");
    let lines: Vec<_> = data.lines().collect();

    // Process each card and create a list of cards, also tally points
    // for Part 1
    let mut cards = Vec::<Card>::new();
    let mut part1: u32 = 0;
    for l in lines {

        // Ignore blank lines
        if l.len() == 0 {
            break
        }

        // Only interested in stuff after colon
        let colon = l.find(":").unwrap() + 2;
        let l1 = &l[colon..];
    
        // Left and right hand sides are lists of numbers separated by bar
        let parts: Vec<&str> = l1.split("|").collect();
        let winning = list_to_set(parse_ints(parts[0]));
        let have = list_to_set(parse_ints(parts[1]));

        // Part 1: get 2^n points, where n is the number cards you have that are winning
        let mut pts: u32 = 0;
        let isect = have.intersection(&winning).count() as u32;
        if isect > 0 {
            pts = u32::pow(2, isect - 1);
        }
        part1 += pts;

        // Create card and add to list, for Part 2
        let c = Card{matches: isect, copies: 1};
        cards.push(c);
    }

    // Part 1 answer
    println!("Part 1 (s/b 22193) = {}", part1);

    // Part 2: for each card, create copies of subsequent cards for each match
    // (just update counters, don't need to recreate cards), then count up how 
    // many cards at the end
	for cn in 0..cards.len() {
		let card = &cards[cn];
        let ncopies = card.copies; // required here to avoid mutable error
        let nmatches = card.matches+1;
		for i in 1..nmatches {
            let ii: usize = i.try_into().unwrap(); // WTF!!!
			if cn+ii < cards.len() {
				cards[cn+ii].copies += ncopies; // can't use card.matches directly!!!
			}
		}
	}

    // Count up all the copies
    let part2: u32 = cards.iter().map(|c| c.copies).sum();
	println!("Part 2 (s/b 5625994) = {}", part2);
       
}

// Convert list of numbers into a set
// TODO: is there a one-line version of this?
fn list_to_set(l: Vec<u32>) -> HashSet<u32> {
    let mut s = HashSet::<u32>::new();
    for x in l {
        s.insert(x);
    }
    s
}

// Parse a space-separated list of integers
fn parse_ints(s: &str) -> Vec<u32> {
    s.split(" ")
        .map(|n| n.trim())
        .filter(|s| s.len() > 0)
        .map(|s| s.parse::<u32>().unwrap())
        .collect()
}
