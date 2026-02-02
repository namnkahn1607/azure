/* Algorithm: Eager Dijkstra */

package sp

import (
	"azure/data_structures/graph"
	pq "azure/data_structures/priority_queue"
)

/*
Dijkstra Shortest Path on Weighted Digraph (Eager variant).
- Time: O(E.logV) & Space: O(V).
*/
func EagerDijkstraSP(G *graph.Digraph, src int) *SP {
	sp := &SP{
		DistTo: make([]int, G.V),
		EdgeTo: make([]graph.Edge, G.V),
		Source: src,
	}

	for v := range G.V {
		sp.DistTo[v] = INF
	}

	sp.DistTo[src] = 0
	minpq := pq.NewIndexPQ(
		G.V, func(a, b int) bool { return a < b },
	)
	minpq.Enqueue(src, sp.DistTo[src])

	for !minpq.IsEmpty() {
		// Add the closest vertex to source.
		v, dist, ok := minpq.Dequeue()

		if !ok {
			panic("attempt to dequeue an empty PQ")
		}

		for e := range G.Adjacent(v) {
			w := e.Other(v)
			newDist := dist + sp.DistTo[v]
			if newDist < sp.DistTo[w] {
				sp.EdgeTo[w] = e
				sp.DistTo[w] = newDist // Relax edge

				// Keep minimum distance to each vertex.
				if minpq.Contains(w) {
					minpq.ChangeKey(w, newDist) // UPDATE
				} else {
					minpq.Enqueue(w, newDist) // QUERY
				}
			}
		}
	}

	return sp
}
