/* Data Structure: Edge */

package graph

type Edge struct {
	v, w   int
	weight int
}

func NewEdge(u, v, weight int) *Edge {
	return &Edge{u, v, weight}
}

/*
Arbitrary endpoint for Undirected Edge;
Source endpoint for Directed Edge.
*/
func (e *Edge) Head() int {
	return e.v
}

/* Another endpoint of a vertex in an edge. */
func (e *Edge) Other(v int) int {
	switch v {
	case e.v:
		return e.w
	case e.w:
		return e.v
	}

	panic("invalid edge endpoint")
}

/* An edge's weight. */
func (e *Edge) Weight() int {
	return e.weight
}
