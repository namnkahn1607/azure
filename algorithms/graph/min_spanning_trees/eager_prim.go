/* Algorithm: Eager Prim */

package mst

import (
	"azure/data_structures/graph"
	pq "azure/data_structures/priority_queue"
)

/*
Prim's Minimum Spanning Tree on Undirected Graph (Eager variant).
- Time: O(E.logV) & Space: O(V).
*/
func EagerPrimMST(G *graph.Graph, src int) *MST {
	// Assume T is the greedily growing tree.
	mst := &MST{
		EdgeTo: make([]graph.Edge, G.V),
		Weight: 0,
	}

	marked := make([]bool, G.V)
	minpq := pq.NewIndexPQ(G.V, func(a, b int) bool {
		return a < b
	})

	distTo := make([]int, G.V)
	for v := range G.V {
		distTo[v] = INF
	}

	distTo[src] = 0
	minpq.Enqueue(src, distTo[src])

	scan := func(v int) {
		marked[v] = true
		for e := range G.Adjacent(v) {
			w := e.Other(v)

			// Marked ~ Form cycle back to T -> Skip.
			if marked[w] {
				continue
			}

			// Minimum-weight incoming edge to each vertex.
			if e.Weight() < distTo[w] {
				distTo[w] = e.Weight()
				mst.EdgeTo[w] = e

				if minpq.Contains(w) {
					minpq.ChangeKey(w, e.Weight()) // UPDATE
				} else {
					minpq.Enqueue(w, e.Weight()) // QUERY
				}
			}
		}
	}

	for !minpq.IsEmpty() {
		// Add closest min-weight edge to T.
		v, weight, ok := minpq.Dequeue()

		if !ok {
			panic("attempt to dequeue an empty PQ")
		}

		mst.Weight += weight
		scan(v)
	}

	return mst
}
