// Advent of Code 2023, Day 19
//
// Given a set of rules, and a set of parts, determine which parts
// are accepted by the rules. Rules are a series of tests, each of
// which is either a comparison (e.g., "s<1351:dest") or a destination.
// For Part 2, determine how many parts, within a universe of 0..4000
// for each of the four parameters, would be accepted. Did this by
// recursively evaluating the rules, to come up with a list of "and"
// conditions for each "A" terminal node. For each "A" rule, the number
// of parts accepted is the product of the widths of the four resulting
// ranges.
//
// AK, 19-27 Dec 2023

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
	cmp  string // comparator, e.g., "<", ">" or empty if destination
	n    int    // number to compare against
	dest string // destination if comparison is true
}

// Global dictionary of rules, each a list of sequential tests
var Rules map[string][]Test

// For Part 2, a global lists of lists of conditions that lead to acceptance
// or rejection
var Accepts [][]Test

func main() {

	// Read the input file
	fname := "sample.txt"
	fname = "input.txt"
	loadData(fname)

	// Part 1: Process each part, add up attributes of accepted parts
	ans := 0
	for _, p := range Parts {
		if processPart(p) == "A" {
			ans += p["x"] + p["m"] + p["a"] + p["s"]
		}
	}
	fmt.Println("Part 1 (19114, 397643):", ans) // 19114, 397643

	// Part 2: determine many parts, within a universe of 0..4000 for
	// each attribute, would be accepted
	ans2 := part2()
	fmt.Println("Part 2 (s/b 167409079868000, 132392981697081):", ans2)
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
		if len(t.cmp) == 0 {
			return t.dest
		}

		// Evaluate "s<1351:dest"
		val, ok := p[t.a]
		assert(ok, "Attribute not found")
		if t.cmp == "<" && val < t.n {
			return t.dest
		} else if t.cmp == ">" && val > t.n {
			return t.dest
		}
	}

	// Could not find an answer
	return "?"
}

// PART 2

// Part 2: determine how many parts could be accepted, out of
// a universe where each parameter ranges from 0 to 4000.
func part2() int64 {

	// Recursively enumerate rules, starting with rule "in", building up
	// chains of Accept conditions in the global variable, basically
	// root-to-terminal paths in decision tree
	enumerate(Rules["in"], []Test{})

	// Apply each set of "Accept" conditions to determine the ranges
	// implied for each of the four attributes. Multiply the width of
	// these four ranges together, to get the number of parts that
	// would match each set of tests.
	xmas := []string{"x", "m", "a", "s"}
	var ans int64
	for _, tests := range Accepts {

		// Initialize min & max of ranges for x,m,a,s with
		// 1 as the minimum, 4000 as the maximum.
		minVal := map[string]int{}
		maxVal := map[string]int{}
		for _, a := range xmas {
			minVal[a] = 1
			maxVal[a] = 4000
		}

		// Apply each condition, modifying the range limits for the
		// affected attribute accordingly
		for _, t := range tests {
			if t.cmp == "<" {
				if t.n-1 < maxVal[t.a] {
					maxVal[t.a] = t.n - 1
				}
			} else if t.cmp == "<=" {
				if t.n < maxVal[t.a] {
					maxVal[t.a] = t.n
				}
			} else if t.cmp == ">" {
				if t.n+1 > minVal[t.a] {
					minVal[t.a] = t.n + 1
				}
			} else if t.cmp == ">=" {
				if t.n > minVal[t.a] {
					minVal[t.a] = t.n
				}
			}
		}

		// The number of parts affected by this rule is just the
		// widths of the four ranges multiplied together
		var c int64 = 1 // the count for this rule
		for _, a := range xmas {
			rng := maxVal[a] - minVal[a] + 1 // range width
			c *= int64(rng)
		}
		ans += c
		//fmt.Println(tests, minVal, maxVal, c)
	}
	return ans
}

// Enumerate a rule, i.e., recursively build up lists of tests
// that lead to "A"
func enumerate(tests []Test, conds []Test) {

	// Make a copy of the conditions
	conds = copyTests(conds)

	// Process each test
	for _, t := range tests {

		// If just a destination with no comparison, add conditions
		// to "Accept" list if "A", or evaluate this destination
		if len(t.cmp) == 0 {
			if t.dest == "A" {
				Accepts = append(Accepts, conds)
				//fmt.Println("Found Accept chain:", conds)
			} else if t.dest != "R" { // rule destination without condition
				//fmt.Println("Evaluating non-condition rule", t.dest)
				enumerate(Rules[t.dest], copyTests(conds))
			}

			// Otherwise, evaluate left-hand side, then add negation of
			// this test to list of conditions
		} else {

			// Evaluate left hand side, with this test added to conditions,
			// so it can be included in further tests for this rule (right
			// branch implies negation of the left branch)
			conds1 := copyTests(conds)
			conds1 = append(conds1, t)
			if t.dest == "A" { // TODO: This seems redundant
				Accepts = append(Accepts, conds1)
				//fmt.Println("Found Accept chain:", conds1)
			} else if t.dest != "R" {
				//fmt.Println("Evaluating condition rule", t.dest)
				enumerate(Rules[t.dest], conds1)
			}

			// Add negation of this condition to the list of conditions
			conds = append(conds, negate(t))
		}
	}
}

// Make copy of a list of tests
func copyTests(tests []Test) []Test {
	var res []Test
	res = append(res, tests...)
	return res
}

// Negate a condition, by flipping >/<= etc.
func negate(t Test) Test {
	opposite := map[string]string{"<": ">=", "<=": ">", ">": "<=", ">=": "<"}
	cmp1, ok := opposite[t.cmp]
	if !ok {
		panic("Invalid operator")
	}
	t1 := t
	t1.cmp = cmp1
	return t1
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
	t.a = lr[0][:1]          // attribute name, e.g., "s"
	t.cmp = string(lr[0][1]) // comparator, e.g., '<'
	t.n = atoi(lr[0][2:])
	t.dest = lr[1] // destination for this test
	return t
}
