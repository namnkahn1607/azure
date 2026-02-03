/* Algorithm: Bellman-Ford */

package sp

import "azure/data_structures/graph"

/*
Shortest Path on negative edge weight Digraph.
- Time: O(E.V) & Space: O(V).
*/
func BellmanFordSP(G *graph.Digraph, src int) *SP {
	sp := &SP{
		DistTo: make([]int, G.V),
		EdgeTo: make([]graph.Edge, G.V),
		Source: src,
	}

	for v := range G.V {
		sp.DistTo[v] = INF
	}

	sp.DistTo[src] = 0

	// Shortest path can't be longer than V-1.
	for range G.V - 1 {
		relaxed := false

		// Relax all edges in Graph.
		for v := range G.V {
			for e := range G.Adjacent(v) {
				w := e.Other(v)
				newDist := sp.DistTo[v] + e.Weight()
				if newDist < sp.DistTo[w] {
					sp.DistTo[w] = newDist // Relax edge
					sp.EdgeTo[w] = e
					relaxed = true
				}
			}
		}

		if !relaxed { // Early termination
			break
		}
	}

	// Can relax even more -> Negative cycle.
	for v := range G.V {
		if sp.DistTo[v] == INF {
			continue
		}

		for e := range G.Adjacent(v) {
			if sp.DistTo[e.Other(v)] < sp.DistTo[v] + e.Weight() {
				panic("negative cycle detected")
			} 
		}
	}

	return sp
}
