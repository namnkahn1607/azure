/* Algorithm: Edmonds-Karp */

package maxflow

import (
	graph "azure/data_structures/flow_network"
	stackqueue "azure/data_structures/stack_queue"
)

/*
A Shortest Path variant of Ford-Fulkerson for finding Maxflow.
- Time: O(E^2.V) & Space: O(V).
*/
func EdmondsKarpMaxFlow(G *graph.FlowNetwork, s, t int) *MinCut {
	ek := &EdmondsKarp{
		edgeTo: make([]*graph.FlowEdge, G.V),
		marked: make([]bool, G.V),
	}
	flowValue := 0

	// Algorithm proceed when there's augmenting path.
	for ek.hasAugmentingPath(G, s, t) {
		bottleneck := INF

		// Calculate bottleneck of path.
		for v := t; v != s; v = ek.edgeTo[v].Other(v) {
			e := ek.edgeTo[v]
			bottleneck = min(bottleneck, e.ResidualTo(v))
		}

		// Add delta residual flow to path.
		for v := t; v != s; v = ek.edgeTo[v].Other(v) {
			e := ek.edgeTo[v]
			e.AddResidualTo(v, bottleneck)
		}

		flowValue += bottleneck
	}

	// All bottleneck edges of the Flow Network.
	var mincut []*graph.FlowEdge
	for v := range G.V {
		if ek.marked[v] {
			for e := range G.Adjacent(v) {
				w := e.Other(v)
				if w != e.Head() && !ek.marked[w] {
					mincut = append(mincut, &e)
				}
			}
		}
	}

	return &MinCut{
		Capacity: flowValue,
		Edges:    mincut,
	}
}

/* Check & find Shortest augmenting path */
func (ek *EdmondsKarp) hasAugmentingPath(G *graph.FlowNetwork, s, t int) bool {
	for v := range G.V {
		ek.marked[v] = true
		ek.edgeTo[v] = nil
	}

	queue := stackqueue.NewQueue[int](G.V)
	queue.Enqueue(s)
	ek.marked[s] = true

	for !queue.IsEmpty() {
		v, ok := queue.Dequeue()

		if !ok {
			panic("attempt to dequeue an empty Queue")
		}

		for e := range G.Adjacent(v) {
			w := e.Other(v)

			// Avoid full forward & empty backward flow edges.
			if !ek.marked[w] && e.ResidualTo(w) > 0 {
				ek.edgeTo[w] = &e
				ek.marked[w] = true
				queue.Enqueue(w)
			}
		}
	}

	return ek.marked[t]
}

type EdmondsKarp struct {
	edgeTo []*graph.FlowEdge
	marked []bool
}
