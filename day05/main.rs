// Advent of Code 2023, Day 05
//
// Given a list of "seeds" (numbers), go through a series of transformations
// based on tables. Each table is a tuple (dst, src, n). If the number is
// within src .. src+n,  it is assigned to dst. If no dst is found, the dst
// is the same as src for that transformation. For Part 1, find the lowest
// final transformation. Part 2 is the same, but treat each pair of "seeds"
// as start and length of a range. Brute force is okay in Go, but
// would be faster with concurrency, memoization, or a smarter approach.
//
// AK, 5 Oct 2024 for Rust version

use std::fs;

// Section of the map, with name and list of number lists
struct Section {
    name: String,
    maps: Vec<Vec<u64>>,
}

fn main() {

    // Read input file into lists of seed numbers and sections
    //let fname = "sample.txt"; // "input.txt";
    let fname = "input.txt";
    let (seeds, secs) = read_file(fname);

    // Part 1: trace each seed through transformations, return the
    // lowest ending location found
    let mut lowest = 999;
    for s in &seeds {
        let l = process_seed(s, &secs);
        if lowest == 999 || l < lowest {
            lowest = l;
        }
    }
    println!("Part 1 (s/b 424490994): {}", lowest);

    // Part 2: same, but treat seed numbers as ranges
	lowest = 999;
	let mut si = 0; 
    while si < seeds.len() {
		let s1 = seeds[si];  // start of range
		let n = seeds[si+1]; // length of range
		println!("Seed {} for length {}", s1, n);
		for i in 0..n {
			let loc = process_seed(s1 + i, &secs);
			if lowest == 999 || loc < lowest {
				lowest = loc;
			}
		}
        si += 2;
	}
    println!("Part 2 (s/b 15290096): {}", lowest);

}

// Process a seed, by traversing the transformations and getting
// the end location
fn process_seed(seed: u64, sections: &Vec<Section>) -> u64 {

    // Go through each section
    let mut src = seed;
    for sec in sections {

        // Find the destination corresponding to the current source
        let mut dst  = 0;
        let mut found = false;
        for ranges in &sec.maps {
            if src >= ranges[1] && src <= ranges[1]+ranges[2] {
                dst = ranges[0] + src - ranges[1];
                found = true;
                break
            }
        }

        // If not found, use the source value
        if !found  {
            dst = src;
        }

        // Source of next transformation is the destination of this one
        src = dst;
    }

    // Return the final destination
    src
}

// Read and parse input file, return list of seed numbers, and
// list of Section structures
fn read_file(fname: &str) -> (Vec<u64>, Vec<Section>) {

    let mut seeds: Vec<_> = Vec::<u64>::new();
    let mut sections: Vec<_> = Vec::<Section>::new();
    let mut in_section = false;

    // Read data, process each line
    let data = fs::read_to_string(fname).expect("Unable to read");
    for l in data.lines() {

        // First line has list of seeds
        if seeds.len() == 0 {
            let colon = find_char(':', l).unwrap();
            seeds = parse_ints(&l[colon+1..]);
        }
        // Blank line is end of map section
        else if l.len() == 0 {
            in_section = false;
        }
        // Line following blank is name of new section
        else if !in_section  {
            let space = find_char(' ', l).unwrap();
            let sec_name = &l[..space]; // name of section, up to space
            let new_section =  Section{name: sec_name.to_string(), maps: Vec::<Vec<u64>>::new()};
            sections.push(new_section);
            in_section = true;
        } 
        // Truples of numbers, add list to most recent section
        else if sections.len() > 0 {
            let truple = parse_ints(l);
            let last = sections.len()-1;
            sections[last].maps.push(truple);
        }
    }

    // Return list of seeds and list of sections
    (seeds, sections)
}

// Find character in a string, return index
fn find_char(c: char, s: &str) -> Option<usize> {
    for i in 0..s.len() {
        if s.chars().nth(i) == Some(c) {
            return Some(i);
        }
    }
    None
}

// Parse a space-separated list of integers
fn parse_ints(s: &str) -> Vec<u64> {
    s.split(" ")
        .map(|n| n.trim())
        .filter(|s| s.len() > 0)
        .map(|s| s.parse::<u64>().unwrap())
        .collect()
}
