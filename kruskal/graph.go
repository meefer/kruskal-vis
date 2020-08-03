package kruskal

import (
	"fmt"
)

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
