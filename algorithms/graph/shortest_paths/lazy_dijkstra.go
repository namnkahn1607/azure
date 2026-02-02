/* Algorithm: Lazy Dijkstra */

package sp

import (
	"azure/data_structures/graph"
	pq "azure/data_structures/priority_queue"
)

/*
Dijkstra Shortest Path on Weighted Digraph (Lazy variant).
- Time: O(E.logE) & Space: O(E + V).
*/
func LazyDijkstraSP(G *graph.Digraph, src int) *SP {
	sp := &SP{
		DistTo: make([]int, G.V),
		EdgeTo: make([]graph.Edge, G.V),
		Source: src,
	}

	for v := range G.V {
		sp.DistTo[v] = INF
	}

	type VerDist struct {
		vertex int
		dist   int
	}

	minpq := pq.NewPQ(func(a, b VerDist) bool {
		return a.dist < b.dist
	})

	sp.DistTo[src] = 0
	minpq.Enqueue(VerDist{src, sp.DistTo[src]})

	for !minpq.IsEmpty() {
		// Add the closest vertex to source.
		entry := minpq.Dequeue()
		v, dist := entry.vertex, entry.dist

		// Already computed a better distance -> Skip.
		if dist > sp.DistTo[v] {
			continue
		}

		for e := range G.Adjacent(v) {
			w := e.Other(v)
			newDist := sp.DistTo[v] + e.Weight()
			if newDist < sp.DistTo[w] {
				sp.DistTo[w] = newDist // Relax edge
				sp.EdgeTo[w] = e
				// Laziness: QUERY instead of UPDATE.
				minpq.Enqueue(VerDist{w, newDist})
			}
		}
	}

	return sp
}
