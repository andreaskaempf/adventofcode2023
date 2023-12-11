# Advent of Code 2023

My solutions for the Advent of Code 2023, 
see https://adventofcode.com/2023

* **Day 1** (Go): Find digits in a string, combine first and last digit on each
  line to make a number, and add these up. For part 2, look for embedded names
  of numbers as well, complicated as last & first letters are shared if they
  are the same.

* **Day 2** (Go): Given a list of games, each with a list of 
  red/green/blue cubes drawn from a bag, determine which games are possible 
  with a given number of red, green, and blue cubes. For Part 2, determine 
  the minimum number of cubes of each colour required to play all the
  turns in a game, and multiply these together.

* **Day 3** (Go): Given a 2-d surface of digits and special characters, 
  find any embedded numbers that are adjacent to special characters, and
  add them up. For Part 2, find any pairs of numbers that are adjacent to 
  the same asterisk, and sum up the products of these pairs.

* **Day 4** (Go): You are given a list of "cards", each with two lists of 
  numbers separated by a vertical bar. Some of the numbers in the right 
  list are also in the left list. For Part 1, calculate "points" for each 
  card by adding up 1, 2, 4, ... for each match. For Part 2, make a copy 
  of the next n cards, where n is the number of matches, and count up the 
  cards at the end.

* **Day 5** (Go): Given a list of "seeds" (numbers), go through a series 
  of transformations for each seed, based on tables. Each table is a tuple 
  (dst, src, n). If the number is within src ..  src+n,  it is assigned 
  to dst. If no dst is found, the dst is the same as src for that 
  transformation. For Part 1, find the lowest final transformation. 
  Part 2 is the same, but treat each pair of "seeds" as start and length 
  of a range. Brute force takes under 5 minutes to run in Go, but would be 
  faster with concurrency, memoization, or a smarter approach.

* **Day 6** (Go): Simulate a toy boat race, in which you hold down a button
  to recharge a motor, which lets the boat go faster. Distance is the speed
  muliplied by remaining time. Find the number of ways you can hold down the
  button, so as to exceed the previous maximum distance given. For part 2, just
  one race, but much bigger number (still almost instantaneous using brute
  force).

* **Day 7** (Go): Given some hands of 5 cards, sort them by poker rank. 
  For part 2, replace instances of J (joker) with which ever other letter
  from the hand, that yields the best poker hand type. Answer for both 
  parts is the sumproduct of the rank (sorted sequence) times a given bid.

* **Day 8** (Go): Read a list of left/right instructions, and the left/right
  adjacencies for a set of nodes. For  Part 1, calculate the number of steps
  required to get from node "AAA" to "ZZZ". For Part 2, calculate the first
  step where all routes from ??A lead to any ??Z. Part 2 solution finds and
  takes into account the repeating length of each ??A to ??Z, but uses brute
  force to find the time when they align, not very efficient.

* **Day 9** (Go): Given a list of numbers, repeatedly get the deltas between
  each successive number. Then, use the delta to extrapolate back upwards to
  the previous row. For Part 1, extrapolate at the end of each row, for Part 2,
  before the beginning. Answers are the sums of the extrapolations.

* **Day 10** (Go): Given a 2D terrain of symbols, start at 'S' and follow the
  chain of "pipe" shapes back to S. For Part 1, calculate the distance of the
  furthest-away shape, distance measured in either direction from start.  For
  Part 2, count up how many positions are *not* inside the shape (solved using
  ray tracing method).

* **Day 11** (Go): Given a 2D terrain of symbols, start at 'S' and follow the
  Given a 2D field of 'galaxies' surrounded by space, first 'expand' this by
  doubling each row/column that is empty, and add up the manhattan distance
  between every pair. For Part 2, expand by factor of a million instead of
  doubling.

To compile and run a **Go** program
* Change into the directory with the program
* `go mod init day01`  (*only if go.mod does not yet exist*)
* `go build`
* `./day01`  (or whatever name of the executable)

To run a **Python** program
* Change into the directory with the program
* `python day06.py`

AK, Dec 2023
