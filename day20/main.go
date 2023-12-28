// Advent of Code 2023, Day 20
//
// Given a list of components and their outputs, simulate the execution of
// signal transmission between the components when a "button" is pressed.
// For Part 1, determine the total number of low & high signals sent over
// 1000 button presses (done by simple simulation). For Part 2, need to
// determine how many button presses until one "rx" module receives a
// signal. This was done by looking for cycles in the four inputs to the
// only input to rx, and multiplying these cycle lengths together to
// give the times when the cycles coincide.
//
// AK, 28 Dec 2023

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Global map of components
var comps map[string]*Component

// A component
type Component struct {
	ctype   byte           // %, &, or b for broadcaster
	name    string         // name of this component
	targets []string       // list of target names
	on      bool           // on/off state for flip-flops, initially off
	memory  map[string]int // memory state, for conjunction modules
}

// Processing needs to happen breadth-first, so create a global queue of
// messages to be processed. First item in the queue is always the next
// to be processed, and messages are added to the back of the queue.
var Q []Message

// A message to be added to the queue
type Message struct {
	sender, receiver string
	signal           int
}

// Low/high signals
const LOW int = 0
const HIGH int = 1

// Global counters for total number of pulses
var totalLow, totalHigh, totalRx int

// Make number of presses for Part 2 global, to facilitate cycle detection
var presses int

// For part 2, use a map to find out the first time each of the four
// inputs to the sole input of rx is 1.
var cycles map[string]int

func main() {

	// Read the input file into a list of components, use pointers to
	// component so we always change the state of the original.
	// "broadcaster" gets renamed to "roadcaster".
	// e.g. &tb -> sx, qn, vj, qq, sk, pv
	fname := "sample2.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	comps = make(map[string]*Component)
	for _, l := range strings.Split(string(data), "\n") {
		l = strings.ReplaceAll(l, ",", "")
		words := strings.Split(l, " ")
		c := Component{ctype: words[0][0], name: words[0][1:], targets: words[2:]}
		comps[c.name] = &c
	}
	fmt.Println(len(comps), "components read")

	// Create a dummy components for "output" (used by samples) and
	// "rx" (ignored for Part 1, but used by Part 2).
	comps["output"] = &Component{ctype: 'o', name: "output"}
	comps["rx"] = &Component{ctype: 'o', name: "rx"} // found in input when trying to run

	// Part 1: press the button 1000 times, show product of total number of
	// low & high pulses sent (s/b 731517480)
	resetAll()
	for i := 0; i < 1000; i++ {
		pressButton()
	}
	fmt.Println("Part 1:", totalLow, "low, ", totalHigh, "high =>", totalLow*totalHigh)

	// Part 2: what is the fewest number of button presses required to deliver
	// a single low pulse to the module named rx? There is only one input to rx,
	// and it has 4 inputs, all conjunctors. When these are all 1, a pulse will
	// be emitted which turns on rx. So run a few thousand iterations, until we
	// get the four cycle lengths.
	resetAll()
	cycles = map[string]int{} // initialize map
	for presses < 10000 {     // or until totalRx == 0 for brute force
		presses++ // global counter
		pressButton()
	}

	// Multiply the four cycles found to get the answer for Part 2
	fmt.Println("Cycles found:", cycles)
	ans2 := 1
	for _, v := range cycles {
		ans2 *= v
	}
	fmt.Println("Part 2:", ans2) // 244178746156661
}

// Press the button, i.e., create a message to push the button, and
// start processing the queue
func pressButton() {
	Q = append(Q, Message{"button", "roadcaster", LOW})
	for len(Q) > 0 {
		m := Q[0] // oldest message at head of queue
		Q = Q[1:] // remove it from the queue
		processMessage(m)
	}
}

// Process a message, which creates messages to send a signal each of its
// receiver components, and add this to the end of the queue
func processMessage(m Message) {

	// For detecting cycles in Part 2, show whenever one of the inputs to
	// the only input to rx is 1. On the input file, the only input to rx
	// is &kh. The four inputs to these are &pv, &qh, &xm, and &hz. The
	// first time this happens is the length of the cycle for this input.
	// rx will turn on when all of these happen at the same time.
	inputs := []string{"pv", "qh", "xm", "hz"}
	if in(m.receiver, inputs) && m.signal != 1 {
		if cycles[m.receiver] == 0 { // catch the first time it happens
			cycles[m.receiver] = presses // for each input
		}
	}

	// Update the pulse totals, for answer
	if m.signal == LOW {
		totalLow++
	} else {
		totalHigh++
	}

	// Update total number of low pulses sent to "rx", for Part 2
	// brute force (takes too long, never reaches this)
	if m.receiver == "rx" && m.signal == LOW {
		fmt.Println("RX received low pulse!")
		totalRx++
	}

	// Get the named receiver module
	comp, ok := comps[m.receiver]
	assert(ok, "Receiver "+m.receiver+" not found")

	// Broadcaster just sends its signal to each destination
	if comp.ctype == 'b' {
		for _, targ := range comp.targets {
			Q = append(Q, Message{m.receiver, targ, m.signal})
		}
	} else if comp.ctype == '%' { // flip flop
		if m.signal == LOW { // ignore high-pulse signals
			sig := HIGH  // if off, turn on and send HIGH
			if comp.on { // if on, turn off and send LOW
				sig = LOW
			}
			comp.on = !comp.on // toggle state
			for _, targ := range comp.targets {
				Q = append(Q, Message{m.receiver, targ, sig})
			}
		}
	} else if comp.ctype == '&' { // conjunction

		// Remember input received, and who it came from
		comp.memory[m.sender] = m.signal

		// Send signal to each target, LOW if the last input from each
		// source was HIGH, otherwise send a HIGH
		sig := HIGH
		if allHigh(comp.memory) {
			sig = LOW
		}
		for _, targ := range comp.targets {
			Q = append(Q, Message{m.receiver, targ, sig})
		}
	}
}

// Are all values in the map HIGH?
func allHigh(m map[string]int) bool {
	for _, v := range m {
		if v != HIGH {
			return false
		}
	}
	return true
}

// Reset all components to their initial state
func resetAll() {

	// Initialize "on" state for flip-flops and map of upstream senders for
	// each conjunction module
	for _, c := range comps { // initialize empty map for each conjunction
		c.on = false // state for flip-flops, initially off
		if c.ctype == '&' {
			c.memory = map[string]int{}
		}
	}

	// Update conjunction map with its senders
	for nm, c := range comps { //
		for _, targ := range c.targets {
			t, ok := comps[targ] // the downstream receiver
			assert(ok, targ+" not found when initializing memory")
			if t.ctype == '&' { // if it's a conjunction
				t.memory[nm] = LOW
			}
		}
	}
}

// Is element in a list?
func in[T int | float64 | byte | string](c T, s []T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

// Panic if a test condition is not true
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
