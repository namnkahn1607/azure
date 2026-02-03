/* Algorithm: Shortest Path Faster Algorithm */

package sp

import (
	"azure/data_structures/graph"
	stackqueue "azure/data_structures/stack_queue"
)

/*
Queue-optimized Shortest Path on negative edge weight Digraph.
- Time: O(k.E) average, O(E.V) worst & Space: O(V).
*/
func ShortestPathFasterSP(G *graph.Digraph, src int) *SP {
	sp := &SP{
		DistTo: make([]int, G.V),
		EdgeTo: make([]graph.Edge, G.V),
		Source: src,
	}

	for v := range G.V {
		sp.DistTo[v] = INF
	}

	// Avoid bloating by tracking relaxed vertices.
	queue := stackqueue.NewQueue[int](G.V)
	onQueue := make([]bool, G.V)
	// Count number of relaxation to each vertex.
	relaxCount := make([]int, G.V)

	sp.DistTo[src] = 0
	queue.Enqueue(src)
	onQueue[src] = true

	for !queue.IsEmpty() {
		v, ok := queue.Dequeue()
	
		if !ok {
			panic("attempt to dequeue an empty Queue")
		}

		onQueue[v] = false

		for e := range G.Adjacent(v) {
			w := e.Other(v)
			newDist := sp.DistTo[v] + e.Other(w)
			if newDist < sp.DistTo[w] {
				sp.DistTo[w] = newDist // Relax edge
				sp.EdgeTo[w] = e
			
				if !onQueue[w] {
					// Enqueue relaxed vertex, since its new distance
					// will be the clue to relax more vertices.
					queue.Enqueue(w)
					onQueue[w] = true
					relaxCount[w]++

					// This vertex has relaxed V times -> Negative Cycle.
					if relaxCount[w] >= G.V {
						panic("negative cycle detected")
					}
				}
			}
		}
	}

	return sp
}
