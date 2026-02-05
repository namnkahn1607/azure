/* API: Minimum Spanning Tree */

package mst

import "azure/data_structures/graph"

const INF = 1<<63 - 1

type MST struct {
	EdgeTo []graph.Edge
	Weight int
}
