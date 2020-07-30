package kruskal

import (
	"math"
	"sort"

	"github.com/meefer/kruskal-vis/unionfind"
)

// Recorder represents a step-by-step Kruskal's algorithm execution recorder
type Recorder interface {
	CheckEdge(from, to int)
	SetEdge(from, to int)
}

type edge [2]int
type byDistance struct {
	*Graph
	es []edge
}

func (g byDistance) d(i int) float64 {
	e := g.es[i]
	u, v := g.Nodes[e[0]], g.Nodes[e[1]]
	return math.Sqrt(float64((v.X-u.X)*(v.X-u.X) + (v.Y-u.Y)*(v.Y-u.Y)))
}
func (g byDistance) Len() int           { return len(g.es) }
func (g byDistance) Less(i, j int) bool { return g.d(i) < g.d(j) }
func (g byDistance) Swap(i, j int)      { g.es[i], g.es[j] = g.es[j], g.es[i] }

// Kruskal finds a minimum spanning tree for a connected weighted graph
func Kruskal(r Recorder, g *Graph) (sptree *Graph) {
	N := len(g.Nodes)
	uf := unionfind.NewUnionFind(N)

	var edges []edge
	for i, nodes := range g.Edges {
		for _, j := range nodes {
			if j > i {
				edges = append(edges, edge{i, j})
			}
		}
	}
	sort.Sort(byDistance{g, edges})

	sptree = &Graph{g.Nodes, make([][]int, N)}
	for _, e := range edges {
		r.CheckEdge(e[0], e[1])
		if !uf.Connected(e[0], e[1]) {
			r.SetEdge(e[0], e[1])

			sptree.SetEdge(e[0], e[1])
			uf.Union(e[0], e[1])
		}
	}

	return
}
