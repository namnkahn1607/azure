/* Algorithm: Acyclical Shortest Path */

package sp

import (
	"azure/algorithms/graph/topological"
	"azure/data_structures/graph"
)

/*
Acyclical Shortest Path on a DAG.
- Time: O(E + V) & Space: O(V).
*/
func AcyclicalSP(G *graph.Digraph, src int) *SP {
	// Always create the SP object first.
	sp := &SP{
		DistTo: make([]int, G.V),
		EdgeTo: make([]graph.Edge, G.V),
		Source: src,
	}

	for v := range G.V {
		sp.DistTo[v] = INF
	}

	// Topological order of DAG.
	topo := topological.TopologicalDFS(G)
	sp.DistTo[src] = 0

	for _, v := range topo {
		if sp.DistTo[v] != INF { // Only reachable vertex
			for e := range G.Adjacent(v) {
				w := e.Other(v)
				newDist := sp.DistTo[v] + e.Weight()
				if newDist < sp.DistTo[w] { 
					sp.DistTo[w] = newDist // Relax edge
					sp.EdgeTo[w] = e
				}
			}
		}
	}

	return sp
}
