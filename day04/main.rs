use std::fs;
use std::collections::HashSet;

fn main() {

    // Read file into a list of strings
    let fname = "input.txt";
    let data = fs::read_to_string(fname)
        .expect("Cannot read");
    let lines: Vec<_> = data.lines().collect();

    // Process each card
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

        // Get 2^n points, where n is the number cards you have that are winning
        let mut pts: u32 = 0;
        let isect = have.intersection(&winning).count() as u32;
        if isect > 0 {
            pts = u32::pow(2, isect - 1);
        }
        
        part1 += pts;
    }
    println!("Part 1 = {}", part1);
       
}

// Convert list to set
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

