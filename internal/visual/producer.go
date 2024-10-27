package visual

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Producer struct {
	im *ebiten.Image

	x, y int
	size int
}

func NewProducer(x, y, size int, cl color.RGBA) Producer {
	im := ebiten.NewImage(size, size)
	im.Fill(cl)

	return Producer{
		im:   im,
		x:    x,
		y:    y,
		size: size,
	}
}

func (p Producer) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.x), float64(p.y))

	screen.DrawImage(p.im, op)
}
