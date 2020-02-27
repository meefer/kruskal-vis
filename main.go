package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/meefer/it/kruskal"
)

func main() {
	g := kruskal.NewGraph(20)
	fmt.Println(g)

	f, _ := os.Create("image.png")
	kruskal.DrawGraph(f, color.White, g)

	anim, _ := os.Create("kruskal_anim.gif")
	a := kruskal.Kruskal(anim, g)
	fmt.Println(a)
	k, _ := os.Create("kruskal.png")
	kruskal.DrawGraph(k, color.White, a)
}
