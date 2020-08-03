package main

import (
	"flag"
	"fmt"
	"image/color"
	"math/rand"
	"os"

	gimage "github.com/meefer/kruskal-vis/image"
	"github.com/meefer/kruskal-vis/kruskal"
)

const usage = `%s generates a random graph with a fixed number of nodes
and runs it through Kruskal's minimum-spanning-tree algorithm
producing three files in a working directory:
	- graph.png -- original graph
	- kruskal_anim.gif -- visualization of the Kruskal's algorithm 
	- kruskal.png -- minimum spanning tree of the original graph
Usage:
`
const insufficientRights = "don't have sufficient rights to create a file"

var (
	n = flag.Int("N", 10, "number of graph nodes")
	t = flag.Bool("t", false, "if set, a textual graph presentation will be written to the standard output")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	g := generateGraph(*n)
	if *t {
		fmt.Println(g)
	}

	f, e := os.Create("graph.png")
	if e != nil {
		exit(insufficientRights)
	}
	gimage.DrawGraph(f, color.White, g)

	recorder := gimage.NewRecorder(g)
	sptree := kruskal.Kruskal(recorder, g)
	if *t {
		fmt.Println(sptree)
	}

	gifile, e := os.Create("kruskal_anim.gif")
	if e != nil {
		exit(insufficientRights)
	}
	recorder.Gif(gifile)

	k, e := os.Create("kruskal.png")
	if e != nil {
		exit(insufficientRights)
	}
	gimage.DrawGraph(k, color.White, sptree)
}

func generateGraph(N int) (g *kruskal.Graph) {
	// generate a random graph with N nodes
	nodes := make([]*kruskal.Node, N)
	edges := make([][]int, N)

	for i := range nodes {
		nodes[i] = &kruskal.Node{X: rand.Intn(gimage.MaxWidth), Y: rand.Intn(gimage.MaxHeight)}
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

func exit(err string) {
	panic(os.Args[0] + ":" + err)
}
