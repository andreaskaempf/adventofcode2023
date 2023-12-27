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
	cmp  string // comparator, e.g., "<", ">" or empty if destination
	n    int    // number to compare against
	dest string // destination if comparison is true
}

// Global dictionary of rules, each a list of sequential tests
var Rules map[string][]Test

// For Part 2, a global lists of lists of conditions that lead to acceptance
var Accepts [][]Test

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
	ans2 := part2()
	fmt.Println("Part 2 (167409079868000):", ans2)
	fmt.Println(float64(ans2) / float64(167409079868000))
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

	// Recursively enumerate, starting with rule "in", building up
	// chains of Acceptance conditions in the global variable Accept
	enumerate(Rules["in"], []Test{})

	// Process each list of Accept conditions, to get the number of
	// parts that would be accepted, assuming that x/m/a/s can each range
	// from 1..4000
	var ans int64
	xmas := []string{"x", "m", "a", "s"}
	for _, tests := range Accepts {

		fmt.Println(tests)

		// Create a list of 4000 zeros for each of the four attributes
		bits := map[string][]int{}
		for _, a := range xmas {
			bits[a] = fill(4000, 0)
		}

		// Apply each test
		for _, t := range tests {
			if t.cmp == "<" {
				setBits(bits[t.a], 1, t.n-1, 1) // turn on 1..n-1
			} else if t.cmp == "<=" {
				setBits(bits[t.a], 1, t.n, 1) // turn on 1..n
			} else if t.cmp == ">" {
				setBits(bits[t.a], t.n+1, 4000, 1) // turn on n+1..4000
			} else if t.cmp == ">=" {
				setBits(bits[t.a], t.n, 4000, 1) // turn on n..4000
			} else {
				panic("Invalid operator: " + t.cmp)
			}
		}

		// Any attribute that has no bits set is assumed to be all ones
		for _, a := range xmas {
			if sum(bits[a]) == 0 {
				setBits(bits[a], 1, 4000, 1)
			}
		}

		// Determine the total number of parts enabled by this rule
		// TODO: Isn't this double counting, as the same parts may be
		// affected by multiple rules?
		var accepts int64 = 1
		for _, a := range xmas {
			accepts *= sum(bits[a])
		}

		fmt.Println("Subtotal for this rule:", accepts)
		ans += accepts
	}

	return ans
}

// Enumerate a rule, i.e., recursively build up lists of tests
// that lead to "A" or "R"
// Rule at front of list gets enumerated, along with all the
func enumerate(tests []Test, conds []Test) {

	// Make a copy of the conditions
	conds = copyTests(conds)

	// Process each test
	for _, t := range tests {

		// If no comparison, add to list if Accept, or evaluate
		// destination
		if len(t.cmp) == 0 {
			if t.dest == "A" {
				Accepts = append(Accepts, conds)
			} else if t.dest != "R" { // ignore Reject conditions
				enumerate(Rules[t.dest], copyTests(conds))
			}

			// Otherwise, evaluate left-hand side, then add negation of
			// this test to list of conditions
		} else {
			conds1 := copyTests(conds)
			conds1 = append(conds1, t)
			enumerate(Rules[t.dest], conds1)
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

// Turn on bits (numbers) in an array, setting them to 1
// Numbers are 1 indexed, but array indices are zero based
func setBits(bits []int, start, end, value int) {
	for i := start - 1; i < end-1; i++ {
		bits[i] = value
	}
}

// Create an array filled with given value
func fill(n, value int) []int {
	res := make([]int, n, n)
	for i := 0; i < n; i++ {
		res[i] = value
	}
	return res
}

// Sum up a list of ints
func sum(nums []int) int64 {
	var res int
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return int64(res)
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
	t.a = lr[0][:1]          // attribute name, e.g., "s"
	t.cmp = string(lr[0][1]) // comparator, e.g., '<'
	t.n = atoi(lr[0][2:])
	t.dest = lr[1] // destination for this test
	return t
}
