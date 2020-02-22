package kruskal

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	width  = 1000
	height = 1000
	e      = 1
	r      = 6
)

var (
	red = color.RGBA{255, 0, 0, 0xff}
)

func point(n *Node) image.Point {
	return image.Point{
		round(n.x * width),
		round(n.y * height),
	}
}

// DrawGraph creates a picture of graph and stores it to fs
func DrawGraph(g *Graph, path string) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for i, nodes := range g.edges {
		fromp := point(g.nodes[i])
		fmt.Println(fromp)
		DrawCircle(img, fromp)
		for _, j := range nodes {
			top := point(g.nodes[j])
			DrawLine(img, fromp, top)
		}
	}

	// DrawCircle(img, image.Point{47, 68})
	// DrawLine(img, image.Point{13, 17}, image.Point{47, 68})
	// DrawLine(img, image.Point{975, 79}, image.Point{66, 157})

	f, _ := os.Create(path)
	png.Encode(f, img)
}

func round(f float64) int {
	return int(math.Round(f))
}

// DrawCircle draws circle with a center in c on the provided image
func DrawCircle(img *image.RGBA, p image.Point) {
	for x := p.X - r; x <= p.X+r; x += e {
		y := round(math.Sqrt(float64(r*r - (x-p.X)*(x-p.X))))
		img.Set(x, p.Y+y, red)
		img.Set(x, p.Y-y, red)
	}
	for y := p.Y - r; y <= p.Y+r; y += e {
		x := round(math.Sqrt(float64(r*r - (y-p.Y)*(y-p.Y))))
		img.Set(p.X+x, y, red)
		img.Set(p.X-x, y, red)
	}
}

// DrawLine draws line from a to b on the provided image
func DrawLine(img *image.RGBA, a, b image.Point) {
	m := float64(b.Y-a.Y) / float64(b.X-a.X) // TODO: handle b.X = a.X == 0
	c := float64(b.Y) - m*float64(b.X)

	if a.X > b.X {
		a, b = b, a
	}
	for x := a.X; x <= b.X; x += e {
		y := round(m*float64(x) + c)
		img.Set(x, y, red)
	}

	if a.Y > b.Y {
		a, b = b, a
	}
	for y := a.Y; y <= b.Y; y += e {
		x := round((float64(y) - c) / m)
		img.Set(x, y, red)
	}
}
