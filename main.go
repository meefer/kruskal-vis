package main

import (
	"fmt"

	"github.com/meefer/it/kruskal"
)

func main() {
	g := kruskal.NewGraph(20)
	fmt.Println(g)
	kruskal.DrawGraph(g, "image.png")
}
