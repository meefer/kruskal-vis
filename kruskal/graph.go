package kruskal

import (
	"fmt"
	"math/rand"
)

// Node represents a node of a graph
type Node struct {
	x, y float64
}

// Graph represents a graph math structure
type Graph struct {
	nodes []*Node
	edges [][]int
}

func (g *Graph) String() string {
	nodes := ""
	for _, n := range g.nodes {
		nodes += fmt.Sprintf("%v ", *n)
	}
	return fmt.Sprintf("%s%v", nodes, g.edges)
}

// NewGraph constructs new Graph value
func NewGraph(N int) *Graph { // TODO: fix graph edges generation & extract to external package
	nodes := make([]*Node, N)
	edges := make([][]int, N)

	for i := range nodes {
		nodes[i] = &Node{rand.Float64(), rand.Float64()}
	}
	for from := range edges {
		tos := make([]int, 0, N-1)
		for to := range nodes {
			connected := rand.Float64()
			if from != to && connected < 0.5 {
				l := len(tos)
				tos = tos[:l+1]
				tos[l] = to
			}
		}
		edges[from] = tos
	}

	return &Graph{nodes, edges}
}
