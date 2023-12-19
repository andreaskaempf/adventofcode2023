// Advent of Code 2023, Day 19
//
// AK, 19 Dec 2023

package main

import (
	"fmt"
	"strings"
)

// Global list of parts, where each part is represented as a dictionary x/m/a/s -> value
type Part map[string]int

var Parts []Part

// Global dictionary of rules, stored in text
var Rules map[string][]string

func main() {

	// Read the input file
	//loadData("sample.txt")
	loadData("input.txt")

	// Process each part
	ans := 0
	for _, p := range Parts {
		if processPart(p) == "A" {
			n := p["x"] + p["m"] + p["a"] + p["s"]
			ans += n
		}
	}
	fmt.Println("Part 1:", ans)
}

func processPart(p Part) string {

	fmt.Println("Checking part", p)

	r := Rules["in"] // start with rule "in"
	for {
		nextRule := applyRule(r, p)
		//fmt.Println(r, "=>", nextRule)
		if nextRule == "A" || nextRule == "R" {
			return nextRule
		}
		r = Rules[nextRule]
	}
	return "?"
}

// Apply one rule, return the target
// E.g., qqz{s>2770:qs,m<1801:hdj,R}
func applyRule(r []string, p Part) string {

	for _, test := range r {

		// Expect "s<1351:dest" for comparison
		// Otherwise it's a destination
		if !in(':', []byte(test)) {
			return test
		}

		// Evaluate "s<1351:dest"
		lr := strings.Split(test, ":")
		dest := lr[1]
		a := lr[0][:1]  // attribute name, e.g., "s"
		cmp := lr[0][1] // comparator, e.g., '<'
		n := atoi(lr[0][2:])
		val, ok := p[a]
		//fmt.Println(a, string(cmp), n, val)
		assert(ok, "Attribute not found")
		if cmp == '<' && val < n {
			return dest
		} else if cmp == '>' && val > n {
			return dest
		}
	}

	// Could not find an answer
	return "?"
}

// Read and parse the input data, into global variables.
// Rule: qqz{s>2770:qs,m<1801:hdj,R}
// Part: {x=787,m=2655,a=1222,s=2876}
func loadData(fname string) {

	Rules = make(map[string][]string)
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
			Rules[ab[0]] = strings.Split(ab[1], ",")
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
