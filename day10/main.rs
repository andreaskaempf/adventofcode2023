// Advent of Code 2023, Day 10
//
// Given a 2D terrain of symbols, start at 'S' and follow the chain of "pipe"
// shapes back to S. For Part 1, calculate the distance of the furthest-away
// shape, distance measured in either direction from start. For Part 2,
// count up how many positions are *not* inside the shape.
//
// AK, Rust solution 12 Oct 2024, more efficient algorithm than the Go
// version, which assumed there could be multiple entry points from any
// location, and used recursion to explore all possible paths.

use std::fs;
use std::collections::HashMap;

// A point on the map, just x,y coordinates. The "derive"
// stuff at the top gets the compiler to create traits to
// copy, compare, or hash a Point object.
#[derive(Copy, Clone, Eq, PartialEq, Hash)]
struct Point {
    x: usize,
    y: usize
}

fn main() {

    // Read input file into a list of rows, each a list of chars
    let fname = "input.txt";
    let map = read_char_matrix(fname);

	// Part 1: starting at 'S', explore the pipes until you get back to 'S',
	// recording the distance along the way.

	// Find the starting point, where 'S' is
    let start = find_char_in_matrix('S', &map);
    println!("Starting at row {}, col {}", start.y, start.x);

	// Dictionary of distances from start, also indicates whether a point has been visited
	let mut dists = HashMap::<Point,i32>::new();
	dists.insert(start, 0); // mark start as visited

	// Explore from here, until you get back to 'S'
	let mut dist = 0;  // current distance from start
	let mut p = start; // the current point
	loop {

        // Find the possible locations you could go from here (should only be one or two, as there
        // are no T-junctions)
		let opts = next_steps(&p, &map);
		if opts.len() > 2 {  // should never happen
			println!("More than 2 options!");
			break;
		}

        // Move to the unvisited option from this point; if no more unvisited locations, we're done
        // (don't need to explictely check if we're back at start)
		p = opts[0];
		if dists.contains_key(&p) && opts.len() > 1 { // i.e., already visited
			p = opts[1]; // take the other unvisited one
		}
		if dists.contains_key(&p) {
			println!("No more unvisited points to explore");
			break;
		}

		// Set the distance of the new point
		dist += 1;
		dists.insert(p, dist);
	}

	// Part 1 is the longest distance from start, equal to half of total distance traversed
	println!("Part 1 (s/b 6927) = {}", (dist + 1) / 2);

    // Part 2: Use ray tracing method on rows only, to count up points that are enclosed by path.
    // Involves traversing each row, and counting up how many times we cross the outline of our
    // path. Where the count is odd, we are inside the shape, otherwise outside.
	// Source: https://www.quora.com/What-are-the-algorithms-for-determining-if-a-point-is-inside-an-arbitrary-closed-shape-or-not
	let mut ans2 = 0;
	for y in 0..map.len()  {        // each row
		let mut bars = 0;           // number of borders encountered on this row
		for x in 0..map.len()  {    // each character on row
			let p = Point{x: x, y: y};
			if dists.contains_key(&p) && one_of(at(&p, &map), "S|LJ") {
				bars += 1;          // we just crossed another border
			} else if !dists.contains_key(&p) && bars % 2 == 1 {
				ans2 += 1;          // odd number of borders crossed, inside shape
			}
		}
	}

	println!("Part 2 (s/b 467) = {}", ans2);
}

// Return a list of the viable next steps from this position, depends on the
// current and next shape being compatible connectors.
// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
fn next_steps(p: &Point, map: &Vec<Vec<char>>) -> Vec<Point> {

	// For list of possible moves
	let mut res = Vec::<Point>::new();
  
 	// character at this position
	let c0 = at(p, map);     

	// Up: bar, F, 7
	if let Some(p1) = make_point(&p, 0, -1, map) {
		let c = at(&p1, map);
		if one_of(c0, "S|LJ") && one_of(c, "|F7") {
			res.push(p1);
		}
	}

	// Down: vert bar, L, J
	if let Some(p1) = make_point(&p, 0, 1, map) {
		let c = at(&p1, map);
		if one_of(c0, "S|F7") && one_of(c, "|LJ") {
			res.push(p1);
		}
	}	

	// Left: dash, F, L
	if let Some(p1) = make_point(&p, -1, 0, map) {
		let c = at(&p1, map);
		if one_of(c0, "S-J7") && one_of(c, "-FL") {
			res.push(p1);
		}
	}

	// Right: dash, 7, J
	if let Some(p1) = make_point(&p, 1, 0, map) {
		let c = at(&p1, map);
		if one_of(c0, "S-FL") && one_of(c, "-7J") {
			res.push(p1);
		}
	}

	// Return list of points
	res
}

// Make a new point based on original point, x and y adjusted,
// return None if out of bounds. 
// Warning: this may fail if dx or dy are < -1
// This is absolutely horrible, so much easier in Go or Python, where
// it's just one line of code, e.g., p = Point{x+2, y+1}
fn make_point(p: &Point, dx: i32, dy: i32, map: &Vec<Vec<char>>) -> Option<Point> {

    // Convert everything to signed integers, because usize won't allow adding 
    // integer, what an abomination is Rust
	let px: i32 = p.x.try_into().unwrap();
	let py: i32 = p.y.try_into().unwrap();
	let nrows: i32 = map.len().try_into().unwrap();
	let ncols: i32 = map[0].len().try_into().unwrap();

	// New coordinates, check for out of bounds, return None if so
	let x1 = px + dx;
	let y1 = py + dy;
	if y1 < 0 || y1 >= nrows || x1 < 0 || x1 >= ncols { // y out of bounds
		return None;
	} 

	// Create new point, need to convert back to unsigned
	let new_x: usize = (px + dx).try_into().unwrap();
	let new_y: usize = (py + dy).try_into().unwrap();
	Some(Point{x: new_x, y: new_y})
}

// Character at a point, 0 if out of bounds
// Note that we don't need to check for < 0 since coordinates are usize
fn at(p: &Point, map: &Vec<Vec<char>>) -> char {
	if  p.y >= map.len() || p.x >= map[0].len() {
		return '\0';
	} else {
		return map[p.y][p.x];
	}
}

// Check if a char is in string, could be done inline, but matches Go version
fn one_of(c: char, s: &str) -> bool {
	s.contains(c)
}

// Find a character in a "matrix", returns (0,0) if not found (but this 
// may be the valid answer, so check). Only used to find the initial 'S'.
fn find_char_in_matrix(ch: char, map: &Vec<Vec<char>>) -> Point {
    let mut r = 0;
    let mut c = 0;
    for ri in 0..map.len() {
        for ci in 0..map[ri].len() {
            if map[ri][ci] == ch {
                r = ri;
                c = ci;
                break;
            }
        }
        if map[r][c] == 'S' {
            break;
        }
    }
    Point{x: c, y: r}
}

// Read a text file as a matrix of chars, i.e., a list of rows, each a list of chars
fn read_char_matrix(fname: &str) -> Vec<Vec<char>> {
    let data = fs::read_to_string(fname).expect("Read error");
    let mut map = Vec::<Vec<char>>::new();
    for l in data.lines() {
        if l.len() > 0 {
            map.push(l.chars().collect());
        }
    }
    map
}
