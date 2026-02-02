/* Data Structure: Edge */

package graph

type Edge struct {
	v, w   int
	weight int
}

func NewEdge(u, v, weight int) *Edge {
	return &Edge{u, v, weight}
}

func (e *Edge) Other(v int) int {
	switch v {
	case e.v:
		return e.w
	case e.w:
		return e.v
	}

	panic("invalid edge endpoint")
}

func (e *Edge) Weight() int {
	return e.weight
}
