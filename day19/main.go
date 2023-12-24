// Advent of Code 2023, Day 19
//
// Given a set of rules, and a set of parts, determine which parts
// are accepted by the rules. Rules are a series of tests, each of
// which is either a comparison or a destination. A comparison is
// of the form "s<1351:dest", where "s" is an attribute, "<" is a
// comparator, "1351" is a number, and "dest" is a destination.
// For Part 2, determine how many parts, within a universe of 0..4000
// for each of the four parameters, would be accepted (*to do*).
//
// AK, 19 Dec 2023

package main

import (
	"fmt"
	"strings"
)

// Global list of parts, where each part is represented as a
// dictionary x/m/a/s -> value
type Part map[string]int

var Parts []Part

// A test within a rule, e.g., "s<1351:dest"
// Either a comparison witor a destination
type Test struct {
	a    string // attribute name, e.g., "s"
	cmp  byte   // comparator, e.g., '<', '>' or 0 if destination
	n    int    // number to compare against
	dest string // destination if comparison is true
}

// Global dictionary of rules, each a list of sequential tests
var Rules map[string][]Test

func main() {

	// Read the input file
	loadData("sample.txt")
	//loadData("input.txt")

	// Part 1: Process each part, add up attributes of accepted parts
	ans := 0
	for _, p := range Parts {
		if processPart(p) == "A" {
			ans += p["x"] + p["m"] + p["a"] + p["s"]
		}
	}
	fmt.Println("Part 1 (19114, 397643):", ans) // 19114, 397643

	// Part 2: determine many parts, within a universe of 0..4000,
	// would be accepted
	part2()
}

// PART 1

// For Part 1, process a part, return "A" if accepted, "R" if rejected
func processPart(p Part) string {

	//fmt.Println("Checking part", p)
	r := Rules["in"] // start with rule "in"
	for {
		nextRule := applyRule(r, p)
		if nextRule == "A" || nextRule == "R" {
			return nextRule
		}
		r = Rules[nextRule]
	}
	return "?"
}

// Apply one rule, return the target
// E.g., qqz{s>2770:qs,m<1801:hdj,R}
func applyRule(r []Test, p Part) string {

	for _, t := range r { // each test in the rule

		// Expect "s<1351:dest" for comparison
		// Otherwise it's a destination
		if t.cmp == 0 {
			return t.dest
		}

		// Evaluate "s<1351:dest"
		val, ok := p[t.a]
		assert(ok, "Attribute not found")
		if t.cmp == '<' && val < t.n {
			return t.dest
		} else if t.cmp == '>' && val > t.n {
			return t.dest
		}
	}

	// Could not find an answer
	return "?"
}

// PART 2

var Accepts, Rejects [][]Test

// Part 2: determine how many parts could be accepted, out of
// a universe where each parameter ranges from 0 to 4000.

// Take a rule (series of tests).
// First one: enumerate it, with zero-length "and" list of rules
// Enumerate takes the first rule in list:
// - if it's terminal (A or R), add it to the global list of terminals
// - otherwise, enumerate that rule, with list of "and" rules
// Subsequent rules:
// - take the previous rule, negate it, add it to the "and" list
// - enumerate that rule
func part2() {

	Accepts = [][]Test{}
	Rejects = [][]Test{}

	r := Rules["in"] // start with rule "in"
	enumerate(r, []Test{})

	fmt.Println("Accept:", Accepts)
	fmt.Println("Reject:", Rejects)

	/*for _, t := range r { // each test in the rule
		enumerate(t, []Test{}) // enumerate all rules
	}*/
	//fmt.Println("Part 2:", r)

}

// Enumerate a rule, i.e., recursively build up lists of tests
// that lead to "A" or "R"
// Rule at front of list gets enumerated, along with all the
func enumerate(tests []Test, prevCond []Test) {

	//fmt.Println("Enumerating", r, prevCond)
	// TODO: could be other terminal
	if len(tests) == 1 && tests[0].cmp == 0 {
		if tests[0].dest == "A" {
			Accepts = append(Accepts, prevCond)
		} else if tests[0].dest == "R" {
			Rejects = append(Rejects, prevCond)
		}
		return
	}

	// Enumerate the first rule
	prevCond = append(prevCond, tests[0])
	enumerate(tests[1:], prevCond)
}

// LOAD DATA

// Read and parse the input data, into global variables.
// Rule: qqz{s>2770:qs,m<1801:hdj,R}
// Part: {x=787,m=2655,a=1222,s=2876}
func loadData(fname string) {

	Rules = make(map[string][]Test)
	var readingParts bool
	for _, l := range readLines(fname) {

		//fmt.Println(string(l))
		if len(l) == 0 { // blank line separates rules and parts
			readingParts = true

			// Read a part
		} else if readingParts {
			part := map[string]int{}
			l = l[1 : len(l)-1] // strip off braces
			for _, attr := range strings.Split(l, ",") {
				a := attr[:1]            // the letter
				part[a] = atoi(attr[2:]) // assign number
			}
			Parts = append(Parts, part)

			// Read a rule, e.g., rfg{s<537:gd,x>2440:R,A}
		} else {
			l = l[:len(l)-1] // remove closing brace
			ab := strings.Split(l, "{")
			ruleName := ab[0]
			Rules[ruleName] = []Test{}
			for _, s := range strings.Split(ab[1], ",") {
				Rules[ruleName] = append(Rules[ruleName], parseTest(s))
			}
		}
	}
	/*fmt.Println("Rules:")
	for k, v := range Rules {
		fmt.Println(k, "=", v)
	}
	fmt.Println("Parts:")
	for _, p := range Parts {
		fmt.Println(p)
	}*/
}

// Parse a test, e.g., "s<1351:dest"
func parseTest(test string) Test {

	// Destination if no comparison
	var t Test
	if !in(':', []byte(test)) {
		t.dest = test
		return t
	}
	// Otherwise assign comparator and number
	lr := strings.Split(test, ":")
	t.a = lr[0][:1]  // attribute name, e.g., "s"
	t.cmp = lr[0][1] // comparator, e.g., '<'
	t.n = atoi(lr[0][2:])
	t.dest = lr[1] // destination for this test
	return t
}
