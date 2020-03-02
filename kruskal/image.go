package kruskal

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"io"
	"math"
)

const (
	width  = maxNodeXY
	height = maxNodeXY
	r      = 6
)

// DrawGraph creates a picture of graph and stores it to fs
func DrawGraph(w io.Writer, c color.Color, g *Graph) {
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette.WebSafe)

	for i, nodes := range g.edges {
		fromp := point(g.nodes[i])
		drawCircle(img, c, fromp)
		for _, j := range nodes {
			if j > i {
				top := point(g.nodes[j])
				drawLine(img, c, fromp, top)
			}
		}
	}

	png.Encode(w, img)
}

// DrawNodes creates a picture of graph and stores it to fs
func DrawNodes(c color.Color, g *Graph) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette.WebSafe)

	for i := range g.edges {
		fromp := point(g.nodes[i])
		drawCircle(img, c, fromp)
	}

	return img
}

// DrawEdge creates a picture of graph and stores it to fs
func DrawEdge(img *image.Paletted, c color.Color, g *Graph, u, v int) *image.Paletted {
	drawLine(img, c, point(g.nodes[u]), point(g.nodes[v]))

	return img
}

func point(n *Node) image.Point {
	return image.Pt(n.x, n.y)
}

func round(f float64) int {
	return int(math.Round(f))
}

// drawCircle draws circle with a center in c on the provided image
func drawCircle(img *image.Paletted, clr color.Color, c image.Point) {
	for x := c.X - r; x <= c.X+r; x++ {
		y := round(math.Sqrt(float64(r*r - (x-c.X)*(x-c.X))))
		img.Set(x, c.Y+y, clr)
		img.Set(x, c.Y-y, clr)
	}
	for y := c.Y - r; y <= c.Y+r; y++ {
		x := round(math.Sqrt(float64(r*r - (y-c.Y)*(y-c.Y))))
		img.Set(c.X+x, y, clr)
		img.Set(c.X-x, y, clr)
	}
}

// drawLine draws line from a to b on the provided image
func drawLine(img *image.Paletted, clr color.Color, a, b image.Point) {
	dx := b.X - a.X
	if dx == 0 {
		for y := a.Y; y <= b.Y; y++ {
			img.Set(a.X, y, clr)
		}
		return
	}
	dy := b.Y - a.Y
	if dy == 0 {
		for x := a.X; x <= b.X; x++ {
			img.Set(x, a.Y, clr)
		}
		return
	}

	m := float64(dy) / float64(dx)
	c := float64(b.Y) - m*float64(b.X)

	if a.X > b.X {
		a, b = b, a
	}
	for x := a.X; x <= b.X; x++ {
		y := round(m*float64(x) + c)
		img.Set(x, y, clr)
	}

	if a.Y > b.Y {
		a, b = b, a
	}
	for y := a.Y; y <= b.Y; y++ {
		x := round((float64(y) - c) / m)
		img.Set(x, y, clr)
	}
}
