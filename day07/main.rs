// Advent of Code 2023, Day 07
//
// Given some hands of 5 cards, sort them by poker rank. For part 2, replace
// J (joker) with option that yields the best score. Answer for both parts is
// the sumproduct of the rank (sorted sequence) times a given bid.
//
// AK, Rust version 6-8 Oct 2024

use std::cmp::Ordering;
use std::collections::HashMap;
use std::fs;

// A hand of cards and its bid (from input file)
struct Hand {
    cards: String,
    bid: usize,
}

// If you change or refer to this, needs to be inside 'unsafe' block
static mut PART2: bool = false;

fn main() {
    // Read input file
    let fname = "sample.txt";
    let mut hands = read_hands(fname);
    println!("{} hands", hands.len());

    // Part 1: get sum-product of ranked bids
    let mut ans = compute_answer(&mut hands);
    println!("Part 1 (s/b 6440 / 253205868): {}", ans);

    // Part 2: replace Jokers in each hand to maximize score, then try again
    unsafe {
        PART2 = true; // used to adjust the strength of Jokers
    }
    for i in 0..hands.len() {
        hands[i].cards = replace_jokers(&mut hands[i].cards);
    }
    ans = compute_answer(&mut hands);
    println!("Part 2 (s/b 5905 / 253907829): {}", ans);
}

// Compute the answer, sum-product of rank * bid
fn compute_answer(hands: &mut Vec<Hand>) -> usize {

    // Sort hands by ascending strength of hand (implies rank)
    hands.sort_by(compare_hands);

    // Answer is sum product of rank * bid
    let mut ans = 0;
    let mut rank = 0;
    for h in hands {
        rank += 1;
        ans += rank * h.bid;
    }

    // Return answer
    ans
}

// Compare two hands
fn compare_hands(a: &Hand, b: &Hand) -> Ordering {

    // Compare by type of hand, using Joker replacement for Part 2
    let ta = hand_type(&a.cards);
    let tb = hand_type(&b.cards);
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
        if sa > sb {
            return Ordering::Greater;
        } else if sa < sb {
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
fn hand_type(cards: &String) -> u32 {

    // Count up how many of each letter, create list of frequencies in descending order
    let freq = char_freqs(&cards);
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

// Replace any Jokers in a hand, so as to maximize the value of the hand
fn replace_jokers(cards: &String) -> String {

    // Get frequency of each character
    let freqs = char_freqs(cards);
    let mut njokers = 0;
    if freqs.contains_key(&'J') {
        njokers = *freqs.get(&'J').unwrap();
    }

    // If no jokers, return string unchanged, or highest hand if all Jokers
    if njokers == 0 {
        return cards.to_string();
    } else if njokers == 5 {
        return String::from("AAAAA");
    }

    // Try replacing the jokers with each of the other cards in the hand, and
    // sort the list by score to find the best
    let mut replacements = Vec::<Hand>::new();
    for c in freqs.keys() {
        if *c != 'J' {
            let h = replace_J_chars(cards, *c);
            replacements.push(Hand{cards: h, bid: 0});
        }
    }

    // Return replacement with highest score
    replacements.sort_by(compare_hands);
    let best = replacements.last().unwrap();
    best.cards.to_string()
}

// Replace all the 'J' characters in a string with given character
fn replace_J_chars(s: &String, c: char) -> String {
    let mut cc: Vec<_> = s.chars().collect();
    for i in 0..cc.len() {
        if cc[i] == 'J' {
            cc[i] = c;
        }
    }
    cc.into_iter().collect() // Return as string
}

// Get frequency of each character in a string
fn char_freqs(s: &String) -> HashMap<char, u32> {
    let mut freq = HashMap::<char, u32>::new();
    for c in s.chars() {
        freq.entry(c).or_insert(0); // inserts if no value yet
        freq.insert(c, freq.get(&c).unwrap() + 1);
    }
    freq
}

// Strength of a card
fn strength(card: char) -> usize {
    let mut ranks = "AKQJT98765432";
    unsafe {
        if PART2 {
            ranks = "AKQT98765432J"; // part 2
        }
    }   
    match String::from(ranks).find(card) {
        Some(i) => ranks.len() - i + 1,
        None => {
            println!("*** strength on '{}'", card);
            0
        }
    }
}

// Read input file into a list of Hand structs
fn read_hands(fname: &str) -> Vec<Hand> {
    // Read input file
    let data = fs::read_to_string(fname).expect("Read error");

    // Convert data to list of hands
    let mut hands = Vec::<Hand>::new();
    for l in data.lines() {
        let parts: Vec<_> = l.split(" ").collect();
        if parts.len() == 2 {
            let c = String::from(parts[0]).trim().to_string();
            let b = parts[1].parse::<usize>().unwrap();
            hands.push(Hand { cards: c, bid: b });
        }
    }

    // Return list
    hands
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_hand_types() {
        assert!(hand_type(&String::from("AAAAA")) == 7);
        assert!(hand_type(&String::from("AA8AA")) == 6);
        assert!(hand_type(&String::from("23332")) == 5);
        assert!(hand_type(&String::from("TTT98")) == 4);
        assert!(hand_type(&String::from("23432")) == 3);
        assert!(hand_type(&String::from("A23A4")) == 2);
        assert!(hand_type(&String::from("23456")) == 1);
    }

    #[test]
    fn test_strength() {
        assert!(strength('A') == 1);
        assert!(strength('K') == 2);
        assert!(strength('2') == 13);
    }

    #[test]
    fn test_repl_jokers() {
        assert!(replace_jokers(&String::from("JJAJJ")) == "AAAAA");
    }
}
