package visual

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Consumer struct {
	im *ebiten.Image

	x, y int
	size int
}

func NewConsumer(x, y, size int, cl color.RGBA) Consumer {
	im := ebiten.NewImage(size, size)
	im.Fill(cl)

	return Consumer{
		im:   im,
		x:    x,
		y:    y,
		size: size,
	}
}

func (c Consumer) Draw(screen *ebiten.Image, cellSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))
	op.GeoM.Scale(float64(cellSize), float64(cellSize))

	screen.DrawImage(c.im, op)
}
