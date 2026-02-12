/* Data Structure: Flow Edge */

package graph

type FlowEdge struct {
	v, w      int
	flow, cap int
}

func NewFlowEdge(from, to, cap int) *FlowEdge {
	return &FlowEdge{
		v:   from,
		w:   to,
		cap: cap,
	}
}

/* Add an amount of residual value to one endpoint of a flow edge. */
func (e *FlowEdge) AddResidualTo(v int, delta int) {
	switch v {
	case e.v:
		e.flow -= delta
	case e.w:
		e.flow += delta
	}

	if e.flow < 0 || e.flow > e.cap {
		panic("edge overflow/underflow")
	}
}

/* Residual value to one endpoint of a flow edge. */
func (e *FlowEdge) ResidualTo(v int) int {
	switch v {
	case e.v:
		return e.flow
	case e.w:
		return e.cap - e.flow
	}

	panic("invalid edge endpoint")
}

/* Another endpoint of a vertex in a flow edge. */
func (e *FlowEdge) Other(v int) int {
	switch v {
	case e.v:
		return e.w
	case e.w:
		return e.v
	}

	panic("invalid edge endpoint")
}
