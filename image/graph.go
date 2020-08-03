package image

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"io"

	k "github.com/meefer/kruskal-vis/kruskal"
)

const (
	// MaxWidth is a maximum value of x coordinate for a graph node
	MaxWidth = 1000
	// MaxHeight is a maximum value of y coordinate for a graph node
	MaxHeight = 1000
	r         = 6
)

func point(n *k.Node) image.Point {
	return image.Pt(n.X, n.Y)
}

// DrawGraph creates a picture of graph and stores it to fs
func DrawGraph(w io.Writer, c color.Color, g *k.Graph) {
	img := image.NewPaletted(image.Rect(0, 0, MaxWidth, MaxHeight), palette.WebSafe)

	for i, nodes := range g.Edges {
		fromp := point(g.Nodes[i])
		drawCircle(img, c, fromp, r)
		for _, j := range nodes {
			if j > i {
				top := point(g.Nodes[j])
				drawLine(img, c, fromp, top)
			}
		}
	}

	png.Encode(w, img)
}

// DrawNodes creates a picture of graph and stores it to fs
func DrawNodes(c color.Color, g *k.Graph) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, MaxWidth, MaxHeight), palette.WebSafe)

	for i := range g.Edges {
		fromp := point(g.Nodes[i])
		drawCircle(img, c, fromp, r)
	}

	return img
}

// DrawEdge creates a picture of graph and stores it to fs
func DrawEdge(img *image.Paletted, c color.Color, g *k.Graph, u, v int) *image.Paletted {
	drawLine(img, c, point(g.Nodes[u]), point(g.Nodes[v]))

	return img
}
