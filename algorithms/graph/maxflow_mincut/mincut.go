/* API: Maxflow Mincut */

package maxflow

import (
	graph "azure/data_structures/flow_network"
	"math"
)

const INF = math.MaxInt64

type MinCut struct {
	Edges    []*graph.FlowEdge
	Capacity int
}
