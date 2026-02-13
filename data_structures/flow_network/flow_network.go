/* Data Structure: Flow Network */

package graph

import (
	"bufio"
	"io"
	"iter"
	"strconv"
)

type FlowNetwork struct {
	E, V int
	adj  [][]*FlowEdge
}

/* Create a Flow Network with V vertices. */
func NewFlowNetwork(V int) *FlowNetwork {
	if V < 0 {
		panic("negative number of vertices")
	}

	return &FlowNetwork{
		V:   V,
		adj: make([][]*FlowEdge, V),
	}
}

/* Create a Flow Network from input stream. */
func NewFlowNetworkIO(r io.Reader) *FlowNetwork {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	readInt := func() int {
		if !scanner.Scan() {
			panic("error reading input")
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("invalid integer format")
		}

		return val
	}

	V := readInt()
	if V < 0 {
		panic("negative number of vertices")
	}

	G := &FlowNetwork{
		V:   V,
		adj: make([][]*FlowEdge, V),
	}

	E := readInt()
	if E < 0 {
		panic("negative number of edges")
	}

	for range E {
		from := readInt()
		to := readInt()
		cap := readInt()
		G.AddEdge(FlowEdge{from, to, 0, cap})
	}

	return G
}

/* Add a flow edge onto the Flow Network. */
func (G *FlowNetwork) AddEdge(e FlowEdge) {
	v, w := e.from, e.to
	G.IsVertexOf(v)
	G.IsVertexOf(w)

	G.adj[v] = append(G.adj[v], &e)
	G.adj[w] = append(G.adj[w], &e)
	G.E++
}

/* Adjacency List (flow edges) of a given vertex. */
func (G *FlowNetwork) Adjacent(v int) iter.Seq[FlowEdge] {
	G.IsVertexOf(v)
	return func(yield func(FlowEdge) bool) {
		for _, e := range G.adj[v] {
			if !yield(*e) {
				return
			}
		}
	}
}

/* Validate if a vertex belongs to a Flow Network. */
func (G *FlowNetwork) IsVertexOf(v int) {
	if v < 0 || v >= G.V {
		panic("vertex out of bounds")
	}
}
