// Advent of Code 2023, Day 02
//
// Given a list of games, each with a list of red/green/blue cubes
// drawn from a bag, determine which games are possible with a given
// number of red, green, and blue cubes. For Part 2, determine the
// minimum number of cubes of each colour required to play all the
// turns in a game, and multiply these together.
//
// AK, 3 Oct 2024 for Rust version

use std::fs;
use std::collections::HashMap;

fn main() {

    // Number of balls available of each colour, for Part 1
    let mut available = HashMap::new();
    available.insert("red", 12);
    available.insert("green", 13);
    available.insert("blue", 14);
    
    // Read input file
    let fname = "input.txt";
    let data = fs::read_to_string(fname)
        .expect("Cannot read");

    // Process each line
    let mut id = 0;  // ID of the next line
    let mut part1 = 0; // total for Part 1
    let mut part2 = 0; // total for Part 2
    for l in data.lines() { // each line

        // Parse the line into a list of color->count hashmaps
        let game = parse_line(l);

        // For Part 2, we need the minimum number of red/green/blue 
        // balls required to play each game
        let mut min_balls = HashMap::new();
        min_balls.insert("red", 0);
        min_balls.insert("green", 0);
        min_balls.insert("blue", 0);

        // Check whether this game is feasible given available number 
        // of balls (part 1), keep track of the number of balls that 
        // would make this game feasible (part 2)
        let mut okay = true; // for part 1
        for draw in game {
            for (c, n) in draw.iter() {

                // Part 1: mark this game as invalid if insufficient balls 
                // of this colour
                if *n > available[c] {
                    okay = false;
                } 

                // Part 2: keep track of the number of balls required
                if *n > min_balls[c] {
                    min_balls.insert(c, *n);
                }
            }
        }

        // For Part 1, if game is valid, add its ID to the answer
        id += 1;
        if okay {
            part1 += id;
        }
        
        // Part 2: calculate "power" and add it to total
        let power = min_balls["red"] * min_balls["green"] * min_balls["blue"];
        part2 += power;
    }

    // Show the answers
    println!("Part 1 (s/b 2278) = {}, Part 2 (s/b 67953)= {}", part1, part2);
}

// Parse a line representing a game, into a list of maps,
// each color -> number
// "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
fn parse_line(l: &str) -> Vec::<HashMap<&str, i32>> {

    // Create a list of maps
    let mut game: Vec<_> = Vec::<HashMap<&str, i32>>::new();

    // Remove game number, just beyond colon
    let colon = l.find(':').unwrap()+2;
    let l1 = &l[colon..]; // skip past the colon

    // Turn each draw of balls (separated by semicolons) into a hashmap
    for d in l1.split(';') { // each draw of multiple balls
        let mut balls= HashMap::new();
        for b in d.trim().split(',') { // e.g., "3 blue"
            let parts: Vec<_> = b.trim().split(' ').collect();
            let n = parts[0].trim().parse::<i32>().unwrap(); // count
            let c = parts[1].trim(); // color
            balls.insert(c, n); // add to map
        }
        game.push(balls);  // add hashmap to list
    }

    // Return the list of hashmaps
    game
}
