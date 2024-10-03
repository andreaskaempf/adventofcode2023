// Advent of Code 2023, Day 03
//
// Given a surface of digits and special characters, find any numbers that
// are adjacent to special characters. For Part 2, find any pairs of numbers
// that are adjacent to the same asterisk.
//
// AK, Rust version 3 Oct 2024

use std::fs;
use regex::Regex; // Cargo.toml: regex = "1.3.9"

struct Symbol {
    y: usize,
    x: usize,
    symbol: char,
}

struct SymbolNums {
    s: Symbol,
    nums: Vec<i32>,
}

fn main() {
    
    // Read input file into a list of strings
    //let fname = "sample.txt";
    let fname = "input.txt";
    let data = fs::read_to_string(fname)
        .expect("Unable to read");
    let lines: Vec<_> = data.lines().collect();

    // Regular expression for a sequence of digits
    let re = Regex::new(r"\d+").unwrap();

    // Process each line
    let mut part1 = 0;
    let mut all_symbols = Vec::<SymbolNums>::new();
    for i in  0..lines.len() {

        // Find all sequences of digits in this line
        let l = lines[i];
        for mat in re.find_iter(l) {

            // Get position and value of match
            let p0 = mat.start();
            let p1 = mat.end();
            let num =  &l[p0..p1];
            let n: i32 = num.parse::<i32>().unwrap();

            // Find the symbols adjacent to this number
            let adj_symbols = adj_symbols(i, p0, p1, &lines); 

            // Part 1: add up numbers that are adjacent to any symbol
            if adj_symbols.len() > 0 {
                part1 += n;
            }

            // For part 2, accumulate list of part numbers adjacent to each symbol found
            for s in adj_symbols {

                // Find the symbol in list, add number to list if found
                match find_symbol(&s, &all_symbols) {
                    Some(i) => all_symbols[i].nums.push(n),
                    None => {
                        let mut s1 = SymbolNums{s: s, nums: Vec::<i32>::new()};
                        s1.nums.push(n);
                        all_symbols.push(s1);
                    }
                }
            }
        }
       
    }

    // Answer for Part 1
    println!("Part 1 (s/b 4361, 560670) = {}", part1);

     // Part 2: add up products of pairs of part numbers that are adjacent to the 
     // same asterisk
     let mut part2 = 0;
     for s in all_symbols {
        if s.s.symbol == '*' && s.nums.len() == 2 {
            part2 += s.nums[0] * s.nums[1];
        }
     }
     println!("Part 2 (s/b 467835, 91622824) = {}", part2);
}


// Find symbol in a list, return index to the symbol found
// TODO: tried to return pointer but did not succeed
fn find_symbol(s: &Symbol, ss: &Vec<SymbolNums> ) -> Option<usize> {
    for i in 0..ss.len() {
        let s1 = &ss[i].s;
        if s1.x == s.x && s1.y == s.y && s1.symbol == s.symbol {
            return Some(i);
        }
    }
    None
}

// Check if there are any symbols adjacent to sequence of characters on given line,
// return a list of the symbols found with their coordinates
fn adj_symbols(line_no: usize, start: usize, end: usize, lines: &Vec<&str>) -> Vec<Symbol> {

    // We will return a list of all the symbols this number is adjacent to,
    // with their x,y position
    let mut symbols = Vec::new();

    // Row before: check all positions from before number to one after
    if line_no > 0 {
        let l = lines[line_no - 1];
        let mut x = start;
        if x > 0 {
            x -= 1;
        }
        while x < l.len() && x <= end {
            let c = l.chars().nth(x).unwrap(); // WTF???!!!
            if issymbol(c) {
                let symb = Symbol{y: line_no-1, x: x, symbol: c};
                symbols.push(symb);
            }
            x += 1;
        }
    }

    // Row after: check all positions from before number to one after
    if line_no < lines.len() - 1 {
        let l = lines[line_no + 1];
        let mut x = start;
        if x > 0 {
            x -= 1;
        }
        while x < l.len() && x <= end  {
            let c = l.chars().nth(x).unwrap(); // WTF???!!!
            if issymbol(c) {
                let symb = Symbol{y: line_no+1, x: x, symbol: c};
                symbols.push(symb);
            }
            x += 1;
        }
    }

    // This row: check character before and after
    let l = lines[line_no];
    if start > 0 {
        let c = l.chars().nth(start-1).unwrap();
        if issymbol(c) {
            let symb = Symbol{y: line_no, x: start-1, symbol: c};
            symbols.push(symb);
        }
    }
    if end < l.len() {
        let c = l.chars().nth(end).unwrap();
        if issymbol(c) {
            let symb = Symbol{y: line_no, x: end, symbol: c};
            symbols.push(symb);
        }
    }

    // Return list of symbols adjacent to this number
    symbols
}

// Check if character is a symbol, e.g., not a period or digit
fn issymbol(c: char) -> bool {
    c != '.' && !isdigit(c)
}

// Check if character is a digit
fn isdigit(c: char) -> bool {
    c >= '0' && c <= '9'
}
