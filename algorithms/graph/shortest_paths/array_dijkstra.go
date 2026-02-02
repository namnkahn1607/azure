/* Algorithm: Array Dijkstra */

package sp

import "azure/data_structures/graph"

/*
Array-variant Dijkstra Shortest Path on Weighted Digraph.
- Time: O(V^2) & Space: O(V).
*/
func ArrayDijkstraSP(G *graph.Digraph, src int) *SP {
	sp := &SP{
		DistTo: make([]int, G.V),
		EdgeTo: make([]graph.Edge, G.V),
		Source: src,
	}

	marked := make([]bool, G.V)
	for v := range G.V {
		sp.DistTo[v] = INF
	}

	sp.DistTo[src] = 0

	for range G.V {
		// Choose the closest vertex to source.
		minV := -1
		minDist := INF

		for v, distV := range sp.DistTo {
			if !marked[v] && distV < minDist {
				minV = v
				minDist = distV
			}
		}

		// Mark computed vertex.
		marked[minV] = true

		// Relax all adjacent edges of the newly vertex.
		for e := range G.Adjacent(minV) {
			w := e.Other(minV)
			newDist := minDist + e.Weight()
			if !marked[w] && newDist < sp.DistTo[w] {
				sp.DistTo[w] = newDist
				sp.EdgeTo[w] = e
			}
		}
	}

	return sp
}
