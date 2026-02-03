/* Data Structure: Directed Graph */

package graph

import (
	"bufio"
	"io"
	"iter"
	"strconv"
)

type Digraph struct {
	E, V   int
	Adj    [][]Edge
	Outdeg []int
	Indeg  []int
}

/* Create a Directed Graph with V vertices. */
func NewDigraph(V int) *Digraph {
	if V < 0 {
		panic("negative number of vertices")
	}

	return &Digraph{
		E:      0,
		V:      V,
		Adj:    make([][]Edge, V),
		Outdeg: make([]int, V),
		Indeg:  make([]int, V),
	}
}

/* Create a Directed Graph from input stream. */
func NewDigraphIO(r io.Reader) *Digraph {
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

	G := &Digraph{
		V:   V,
		Adj: make([][]Edge, V),
	}

	E := readInt()
	if E < 0 {
		panic("negative number of edges")
	}

	for range E {
		from := readInt()
		to := readInt()
		weight := readInt()
		G.AddEdge(Edge{from, to, weight})
	}

	return G
}

/* Add an edge onto the Directed Graph. */
func (G *Digraph) AddEdge(e Edge) {
	from, to := e.v, e.w
	G.IsVertexOf(from)
	G.IsVertexOf(to)

	G.Adj[from] = append(G.Adj[from], e)
	G.Outdeg[from]++
	G.Indeg[to]++
	G.E++
}

/* Adjacency List (edges) of a given vertex. */
func (G *Digraph) Adjacent(v int) iter.Seq[Edge] {
	G.IsVertexOf(v)
	return func(yield func(Edge) bool) {
		N := len(G.Adj[v])
		for i := range N {
			if !yield(G.Adj[v][i]) {
				return
			}
		}
	}
}

/* All directed edges from a Directed Graph. */
func (G *Digraph) Edges() iter.Seq[Edge] {
	return func(yield func(Edge) bool) {
		for v := range G.V {
			for e := range G.Adjacent(v) {
				if !yield(e) {
					return
				}
			}
		}
	}
}

/* Make a reversed clone of a Directed Graph. */
func (G *Digraph) Reversed() *Digraph {
	G_R := &Digraph{
		V:   G.V,
		Adj: make([][]Edge, G.V),
	}

	for e := range G.Edges() {
		G_R.AddEdge(Edge{e.w, e.v, e.weight})
	}

	return G_R
}

/* Traverse the Directed Graph in Preorder fashion. */
func (G *Digraph) PreOrder() iter.Seq[int] {
	return func(yield func(int) bool) {
		marked := make([]bool, G.V)

		var dfs func(int)
		dfs = func(v int) {
			marked[v] = true
			if !yield(v) {
				return
			}

			for e := range G.Adjacent(v) {
				w := e.Other(v)
				if !marked[w] {
					dfs(w)
				}
			}
		}

		for v := range G.V {
			if !marked[v] {
				dfs(v)
			}
		}
	}
}

/* Traverse the Directed Graph in Postorder fashion. */
func (G *Digraph) PostOrder() iter.Seq[int] {
	return func(yield func(int) bool) {
		marked := make([]bool, G.V)

		var dfs func(int)
		dfs = func(v int) {
			marked[v] = true

			for e := range G.Adjacent(v) {
				w := e.Other(v)
				if !marked[w] {
					dfs(w)
				}
			}

			if !yield(v) {
				return
			}
		}

		for v := range G.V {
			if !marked[v] {
				dfs(v)
			}
		}
	}
}

/* Validate if a vertex belongs to a Directed Graph. */
func (G *Digraph) IsVertexOf(v int) {
	if v < 0 || v >= G.V {
		panic("vertex out of bounds")
	}
}
