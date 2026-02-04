/* Data Structure: Undirected Graph */

package graph

import (
	"bufio"
	"io"
	"iter"
	"strconv"
)

type Graph struct {
	E, V int
	adj  [][]Edge
	deg  []int
}

/* Create a Undirected Graph with V vertices. */
func NewGraph(V int) *Graph {
	if V < 0 {
		panic("negative number of vertices")
	}

	return &Graph{
		E:   0,
		V:   V,
		adj: make([][]Edge, V),
		deg: make([]int, V),
	}
}

/* Create a Undirected Graph from input stream. */
func NewGraphIO(r io.Reader) *Graph {
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

	G := &Graph{
		V:   V,
		adj: make([][]Edge, V),
		deg: make([]int, V),
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

/* Add an edge onto the Undirected Graph. */
func (G *Graph) AddEdge(e Edge) {
	from, to := e.v, e.w
	G.IsVertexOf(from)
	G.IsVertexOf(to)

	G.adj[from] = append(G.adj[from], e)
	G.deg[from]++
	G.adj[to] = append(G.adj[to], e)
	G.deg[to]++
	G.E++
}

/* Adjacency List (edges) of a given vertex. */
func (G *Graph) Adjacent(v int) iter.Seq[Edge] {
	G.IsVertexOf(v)
	return func(yield func(Edge) bool) {
		N := len(G.adj[v])
		for i := range N {
			if !yield(G.adj[v][i]) {
				return
			}
		}
	}
}

/* Degree of a Undirected Graph's vertex. */
func (G *Graph) Degree(v int) int {
	G.IsVertexOf(v)
	return G.deg[v]
}

/* All edges from a Undirected Graph. */
func (G *Graph) Edges() iter.Seq[Edge] {
	return func(yield func(Edge) bool) {
		for v := range G.V {
			for e := range G.Adjacent(v) {
				if v > e.Other(v) {
					if !yield(e) {
						return
					}
				}
			}
		}
	}
}

/* Validate if a vertex belongs to a Directed Graph. */
func (G *Graph) IsVertexOf(v int) {
	if v < 0 || v >= G.V {
		panic("vertex out of bounds")
	}
}
