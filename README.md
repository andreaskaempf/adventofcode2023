# Advent of Code 2023

My solutions for the Advent of Code 2023, 
see https://adventofcode.com/2023

* **Day 1** (Go, Rust): Find digits in a string, combine first and last digit
  on each line to make a number, and add these up. For part 2, look for
  embedded names of numbers as well, complicated as last & first letters are
  shared if they are the same.

* **Day 2** (Go, Rust): Given a list of games, each with a list of 
  red/green/blue cubes drawn from a bag, determine which games are possible 
  with a given number of red, green, and blue cubes. For Part 2, determine 
  the minimum number of cubes of each colour required to play all the
  turns in a game, and multiply these together.

* **Day 3** (Go, Rust): Given a 2-d surface of digits and special characters, 
  find any embedded numbers that are adjacent to special characters, and
  add them up. For Part 2, find any pairs of numbers that are adjacent to 
  the same asterisk, and sum up the products of these pairs.

* **Day 4** (Go, Rust): You are given a list of "cards", each with two lists of 
  numbers separated by a vertical bar. Some of the numbers in the right 
  list are also in the left list. For Part 1, calculate "points" for each 
  card by adding up 1, 2, 4, ... for each match. For Part 2, make a copy 
  of the next n cards, where n is the number of matches, and count up the 
  cards at the end.

* **Day 5** (Go, Rust, Zig, C): Given a list of "seeds" (numbers), go through a
  series of transformations for each seed, based on tables. Each table is a
  tuple (dst, src, n). If the number is within src ..  src+n,  it is assigned
  to dst. If no dst is found, the dst is the same as src for that
  transformation. For Part 1, find the lowest final transformation. 
  Part 2 is the same, but treat each pair of "seeds" as start and length 
  of a range. Computationally intensive brute force approach takes about 5:40
  minutes to run in **Go**, and the same algorithm about 3:45 in **Rust** in 
  release mode (197 mins in debug mode!), suggesting that Go takes about 50% 
  more time than Rust. **Zig** implementation compiled with ReleaseFast 
  option takes 1:51, half of the Rust version. And Zig is even faster than 
  **C**, which takes 2:40 using gcc with the -O3 optimization option. This is
  just an indicative benchmark using a naive implementation -- in any 
  language, a better algorithm would be faster, as would the addition of 
  concurrency and/or memoization.

* **Day 6** (Go, Rust): Simulate a toy boat race, in which you hold down a button
  to recharge a motor, which lets the boat go faster. Distance is the speed
  muliplied by remaining time. Find the number of ways you can hold down the
  button, so as to exceed the previous maximum distance given. For part 2, just
  one race, but much bigger number (still almost instantaneous using brute
  force).

* **Day 7** (Go, Rust): Given some hands of 5 cards, sort them by poker rank. 
  For part 2, replace instances of J (joker) with which ever other letter
  from the hand, that yields the best poker hand type. Answer for both 
  parts is the sumproduct of the rank (sorted sequence) times a given bid.

* **Day 8** (Go, Rust): Read a list of left/right instructions, and the left/right
  adjacencies for a set of nodes. For  Part 1, calculate the number of steps
  required to get from node "AAA" to "ZZZ". For Part 2, calculate the first
  step where all routes from ??A lead to any ??Z. Part 2 solution finds and
  takes into account the repeating length of each ??A to ??Z, but uses brute
  force to find the time when they align, not very efficient.

* **Day 9** (Go, Rust): Given a list of numbers, repeatedly get the deltas between
  each successive number. Then, use the delta to extrapolate back upwards to
  the previous row. For Part 1, extrapolate at the end of each row, for Part 2,
  before the beginning. Answers are the sums of the extrapolations.

* **Day 10** (Go, Rust): Given a 2D terrain of symbols, start at 'S' and follow the
  chain of "pipe" shapes back to S. For Part 1, calculate the distance of the
  furthest-away shape, distance measured in either direction from start.  For
  Part 2, count up how many positions are *not* inside the shape (solved using
  ray tracing method).

* **Day 11** (Go, Rust): Given a 2D field of 'galaxies' surrounded by space, 
  first 'expand' this by doubling each row/column that is empty, and add up 
  the manhattan distances between all the pairs. For Part 2, expand by factor 
  of a million instead of doubling. 

* **Day 12** (Go): Given a pattern of '#', '.', and '?' symbols, and a list of
  numbers representing the lengths of blocks of '#', find the number of
  possible replacements of '?' with either '.' or '#', so that the lengths of
  the '#' sequences match the given list of numbers. For Part 2, the pattern is
  repeated five times, with '?' in between, and the list of numbers is also
  repeated five times, defeating a brute force solution.

* **Day 13** (Go): Given a set of 2D fields consisting of '.' and '#'
  characters, find the row or column in each that gives a mirror reflection,
  i.e., left/right sides are mirror images, or top/bottom sides. For Part 2,
  flip every character on each field, to find a different reflection
  (ignoring the previous one).

* **Day 14** (Go): Given a square field of 'O' (round rocks), '#' (square rocks)
  and '.', move all the round rocks as much as possible to the top. For Part 2,
  do this four times (once in each direction) to make a cycle, and do 10e9
  cycles. For both, answer is the score calculated as the sum of the number of
  round rocks below in each column, for each row. Part 2 cannot be calculated
  using brute force, but the pattern repeats after a while, so use this to
  quickly determine what the pattern would look like after many iterations.

* **Day 15** (Go): For Part 1, sum up the hashes of a list of commands, using a
  custom has function. For Part 2, simulate the movement of "lenses" among a
  chain of 256 boxes, by executing the commands.

