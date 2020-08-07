// +build !js,!wasm

package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"

	gimage "github.com/meefer/kruskal-vis/image"
	"github.com/meefer/kruskal-vis/kruskal"
)

const (
	usage = `kruskal-vis generates a random graph with a fixed number of nodes
and runs it through Kruskal's minimum-spanning-tree algorithm
producing three files in a working directory:
  - %s -- original graph
  - %s -- visualization of the Kruskal's algorithm 
  - %s -- minimum spanning tree of the original graph
Usage:
`
	graphFilename      = "graph.png"
	animFilename       = "kruskal_anim.gif"
	sptreeFilename     = "kruskal.png"
	insufficientRights = "don't have sufficient rights to create a file"
)

var (
	n = flag.Int("N", 10, "number of graph nodes")
	d = flag.Int("d", 20, "duration of an animation step")
	t = flag.Bool("t", false, "if set, a textual graph presentation will be written to the standard output")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, graphFilename, animFilename, sptreeFilename)
		flag.PrintDefaults()
	}
	flag.Parse()

	g := generateGraph(*n)
	if *t {
		fmt.Println(g)
	}

	f, e := os.Create(graphFilename)
	if e != nil {
		exit(insufficientRights)
	}
	gimage.DrawGraph(f, color.White, g)

	recorder := gimage.NewRecorder(g)
	sptree := kruskal.Kruskal(recorder, g)
	if *t {
		fmt.Println(sptree)
	}

	gifile, e := os.Create(animFilename)
	if e != nil {
		exit(insufficientRights)
	}
	recorder.WriteGif(gifile, *d)

	k, e := os.Create(sptreeFilename)
	if e != nil {
		exit(insufficientRights)
	}
	gimage.DrawGraph(k, color.White, sptree)
}

func exit(err string) {
	panic(os.Args[0] + ": " + err)
}
