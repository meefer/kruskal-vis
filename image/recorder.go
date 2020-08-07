package image

import (
	"image"
	"image/color"
	"image/gif"
	"io"

	k "github.com/meefer/kruskal-vis/kruskal"
)

var (
	red   = color.RGBA{0xff, 0x0, 0x0, 0xff}
	green = color.RGBA{0x0, 0xff, 0x0, 0xff}
)

// Recorder renders the recorded Kruskal's algorithm execution into animated gif
type Recorder struct {
	g       *k.Graph
	currImg *image.Paletted
	imgs    []*image.Paletted
}

// NewRecorder returns a new Recorder with the given graph
func NewRecorder(g *k.Graph) *Recorder {
	currImg := DrawNodes(red, g)
	return &Recorder{
		g,
		currImg,
		[]*image.Paletted{currImg},
	}
}

// CheckEdge adds a new image to the gif with the highlighted edge that is checked for a spanning tree
func (r *Recorder) CheckEdge(from, to int) {
	r.imgs = append(r.imgs, DrawEdge(copyImg(r.currImg), red, r.g, from, to))
}

// SetEdge adds a new image to the gif with the highlighted edge that is set into spanning tree
func (r *Recorder) SetEdge(from, to int) {
	r.currImg = copyImg(r.currImg)
	r.imgs = append(r.imgs, DrawEdge(r.currImg, green, r.g, from, to))
}

// WriteGif creates an animated gif from the recorded Kruskal's algorithm execution and writes it to the given writer
func (r *Recorder) WriteGif(w io.Writer, delay int) {
	delays := make([]int, len(r.imgs))
	for i := range delays {
		delays[i] = delay
	}
	anim := gif.GIF{Delay: delays, Image: r.imgs}

	gif.EncodeAll(w, &anim)
}