* **Day 16** (Go): Given square a field of characters, simulate the movement of
  a "beam of light" entering the field. Slashes and backslashes change the
  direction of travel 90 degrees, and dashes or vertical bars "split" the beam
  in two. "Energy" is defined as the number of cells in the field that are
  eventually  touched by the beam. For Part 1, simulate a single beam entering
  top left, and report the energy. For Part 2, try entering from every position
  on every edge, and report the highest energy found.  Quite easy, but need to
  know when to stop, ignore beams that leave the field and avoid creating
  duplicates of beams that are identical to any that already exist.

* **Day 17** (Go): Given a matrix of digits, find the sum of the digits along
  the shortest path from the top left to the bottom right, such that you
  never go in the same direction more than 3 steps. For Part 2, maximum 10
  steps in the same direction, minimum 4 steps before changing direction.
  Solved by modifying the Djistra algorithm, transcribed from my Julia
  solution to AoC 2021, Day 15, to enforce constraints on direction and
  number of steps in the same direction.

* **Day 18** (Go): Given a set of instructions to draw a polygon consisting of
  just horizontal and vertical lines, count the number of points that are
  inside the polygon. For Part 2, instructions are revised to give a
  much bigger shape, that cannot be computed in memory. Did Part 1 using
  simple recursive flood fill, Part 2 using go-geom library.

* **Day 19** (Go): Given a set of rules, and a set of parts, determine which
  parts are accepted by the rules. Rules are a series of tests, each of which
  is either a comparison (e.g., "s<1351:dest") or a destination.  For Part 2,
  determine how many parts, within a universe of 0..4000 for each of the four
  parameters, would be accepted. Did this by recursively evaluating the rules,
  to come up with a list of "and" conditions for each "A" terminal node. For
  each "A" rule, the number of parts accepted is the product of the widths of
  the four resulting ranges.

* **Day 20** (Go): Given a list of components and their outputs, simulate the 
  execution of signal transmission between the components when a "button" is
  pressed.  For Part 1, determine the total number of low & high signals sent
  over 1000 button presses (done by simple simulation). For Part 2, need to
  determine how many button presses until one "rx" module receives a signal.
  This was done by looking for cycles in the four inputs to the only input to
  rx, and multiplying these cycle lengths together to give the times when the
  cycles coincide.

* **Day 21** (Go): Given a map of points and rocks, find the number of points
  that can be reached in n steps, starting from a given point.  Used Djikstra's
  algorithm (and later simple walk simulation) to find the number of points
  that can be reached in n steps.  For Part 2, assume a much larger number of
  steps, infeasible using brute force, so count tiles reached in 1x1, 3x3, and
  5x5 area blocks, and extrapolate from these.

* **Day 22** (Go): Given a set of long bricks, each defined by two points 
  in 3d space, determine which bricks can be removed without causing any other
  bricks to fall. For Part 2, determine how many bricks can be removed without
  causing any other bricks to fall. Fairly straightforward simulation, not very
  efficient.

* **Day 23** (Go): Find the maximum number of steps (longest path) from top to
  bottom of a grid, ignorning blocks (hash marks), and not revisiting previous
  cells. In Part 1, pointer chars indicate that you must move in that direction
  (restriction removed in Part 2). Hard problem, used recursive depth-first 
  search (brute force), used brute force, but could probably simplify the graph 
  to reduce the search space.

* **Day 24** (Go, Python): Given a list of stones, each with a position and velocity,
  find the number stones whose trajectories will intersect. For Part 2, find
  the position and velocity of a new stone, whose trajectory will intersect
  with that of every other stone. Part 2 done using constraint solvers 
  (Centipede in Go and Z3 from Python).

* **Day 25** (Go): Given a graph of nodes ("components"), find the three edges
  ("wires") you need to disconnect in order to divide the graph into two
  separate groups, multiply the sizes of these two groups together. Used
  yourbasic/graph to create graph, and calculated the total distance from node
  0 to all other nodes, tried all combinations of the 100 nodes with the
  highest graph distance when removed, to check if graph has 2 components (i.e.,
  subgraphs not connected together). Picking the edges that resulted in
  the highest total graph distance was necessary because it's too much
  computation to consider all 3-edge combinations amount the ~3500 edges.
  This was Part 1, there was no Part 2 (except for being able to push a
  "button" provided that all previous gold stars were earned).

To compile and run a **Go** program
* Change into the directory with the program
* `go mod init day01`  (*only if go.mod does not yet exist*)
* `go build`
* `./day01`  (or whatever name of the executable)

To run a **Python** program
* Change into the directory with the program
* `python day06.py`

To compile and run a **Rust** program
* Change into the directory with the program
* `rustc day01.rs`
* `./day01`  (or whatever name of the executable)
* If the program requires external dependencies ("crates"), you will 
  have to do `cargo init`, move dayXX.rs to src, add the crate to
  Cargo.toml, and then `cargo build` to compile; the executable will
  be somewhere in the ./target directory.

To compile and run a **Zig** program
* Change into the zig directory, e.g., `cd day05/zig`
* `zig build` (*debug mode, day05 runs in ~15 mins*)
* `zig build -Doptimize=ReleaseFast` (*fast mode, takes < 2 mins*)
* To run: `zig-out/bin/day05`

To compile and run a **C** program
* `gcc -O3 day05.c -o day05`
* To run: `./day05`

AK, Dec 2023 and Oct 2024 (for Rust, Zig, C)
