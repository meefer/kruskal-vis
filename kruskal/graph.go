package kruskal

import (
	"fmt"
	"math/rand"
)

const maxNodeXY = 1000

// Node represents a node of a graph
type Node struct {
	x, y int
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
func NewGraph(N int) *Graph {
	nodes := make([]*Node, N)
	edges := make([][]int, N)
	for i := range nodes {
		nodes[i] = &Node{rand.Intn(maxNodeXY), rand.Intn(maxNodeXY)}
	}

	g := &Graph{nodes, edges}
	for from := range edges {
		for to := from + 1; to < N; to++ {
			connected := rand.Float64()
			if connected < 0.5 {
				g.SetEdge(from, to)
			}
		}
	}

	return &Graph{nodes, edges}
}

// AddNode adds new vertex to a graph
func (g *Graph) AddNode(n *Node) int {
	g.nodes = append(g.nodes, n)
	return len(g.nodes) - 1
}

// SetEdge creates new edge in a graph
func (g *Graph) SetEdge(from, to int) {
	g.edges[from] = append(g.edges[from], to)
	g.edges[to] = append(g.edges[to], from)
}
