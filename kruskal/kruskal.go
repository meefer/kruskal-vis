package kruskal

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"io"
	"math"
	"sort"

	"github.com/meefer/kruskal-vis/unionfind"
)

type edge [2]int
type byDistance struct {
	*Graph
	es []edge
}

func (g byDistance) d(i int) float64 {
	e := g.es[i]
	u, v := g.nodes[e[0]], g.nodes[e[1]]
	return math.Sqrt(float64((v.x-u.x)*(v.x-u.x) + (v.y-u.y)*(v.y-u.y)))
}
func (g byDistance) Len() int           { return len(g.es) }
func (g byDistance) Less(i, j int) bool { return g.d(i) < g.d(j) }
func (g byDistance) Swap(i, j int)      { g.es[i], g.es[j] = g.es[j], g.es[i] }

func Kruskal(w io.Writer, g *Graph) *Graph {
	N := len(g.nodes)
	uf := unionfind.NewUnionFind(N)
	uf.Connected(0, 0)
	var edges []edge
	for i, nodes := range g.edges {
		for _, j := range nodes {
			if j > i {
				edges = append(edges, edge{i, j})
			}
		}
	}
	sort.Sort(byDistance{g, edges})

	red := color.RGBA{0xff, 0x0, 0x0, 0xff}
	green := color.RGBA{0x0, 0xff, 0x0, 0xff}

	var imgs []*image.Paletted
	currImg := DrawNodes(red, g)
	imgs = append(imgs, currImg)
	copyImg := func(img *image.Paletted) *image.Paletted {
		n := new(image.Paletted)
		n.Palette = palette.WebSafe
		sl := make([]uint8, len(img.Pix))
		copy(sl, img.Pix)
		n.Pix = sl
		n.Rect = img.Rect
		n.Stride = img.Stride
		return n
	}

	a := &Graph{g.nodes, make([][]int, N)}
	for _, e := range edges {
		imgs = append(imgs, DrawEdge(copyImg(currImg), red, g, e[0], e[1]))
		if !uf.Connected(e[0], e[1]) {
			currImg = copyImg(currImg)
			imgs = append(imgs, DrawEdge(currImg, green, g, e[0], e[1]))

			a.SetEdge(e[0], e[1])
			uf.Union(e[0], e[1])
		}
	}

	delays := make([]int, len(imgs))
	for i := range delays {
		delays[i] = 10
	}
	anim := gif.GIF{Delay: delays, Image: imgs}
	gif.EncodeAll(w, &anim)

	return a
}
