/* Algorithm: Kruskal */

package mst

import (
	"azure/data_structures/graph"
	"cmp"
	"slices"
)

/*
Kruskal's Minimum Spanning Tree of an Undirected Graph.
- Time: O(E.logE) & Space: O(E).
*/
func KruskalMST(G *graph.Graph) *MST {
	mst := &MST{
		EdgeTo: make([]graph.Edge, G.V),
		Weight: 0,
	}

	rank := make([]int, G.V)
	parent := make([]int, G.V)
	for i := range G.V {
		parent[i] = i
	}

	find := func(i int) int {
		for i != parent[i] {
			parent[i] = parent[parent[i]] // Half-way compression
			i = parent[i]
		}

		return parent[i]
	}

	union := func(x, y int) bool {
		rx, ry := find(x), find(y)
		if rx == ry {
			return false
		}

		diff := rank[rx] - rank[ry] // Union by Rank
		if diff < 0 {
			parent[rx] = ry
		} else if diff > 0 {
			parent[ry] = rx
		} else {
			parent[ry] = rx
			rank[ry]++
		}

		return true
	}

	edges := make([]graph.Edge, 0, G.E)
	for e := range G.Edges() {
		edges = append(edges, e)
	}

	// Ascending edge weights sorting.
	slices.SortFunc(edges, func(a, b graph.Edge) int {
		return cmp.Compare(a.Weight(), b.Weight())
	})

	for _, e := range edges {
		v := e.Head()
		w := e.Other(v)
		if union(v, w) {
			mst.EdgeTo[w] = e
			mst.Weight += e.Weight()
		}
	}

	return mst
}
