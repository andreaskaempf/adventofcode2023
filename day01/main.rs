// Advent of Code 2023, Day 01
//
// Find digits in a string, combine first and last digit on each line to
// make a number, and add these up. For part 2, look for embedded names
// of numbers as well, complicated as last & first letters are shared
// if they are the same.
//
// To compile and run without having to use Cargo:
//   rustc main.rs
//   ./main
//
// AK, Rust solution 1 Oct 2024

use std::fs;

fn main() {

    // Read whole file into memory (uncomment relevant file name)
    //let fname = "sample.txt";
    let fname = "input.txt";
    let data = fs::read_to_string(fname)
        .expect("Error reading file");

    // Process each line, add up totals for Parts 1 and 2
    let mut part1: u32 = 0;
    let mut part2: u32 = 0;
    for l in data.lines() {
        part1 += check(l, false);
        part2 += check(l, true);
    }
    println!("Part 1 = {}, Part 2 = {}", part1, part2);
}

// Create number from first and last digits found in the string,
// converting digit names for part 2
fn check(s: &str, convert_names: bool) -> u32 {

    // Convert to a string so we can use replace()
    let mut s1 = String::from(s);

    // For Part 2, convert number names to digits
    if convert_names {

        // Tricky double combinatations
        s1 = s1.replace("oneight", "18");
        s1 = s1.replace("twone", "21");
        s1 = s1.replace("threeight", "38");
        s1 = s1.replace("fiveight", "58");
        s1 = s1.replace("sevenine", "79");
        s1 = s1.replace("eightwo", "82");
        s1 = s1.replace("nineight", "98");

        // Single digits
        s1 = s1.replace("one", "1");
        s1 = s1.replace("two", "2");
        s1 = s1.replace("three", "3");
        s1 = s1.replace("four", "4");
        s1 = s1.replace("five", "5");
        s1 = s1.replace("six", "6");
        s1 = s1.replace("seven", "7");
        s1 = s1.replace("eight", "8");
        s1 = s1.replace("nine", "9");
        s1 = s1.replace("zero", "0");
    }

    // Find first and last digits in the line
    let mut first_dig = 999;
    let mut last_dig = 0;
    for c in  s1.chars() {
        if c >= '0' && c <= '9' {
            let x = c.to_digit(10).unwrap(); // numeric value
            if first_dig == 999 {
                first_dig = x;
            }
            last_dig = x;
        }
    }

    // Return number
    first_dig * 10 + last_dig
}
