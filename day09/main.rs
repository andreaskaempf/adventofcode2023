// Advent of Code 2023, Day 09
//
// Given a list of numbers, repeatedly get the deltas between each successive
// number. Then, use the delta to extrapolate back upwards to the previous
// row. For Part 1, extrapolate at the end of each row, for Part 2, before
// the beginning. Answers are the sums of the extrapolations.
//
// AK, Rust solution 10 Oct 2024

use std::fs;

fn main() {

    // Read data into list of numbers
    let data = read_data("input.txt");

    // Process each row and extrapolate the next number on the first row,
	// at end for Part 1, before beginning for Part 2. Add up extrapolations
	// to get answers for parts 1 and 2.
	let mut ans1 = 0;
    let mut ans2  = 0;
	for row in &data {

		// Get successive lists of deltas for this row
		let mut deltas = Vec::<Vec<i32>>::new();
		let mut dd = get_deltas(&row);
		while dd.len() > 0 {
			deltas.push(dd.clone()); // TODO: avoid need to clone?
			dd = get_deltas(&dd);
		}

		// Now extrapolate each row from the bottom up
		let mut x1 = 0;
        let mut x2 = 0;
		let mut i = deltas.len() - 1;
        loop {
			let dd = &deltas[i];
			x1 += dd[dd.len()-1]; // part 1
			x2 = dd[0] - x2;      // part 2
            if i == 0 {           // usize can't go negative, so break if negative
                break
            }
            i -= 1; 
		}

		// Extrapolate first row, add to answers
		ans1 += row[row.len()-1] + x1;
		ans2 += row[0] - x2;

	}

    // Show both answers
	println!("Part 1 (s/b 1782868781)= {}", ans1);
	println!("Part 2 (s/b 1057)= {}", ans2);
}

// Get list of deltas for a list of numbers
fn get_deltas(nums: &Vec<i32>) -> Vec<i32> {
	let mut deltas = vec![0; nums.len()-1];
	for i in 1..nums.len() {
		deltas[i-1] = nums[i] - nums[i-1];
	}
	deltas
}

// Read a file, each line a list of numbers separated by spaces
fn read_data(fname: &str) -> Vec<Vec<i32>> {
    let data = fs::read_to_string(fname).expect("Error reading");
    let mut res = Vec::<Vec<i32>>::new();
    for l in data.lines() {
        let nums: Vec<_> = l.split(' ').map(parse_int).collect();
        res.push(nums);
    }
    res
}

// Parse an integer, used for map over list
fn parse_int(s: &str) -> i32 {
    s.parse::<i32>().unwrap()
}
