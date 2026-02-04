/* Algorithm: Topological Sort */

package topological

import (
	"azure/algorithms/array"
	"azure/data_structures/graph"
)

/*
Topological order of vertices of a Digraph (DFS).
- Time: O(E + V) & Space: O(V).
*/
func TopologicalDFS(G *graph.Digraph) []int {
	const (
		UNMARKED = 0
		MARKING  = 1
		MARKED   = 2
	)

	order := make([]int, 0, G.V)
	marked := make([]int, G.V)

	var dfs func(int) bool
	dfs = func(v int) bool {
		marked[v] = MARKING

		// Add its children first.
		for e := range G.Adjacent(v) {
			w := e.Other(v)

			if marked[w] == MARKING { // Cycle detection
				return false
			}

			// Explore if unmarked + early termination.
			if marked[w] == UNMARKED && !dfs(w) {
				return false
			}
		}

		// Append the parent afterward.
		order = append(order, v)
		marked[v] = MARKED

		return true
	}

	for v := range G.V {
		if marked[v] == UNMARKED && !dfs(v) {
			panic("non-acyclical input Digraph")
		}
	}

	array.Reverse(order)
	return order
}
