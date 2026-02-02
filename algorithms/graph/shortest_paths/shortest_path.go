/* API: Shortest Path */

package sp

import "azure/data_structures/graph"

const INF = 1<<63 - 1

type SP struct {
	EdgeTo []graph.Edge
	DistTo []int
	Source int
}
