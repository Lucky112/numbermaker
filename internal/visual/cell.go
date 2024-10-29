package visual

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cell struct {
	im     *ebiten.Image
	origin BoardPoint
	color  color.RGBA
}

func NewCell(origin BoardPoint, cl color.RGBA) Cell {
	im := ebiten.NewImage(1, 1)
	im.Fill(cl)

	// make transparent padding around cell within its size
	// for i := range size {
	// 	for j := range size {
	// 		if i < padding || i > size-padding || j < padding || j > size-padding {
	// 			im.Set(i, j, color.Transparent)
	// 		}
	// 	}
	// }

	return Cell{
		im:     im,
		origin: origin,
		color:  cl,
	}
}

func (c Cell) Draw(screen *ebiten.Image, cellSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.origin.c), float64(c.origin.r))
	op.GeoM.Scale(float64(cellSize), float64(cellSize))

	screen.DrawImage(c.im, op)
}
