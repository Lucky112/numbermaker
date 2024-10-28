package visual

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Port struct {
	im *ebiten.Image

	x, y int
	size int
}

func NewPort(x, y, size int, cl color.RGBA) Port {
	im := ebiten.NewImage(size, size)
	im.Fill(cl)

	return Port{
		im:   im,
		x:    x,
		y:    y,
		size: size,
	}
}

func (p Port) Draw(screen *ebiten.Image, cellSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.x), float64(p.y))
	op.GeoM.Scale(float64(cellSize), float64(cellSize))

	screen.DrawImage(p.im, op)
}

func (p Port) In(x, y int) bool {
	if x >= p.x && x < p.x+p.size && y >= p.y && y <= p.y+p.size {
		return true
	}

	return false
}
