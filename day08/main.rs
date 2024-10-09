// Advent of Code 2023, Day 08
//
// Read a list of left/right instructions, and the left/right adjacencies
// for a set of nodes. For  Part 1, calculate the number of steps required
// to get from node "AAA" to "ZZZ". For Part 2, calculate the first step
// where all routes from ??A lead to any ??Z. Part 2 solution finds and
// takes into account the repeating length of each ??A to ??Z, but uses
// brute force to find the time when they align, not very efficient.
//
// AK, Rust version 9 Oct 2024

use std::fs;
use std::collections::HashMap;
use num::integer::lcm; // Cargo.toml: num = "0.4.3"

// A left/right decision
struct Step {
    left: String,
    right: String,
}

fn main() {

    // Read data into string of instructions, and hashmap of 
    // src -> (left, right)
    let fname = "input.txt";
    let (instructions, steps) = read_data(fname);

    // Part 1: count steps from "AAA" to "ZZZ"
    println!("Part 1 (s/b 20093): {}", 
        n_steps(String::from("AAA"), String::from("ZZZ"), 
            &instructions, &steps));

    // Part 2: count iterations until all "**A" simulataneously lead to "**Z"
    // - find all nodes that end in 'A'
    // - find the number of steps from each of these nodes, to
    //   any node that ends in 'Z'
    // - get the lcm of these distances

    // Get distances from all "**A" to any "**Z"
    let mut dists = Vec::<i64>::new();
    for n in steps.keys() {    // all the start nodes
        if n.ends_with('A') {  // only interested in "**A"
            let d = n_steps(n.to_string(), String::from("any"), 
                &instructions, &steps);
            dists.push(d);
        }
    }

    // Get lowest common multiplier of all the distances
    let mut a = dists[0];
    for i in 1..dists.len() {
        a = lcm(a, dists[i]);
    }
    println!("Part 2 (s/b 22103062509257): {}", a);
}

// Find number of steps from source node to destination, traverse left/right
// instructions using the transition table. Destination "any" means any node
// that ends in 'Z'.
fn n_steps(src: String, dst: String, instructions: &String, steps: &HashMap<String, Step>) -> i64 {

    // Starting node, e.g., "AAA"
    let mut p = src.to_string();

    // Loop until reaches destination, "any" means any node that ends in 'Z'
    let mut iteration: i64 = 0; // answer will be final value of this
    let mut i = 0; // position in instructions
    let mut done = false;
    while !done  {

        // Get the next instruction
        let instr = instructions.chars().nth(i).unwrap();
        i += 1;
        if i >= instructions.len() {
            i = 0;  // circle back to beginning
        }
    
        // Go left/right from this node depending on instruction
        let step = &steps[&p];
        if instr == 'L' {
            p = step.left.to_string();
        } else {
            p = step.right.to_string();
        }

        // Determine if reached destination
        if p == dst || (dst == "any" && p.ends_with('Z')) {
            done = true;
        }

        // Next iteration
        iteration += 1;
    }

    // Return the last iteration reached
    iteration
}

// Read instructions and src -> left : right tuples into a map
fn read_data(filename: &str) -> (String, HashMap<String, Step>) {

    // Read input file into a string
    let data = fs::read_to_string(filename)
        .expect("Error reading");

    // Process each line
    let mut steps = HashMap::<String,Step>::new();
    let mut instructions = String::from("");
    for l in data.lines() {
        if instructions.len() == 0 {  // instructions on first line
            instructions = String::from(l);
        } else if l.len() > 0 { // e.g., "GSM = (PNH, BVG)"
            let s = String::from(&l[0..3]);
            let lf = String::from(&l[7..10]);
            let rr = String::from(&l[12..15]);
            steps.insert(s, Step{left: lf, right: rr});
        }
    }

    // Return string of L/R instructions, and hashmap of L/R
    (instructions, steps)
}
