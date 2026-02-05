/* Algorithm: Array Prim */

package mst

import "azure/data_structures/graph"

/*
Array-variant Prim's Minimum Spanning Tree on Undirected Graph.
- Time: O(V^2) & Space: O(V).
*/
func ArrayPrimMST(G *graph.Graph, src int) *MST {
	// Assume T is the greedily growing tree.
	mst := &MST{
		EdgeTo: make([]graph.Edge, G.V),
		Weight: 0,
	}

	marked := make([]bool, G.V)
	
	distTo := make([]int, G.V)
	for v := range G.V {
		distTo[v] = INF
	}

	distTo[src] = 0

	for {
		minV := -1
		minDist := INF

		for v := range G.V {
			// Select the closest non-tree vertex to tree.
			if !marked[v] && distTo[v] < minDist {
				minDist = distTo[v]
				minV = v
			}
		}

		// Found no more -> No further growing.
		if minV == -1 {
			break
		}

		// Add the vertex to tree T.
		marked[minV] = true
		mst.Weight += minDist

		// Distance update to all non-tree adjacent vertices.
		for e := range G.Adjacent(minV) {
			w := e.Other(minV)
			if !marked[w] && e.Weight() < distTo[w] {
				distTo[w] = e.Weight()
				mst.EdgeTo[w] = e
			}
		}
	}

	return mst
}
