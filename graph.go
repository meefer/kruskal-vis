package main

import (
	"math/rand"

	"github.com/meefer/kruskal-vis/image"
	"github.com/meefer/kruskal-vis/kruskal"
)

func generateGraph(N int) (g *kruskal.Graph) {
	// generate a random graph with N nodes
	nodes := make([]*kruskal.Node, N)
	edges := make([][]int, N)

	for i := range nodes {
		nodes[i] = &kruskal.Node{X: rand.Intn(image.MaxWidth), Y: rand.Intn(image.MaxHeight)}
	}

	g = &kruskal.Graph{Nodes: nodes, Edges: edges}
	for from := range edges {
		for to := from + 1; to < N; to++ {
			connected := rand.Float64()
			if connected < 0.5 {
				g.SetEdge(from, to)
			}
		}
	}
	return
}
