/* Algorithm: Lazy Prim */

package mst

import (
	"azure/data_structures/graph"
	pq "azure/data_structures/priority_queue"
)

/*
Prim's Minimum Spanning Tree on Undirected Graph (Lazy variant).
- Time: O(E.logE) & Space: O(E).
*/
func LazyPrimMST(G *graph.Graph, src int) *MST {
	// Assume T is the greedily growing tree.
	mst := &MST{
		EdgeTo: make([]graph.Edge, G.V),
		Weight: 0,
	}

	marked := make([]bool, G.V)
	minpq := pq.NewPQ(func(a, b graph.Edge) bool {
		return a.Weight() < b.Weight()
	})
	
	scan := func(v int) {
		marked[v] = true
		for e := range G.Adjacent(v) {
			// Marked ~ Form cycle back to T -> Skip.
			if !marked[e.Other(v)] {
				minpq.Enqueue(e)
			}
		}
	}

	scan(src);

	for !minpq.IsEmpty() {
		// Add closest min-weight edge to T.
		e := minpq.Dequeue()
		v := e.Head()
		w := e.Other(v)

		if marked[v] && marked[w] { // Stale edge
			continue	
		}

		// One edge's endpoint belongs to T.
		mst.EdgeTo[w] = e
		mst.Weight += e.Weight()

		// Discover unmarked endpoint.
		if !marked[v] {
			scan(v)
		} else {
			scan(w)
		}
	}

	return mst
}
