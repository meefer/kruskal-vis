package kruskal

import (
	"fmt"
	"math/rand"
)

// MaxNodeXY is a maximum value of x and y coordinates for a graph node
const MaxNodeXY = 1000

// Node represents a node of a graph
type Node struct {
	X, Y int
}

// Graph represents a graph math structure
type Graph struct {
	Nodes []*Node
	Edges [][]int
}

func (g *Graph) String() string {
	nodes := ""
	for _, n := range g.Nodes {
		nodes += fmt.Sprintf("%v ", *n)
	}
	return fmt.Sprintf("%s%v", nodes, g.Edges)
}

// NewGraph constructs new Graph value
func NewGraph(N int) *Graph {
	nodes := make([]*Node, N)
	edges := make([][]int, N)
	for i := range nodes {
		nodes[i] = &Node{rand.Intn(MaxNodeXY), rand.Intn(MaxNodeXY)}
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
	g.Nodes = append(g.Nodes, n)
	return len(g.Nodes) - 1
}

// SetEdge creates new edge in a graph
func (g *Graph) SetEdge(from, to int) {
	g.Edges[from] = append(g.Edges[from], to)
	g.Edges[to] = append(g.Edges[to], from)
}
