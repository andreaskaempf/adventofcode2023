// Advent of Code 2023, Day 06
//
// Simulate a race, in which you hold down a button to recharge a motor,
// which lets the vehicle go faster. Distance is the speed muliplied by
// remaining time. Find the number of ways you can hold down the button, so
// as to exceed the previous maximum distance given. For part 2, just one
// race, but much bigger number.
//
// AK, Rust version 5 Oct 2024

fn main() {

    // Sample and input data for Part 1 (hard coded, no input files, uncomment as necessary
    //let time = Vec::from([7, 15, 30]);  // sample
    //let distance = Vec::from([9, 40, 200]); //sample
    let time = Vec::from([60, 80, 86, 76]); // input
    let distance = Vec::from([601, 1163, 1559, 1300]); // input

    // Part 1: For each race, count up the number of ways you can hold down
    // the button, so as to exceed the given maximum distance, and multiply
    // these counts together
    let mut part1 = 1; // answer is product, so start with 1
    for i in 0..time.len() { // each race
        let T = time[i];
        let D = distance[i];
        let mut better = 0;
        for t in 1..T {
            let  speed = t;
            let  dist= (T - t) * speed;
            if dist > D {
                better += 1;
            }
        }
        part1 *= better;
    }
    println!("Part 1 (s/b 1155175): {}", part1);

    // Part 2: same idea, but one very long race, time and disance obtained
	// by concatenating the times and distances of the sample values
	//let T = 71530; // sample values
	//let D = 940200;
	let T: i64 = 60808676; // input values
	let D: i64 = 601116315591300;
	let mut part2: i64 = 0;
	for t in 1..T {
		let speed = t;
		let dist = (T - t) * speed;
		if dist > D {
			part2 += 1;
		}
	}
	println!("Part 2 (s/b 35961505): {}", part2)
}
