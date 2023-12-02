// Advent of Code 2023, Day 02
//
// Given a list of games, each with a list of red/green/blue cubes
// drawn from a bag, determine which games are possible with a given
// number of red, green, and blue cubes. For Part 2, determine the
// minimum number of cubes of each colour required to play all the
// turns in a game, and multiply these together.
//
// AK, 2 Dec 2023 (43 mins)

package main

import (
	"fmt"
	"strings"
)

// A game is a series of turns
type Game struct {
	id    int
	turns []Turn
}

// Structure of a game turn, how many cubes of each colour
type Turn struct {
	red, green, blue int
}

func main() {

	// Read games from the input file
	fname := "sample.txt"
	fname = "input.txt"
	var games []Game
	for _, l := range readLines(fname) {
		g := parseGame(l)
		games = append(games, g)
	}

	// Part 1: Add up the IDs of all games that are possible with 12/13/14
	// red/green/blue cubes
	var ans1, ans2 int
	for _, g := range games {
		if possible(g, 12, 13, 14) {
			ans1 += g.id
		}
	}
	fmt.Println("Part 1:", ans1) // correct: 2278

	// Part 2: what is the fewest number of cubes of each color that could
	// have been in the bag to make the game possible?
	// Add up the "powers" of the minimum sets of cubes in each, game, where
	// The power of a set of cubes is equal to the numbers of red, green, and
	// blue cubes multiplied together.
	for _, g := range games {
		ans2 += power(g)
	}
	fmt.Println("Part 2:", ans2) // correct: 67953
}

// For Part 1, is game possible with given number of red, green, blue cubes?
func possible(g Game, red, green, blue int) bool {
	for _, t := range g.turns {
		if t.red > red || t.green > green || t.blue > blue {
			return false
		}
	}
	return true
}

// For Part 2, get the minimum number of cubes of each colour required to play
// all the turns in a game, and multiply these together
func power(g Game) int {
	mg := Turn{} // each colour initialized to zero
	for _, t := range g.turns {
		if t.red > mg.red { // if turn requires more of this colour,
			mg.red = t.red // update the minimum
		}
		if t.green > mg.green {
			mg.green = t.green
		}
		if t.blue > mg.blue {
			mg.blue = t.blue
		}
	}
	return mg.red * mg.green * mg.blue
}

// Parse a game
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
func parseGame(s string) Game {

	// Create a new game with the id
	parts := strings.Split(s, ":")
	gids := parts[0] // e.g., "Game xx"
	id := atoi(gids[5:])
	g := Game{id: id} // empty game (all counters zero)

	// Extract turns from the game
	s = parts[1]                   // strip off "Game xx:"
	turns := strings.Split(s, ";") // separate turns
	for _, ts := range turns {     // e.g., "8 green, 6 blue, 20 red"
		t := Turn{}                                 // new turn
		for _, cs := range strings.Split(ts, ",") { // e.g., "8 green"
			cs = strings.TrimSpace(cs)
			nc := strings.Split(cs, " ")
			n := atoi(nc[0]) // the number
			c := nc[1]       // the colour
			if c == "red" {
				t.red = n
			} else if c == "blue" {
				t.blue = n
			} else if c == "green" {
				t.green = n
			} else {
				fmt.Println("Invalid color:", c)
			}
		}
		g.turns = append(g.turns, t)
	}

	return g
}
