// Advent of Code 2023, Day 25
//
// Given a graph of nodes ("components"), find the three edges ("wires") you
// need to disconnect in order to divide the graph into two separate groups,
// multiply the sizes of these two groups together. Used yourbasic/graph
// to create graph, and calculated the total distance from node 0 to all
// other nodes, tried all combinations of the 100 nodes with the highest
// graph distnce when removed, to check if graph has 2 components (i.e.,
// subgraphs not connected together). Picking the edges that resulted in
// the highest total graph distance was necessary because it's too much
// computation to consider all 3-edge combinations amount the ~3500
// edges.
//
// AK, 25 Dec 2023

package main

import (
	"fmt"
	"github.com/yourbasic/graph" // Graph library
	"io/ioutil"
	"sort"
	"strings"
)

// List of node names, mainly to get the total number of nodes for
// initializing the graph
var nodes []string

// An edge, with source and destination node IDs, and the total distance
// of the graph when the edge is removed (used for prioritizing which
// edges to try removing)
type Edge struct {
	src, dst int
	dist     int
}

func main() {

	// Read the input file into a list of strings
	fname := "sample.txt" // 15 nodes
	fname = "input.txt"   // 1550 nodes
	data, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(data), "\n")

	// Convert input to a graph and make a list of edges
	g := graph.New(1550) // 1550 for input, 15 for sample
	var edges []Edge
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		var srcId, dstId int
		for i, w := range strings.Split(l, " ") {
			w = strings.TrimSpace(w)
			if i == 0 { // first word is source
				src := w[:len(w)-1] // remove colon
				srcId = nodeId(src)
			} else { // rest are destinations from same source
				dstId = nodeId(w)
				g.AddBothCost(srcId, dstId, 1) // cost 1 to get distance
				edges = append(edges, Edge{srcId, dstId, 0})
			}
		}
	}
	fmt.Println(len(nodes), "nodes, ", len(edges), "edges")

	// For each edge, calculate the total distances from node 0 if this
	// edge is removed
	fmt.Println("Finding edges that contribute the largest distances")
	for i, e := range edges {
		g.DeleteBoth(e.src, e.dst)          // remove the edge
		edges[i].dist = totalDistance(g, 0) // total distance from node 0
		g.AddBothCost(e.src, e.dst, 1)      // add the edge back
	}

	// Sort the edges by descending distance
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist > edges[j].dist
	})

	// Remove all combinations of 3 edges until graph is divided
	// into two components, processing edges of descending total graph
	// distances when the edge is removed
	fmt.Println("Finding the three edges that divide the graph")
	var iter int
	for i := 0; i < len(edges); i++ {
		for j := 0; j < len(edges); j++ {
			for k := 0; k < len(edges); k++ {

				// All three edges must be different
				if i == j || i == k || j == k {
					continue
				}

				// Remove the three edges and check if graph is divided
				g.DeleteBoth(edges[i].src, edges[i].dst)
				g.DeleteBoth(edges[j].src, edges[j].dst)
				g.DeleteBoth(edges[k].src, edges[k].dst)

				// Get the number of components in the graph, show
				// answer if more than 1
				iter++
				cc := graph.Components(g)
				if len(cc) > 1 {
					fmt.Println("Iteration", iter, ": removing edges", edges[i], edges[j], edges[k])
					fmt.Println("  results in two components of size", len(cc[0]), len(cc[1]))
					fmt.Println("Answer =", len(cc[0])*len(cc[1]))
					return
				}

				// Otherwise, add the edges back
				g.AddBothCost(edges[i].src, edges[i].dst, 1)
				g.AddBothCost(edges[j].src, edges[j].dst, 1)
				g.AddBothCost(edges[k].src, edges[k].dst, 1)
			}
		}
	}
}

// Get the total distance from node 0 to all other nodes in the graph,
// using yourbasic/graph's shortest paths function
func totalDistance(g *graph.Mutable, node int) int {
	_, dists := graph.ShortestPaths(g, node)
	total := 0
	for _, d := range dists {
		total += int(d)
	}
	return total
}

// Number of this node, basically index in list, add if not there
func nodeId(node string) int {
	for i, n := range nodes {
		if n == node {
			return i
		}
	}
	nodes = append(nodes, node)
	return len(nodes) - 1
}
