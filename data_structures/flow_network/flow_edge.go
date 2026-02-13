/* Data Structure: Directed Flow Edge */

package graph

type FlowEdge struct {
	from, to  int
	flow, cap int
}

/* Create a new hollow Directed Flow Edge. */
func NewFlowEdge(v, w, cap int) *FlowEdge {
	return &FlowEdge{v, w, 0, cap}
}

/* The head endpoint of the Flow Edge. */
func (e *FlowEdge) Head() int {
	return e.from
}

/* Another endpoint of a vertex in a Flow Edge. */
func (e *FlowEdge) Other(v int) int {
	switch v {
	case e.from:
		return e.to
	case e.to:
		return e.from
	}

	panic("invalid edge endpoint")
}

/* Add an amount of residual to a Flow Edge. */
func (e *FlowEdge) AddResidualTo(v int, delta int) {
	switch v {
	case e.from:
		e.flow -= delta
	case e.to:
		e.flow += delta
	default:
		panic("invalid edge endpoint")
	}

	if e.flow < 0 || e.flow > e.cap {
		panic("edge overflow/underflow")
	}
}

/* Residual value corresponding to a vertex in a Flow Edge. */
func (e *FlowEdge) ResidualTo(v int) int {
	switch v {
	case e.from:
		return e.to
	case e.to:
		return e.from
	}

	panic("invalid edge endpoint")
}
