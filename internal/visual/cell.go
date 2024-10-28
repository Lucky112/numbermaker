package visual

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cell struct {
	im      *ebiten.Image
	x, y    int
	size    int
	padding int
	color   color.RGBA
}

func NewCell(x, y, size, padding int, cl color.RGBA) Cell {
	im := ebiten.NewImage(size, size)
	im.Fill(cl)

	// make transparent padding around cell within its size
	for i := range size {
		for j := range size {
			if i < padding || i > size-padding || j < padding || j > size-padding {
				im.Set(i, j, color.Transparent)
			}
		}
	}

	return Cell{
		im:    im,
		x:     x,
		y:     y,
		size:  size,
		color: cl,
	}
}

func (c Cell) Draw(screen *ebiten.Image, color color.RGBA) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))

	if c.color != color {
		c.color = color
		c.im.Fill(color)
	}

	screen.DrawImage(c.im, op)
}

func (c Cell) In(x, y int) bool {
	if x >= c.x && x < c.x+c.size && y >= c.y && y <= c.y+c.size {
		return true
	}

	return false
}
