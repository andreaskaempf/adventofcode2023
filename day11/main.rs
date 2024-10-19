// Advent of Code 2023, Day 11
//
// Given a 2D field of 'galaxies' surrounded by space, first 'expand' this by
// doubling each row/column that is empty, and add up the manhattan distance
// between every pair. For Part 2, expand by factor of a million instead of
// doubling. Go's default int size is 64, enough to handle this.
//
// AK, 19 Oct 2024 for Rust solution

use std::fs;

struct Point {
    x: usize,
    y: usize
}

fn main() {

    // Read input file, convert to pseudomatrix
    let mat = read_char_matrix("input.txt");
    
    // Get the coordinatees of all the '#' characters into a list
    let mut points = Vec::<Point>::new();
    for ri in 0..mat.len() {
        for ci in 0..mat[ri].len() {
            if mat[ri][ci] == '#' {
                points.push(Point{x: ci, y: ri});
            }
        }
    }

    // Increasing the size of empty rows and columns is achieved by adding
    // 1 to subsequent coordinates to double (for Part 1), adding 1000000-1
    // to increase by a million (for Part 2)
    //let diff = 1;  // uncomment for Part 1
    let diff = 1000000-1; // uncomment for Part 2

    // Starting at the right, double the width of every "emtpy" column, by
    // increasing the x coordinates of anything to the right of that column
    let mut x = mat[0].len() - 1; // maximum x
    loop  {
        // If column is empty, increase coords of points to right
        if empty_column(x, &points) {
            for p in &mut points {
                if p.x > x {
                    p.x += diff;
                }
            }
        }

        // Move to previous column, or done if at the first column
        if x == 0 {
            break;
        }
        x -= 1;
    }

    // Repeat for y coordinates, i.e., starting at the bottom, double the 
    // width of every "emtpy" row, by increasing the y coordinates of 
    // anything below that row
    let mut y = mat.len() - 1; // maximum y
    loop  {
        // If column is empty, increase coords of points to right
        if empty_row(y, &points) {
            for p in &mut points {
                if p.y > y {
                    p.y += diff;
                }
            }
        }
        
        // Move to previous row, or done if at the top
        if y == 0 {
            break;
        }
        y -= 1;
    }

    // For answer, add up the manhattan distances between all pairs of hashes
    let mut ans: usize = 0;
	for i in 0..points.len() {
		for j in 0..points.len() {
			if j > i {
				let p1 = &points[i];
				let p2 = &points[j];
				ans += dist(p1.x, p2.x) + dist(p1.y, p2.y);
			}
		}
	}
	println!("Answer (s/b 9521776 for Part1, 553224415344 Part 2): {}", ans);
}

// Distance between two points
fn dist(a: usize, b: usize) -> usize {
    if a > b {
        return a - b;
    }
    b - a
}

// Is there anything in this column?
fn empty_column(x: usize, points: &Vec<Point>) -> bool {
    for p in points {
        if p.x == x {
            return false;
        }
    }
    true
}

// Is there anything in this row?
fn empty_row(y: usize, points: &Vec<Point>) -> bool {
    for p in points {
        if p.y == y {
            return false;
        }
    }
    true
}

// Read a text file as a matrix of chars, i.e., a list of rows, each a 
// list of chars
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
