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
		im:   im,
		x:    x,
		y:    y,
		size: size,
	}
}

func (c Cell) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))

	screen.DrawImage(c.im, op)
}
