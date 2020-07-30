package main

import (
	"fmt"
	"image/color"
	"os"

	gimage "github.com/meefer/kruskal-vis/image"
	"github.com/meefer/kruskal-vis/kruskal"
)

func main() {
	g := kruskal.NewGraph(20)
	fmt.Println(g)

	f, _ := os.Create("image.png")
	gimage.DrawGraph(f, color.White, g)

	recorder := gimage.NewRecorder(g)
	sptree := kruskal.Kruskal(recorder, g)
	fmt.Println(sptree)

	gifile, _ := os.Create("kruskal_anim.gif")
	recorder.Gif(gifile)

	k, _ := os.Create("kruskal.png")
	gimage.DrawGraph(k, color.White, sptree)
}
