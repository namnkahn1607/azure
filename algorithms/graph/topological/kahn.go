/* Algorithm: Kahn */

package topological

import (
	"azure/data_structures/graph"
	stackqueue "azure/data_structures/stack_queue"
)

/*
Topological order of vertices of a Digraph (BFS).
- Time: O(E + V) & Space: O(V).
*/
func TopologicalBFS(G *graph.Digraph) []int {
	indeg := make([]int, G.V)
	topo := make([]int, 0, G.V)

	queue := stackqueue.NewQueue[int](G.V)
	for v := range G.V {
		indeg[v] = G.Indegree(v)
		// Start with 'no-prerequisite' vertices.
		if indeg[v] == 0 {
			queue.Enqueue(v)
		}
	}

	for !queue.IsEmpty() {
		// Peel all surface vertices.
		len := queue.Len()
		for range len {
			v, ok := queue.Dequeue()
			
			if !ok {
				panic("attempt to dequeue an empty Queue")
			}

			// Append them into the order Array.
			topo = append(topo, v)

			for e := range G.Adjacent(v) {
				w := e.Other(v)
				indeg[w] -= 1
				// Neighbor become the new surface -> Query.
				if indeg[w] == 0 {
					queue.Enqueue(w)
				}
			}
		}
	}

	if len(topo) != G.V {
		panic("non-acyclical input Digraph")
	}

	return topo
}
