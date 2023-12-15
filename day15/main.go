// Advent of Code 2023, Day 15
//
// For Part 1, sum up the hashes of a list of commands, using a custom has
// function. For Part 2, simulate the movement of "lenses" among a chain of
// 256 boxes, by executing the commands.
//
// AK, 15 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

// A command
type Command struct {
	label string
	op    byte
	num   int
	text  []byte
}

// A lens
type Lens struct {
	label  string
	focLen int
}

// A box of lenses
type Box struct {
	lenses []Lens
}

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	lines := bytes.Split(data, []byte{'\n'})

	// Flatten into list commands
	var cmds []Command
	for _, s := range lines {
		for _, cmd := range bytes.Split(s, []byte{','}) {
			cmds = append(cmds, parseCmd(cmd))
		}
	}

	// Part 1: sum of the "hash" of every substring of comma-separated data
	ans := 0
	for _, cmd := range cmds {
		ans += hash(cmd.text)
	}
	fmt.Println("Part 1:", ans) // 511343

	// Part 2: simulate steps on the boxes
	boxes := make([]Box, 256, 256) // row of 256 "boxes"
	for _, c := range cmds {

		// Box number is hash of label
		label := c.label
		boxNo := hash([]byte(label))

		if c.op == '-' { // Dash means remove lens from the box
			boxes[boxNo] = removeLens(boxes[boxNo], c.label)
		} else { // Equal sign means update/add lens
			focLen := c.num // focal length
			boxes[boxNo] = updateLens(boxes[boxNo], label, focLen)
		}
	}

	// Part 2 answer is the sum of the "focal length" of all the lenses
	ans2 := 0
	for i := 0; i < len(boxes); i++ {
		box := boxes[i]
		for j := 0; j < len(box.lenses); j++ {
			lens := box.lenses[j]
			focPower := (i + 1) * (j + 1) * lens.focLen
			ans2 += focPower
		}
	}
	fmt.Println("Part 2:", ans2) // 294474
}

// Hash algorithm from problem
// Start with zero
// Ignore newlines (assumed removed from input)
// For each character:
// 1. Increase the current value by the ASCII code of the current character.
// 2. Set the current value to itself multiplied by 17.
// 3. Set the current value to the remainder of dividing itself by 256.
func hash(s []byte) int {
	h := 0
	for _, c := range s {
		h += int(c)
		h *= 17
		h = h % 256
	}
	return h
}

// Remove lens with given label from the box
func removeLens(b Box, label string) Box {
	b1 := Box{}
	for _, lens := range b.lenses {
		if lens.label != label {
			b1.lenses = append(b1.lenses, lens)
		}
	}
	return b1
}

// Update lens in a box: if there is already a lens with this label, \
// update its focal length. Otherwise, add lens to box.
func updateLens(b Box, label string, focLen int) Box {

	// Update lens with existing label
	for i := 0; i < len(b.lenses); i++ {
		if b.lenses[i].label == label {
			b.lenses[i].focLen = focLen
			return b
		}
	}

	// Add new lens to box
	lens := Lens{label, focLen}
	b.lenses = append(b.lenses, lens)
	return b
}

// Parse a command
func parseCmd(s []byte) Command {
	for i := 1; i < len(s); i++ {
		if s[i] == '-' || s[i] == '=' {
			label := string(s[:i])
			op := s[i] // - or =
			num, err := strconv.Atoi(string(s[i+1:]))
			if err != nil {
				num = -1 // indicates no number present
			}
			return Command{label, op, num, s}
		}
	}
	fmt.Println(string(s))
	panic("Invalid command")
}

// For debugging
func printBoxes(boxes []Box) {
	for i, b := range boxes {
		if len(b.lenses) > 0 {
			fmt.Println(i, b.lenses)
		}
	}
}
