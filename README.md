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

To compile and run a **Go** program
* Change into the directory with the program
* `go mod init day01`  (*only if go.mod does not yet exist*)
* `go build`
* `./day01`  (or whatever name of the executable)

To run a **Python** program
* Change into the directory with the program
* `python day06.py`

AK, Dec 2023
