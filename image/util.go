package image

import (
	"image"
	"image/color"
	"image/color/palette"
	"math"
)

func round(f float64) int {
	return int(math.Round(f))
}

func copyImg(img *image.Paletted) (nimg *image.Paletted) {
	nimg = new(image.Paletted)
	nimg.Palette = palette.WebSafe
	sl := make([]uint8, len(img.Pix))
	copy(sl, img.Pix)
	nimg.Pix = sl
	nimg.Rect = img.Rect
	nimg.Stride = img.Stride
	return
}

// drawCircle draws a circle with a center in c on the given image
func drawCircle(img *image.Paletted, clr color.Color, c image.Point, r int) {
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

// drawLine draws a line from a to b on the given image
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
