package visual

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Port struct {
	im *ebiten.Image

	origin    BoardPoint
	direction Direction
}

func NewPort(origin BoardPoint, direction Direction, cl color.RGBA) Port {
	var im *ebiten.Image

	base := ebiten.NewImage(5, 5)
	base.Fill(cl)

	base.Set(0, 1, color.Transparent)
	base.Set(0, 2, color.Transparent)
	base.Set(1, 2, color.Transparent)
	base.Set(3, 2, color.Transparent)
	base.Set(4, 2, color.Transparent)
	base.Set(4, 1, color.Transparent)
	base.Set(0, 3, color.Transparent)
	base.Set(1, 3, color.Transparent)
	base.Set(2, 3, color.Transparent)
	base.Set(3, 3, color.Transparent)
	base.Set(4, 3, color.Transparent)
	base.Set(0, 4, color.Transparent)
	base.Set(1, 4, color.Transparent)
	base.Set(2, 4, color.Transparent)
	base.Set(3, 4, color.Transparent)
	base.Set(4, 4, color.Transparent)

	switch direction {
	case DOWN:
		im = base
	case UP:
		im = ebiten.NewImage(5, 5)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(base.Bounds().Dx())/2, -float64(base.Bounds().Dy())/2)
		op.GeoM.Rotate(math.Pi)
		op.GeoM.Translate(float64(base.Bounds().Dx())/2, float64(base.Bounds().Dy())/2)
		im.DrawImage(base, op)
	case RIGHT:
		im = ebiten.NewImage(5, 5)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(base.Bounds().Dx())/2, -float64(base.Bounds().Dy())/2)
		op.GeoM.Rotate(-math.Pi / 2)
		op.GeoM.Translate(float64(base.Bounds().Dx())/2, float64(base.Bounds().Dy())/2)
		im.DrawImage(base, op)
	case LEFT:
		im = ebiten.NewImage(5, 5)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(base.Bounds().Dx())/2, -float64(base.Bounds().Dy())/2)
		op.GeoM.Rotate(math.Pi / 2)
		op.GeoM.Translate(float64(base.Bounds().Dx())/2, float64(base.Bounds().Dy())/2)
		im.DrawImage(base, op)
	}

	return Port{
		im:        im,
		origin:    origin,
		direction: direction,
	}
}

func (p Port) Draw(screen *ebiten.Image, cellSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(cellSize)/float64(p.im.Bounds().Dx()), float64(cellSize)/float64(p.im.Bounds().Dy()))
	op.GeoM.Translate(float64(p.origin.c)*float64(cellSize), float64(p.origin.r)*float64(cellSize))

	switch p.direction {
	case DOWN:
		op.GeoM.Translate(0, float64(cellSize))
	case UP:
		op.GeoM.Translate(0, float64(-cellSize))
	case LEFT:
		op.GeoM.Translate(float64(-cellSize), 0)
	case RIGHT:
		op.GeoM.Translate(float64(cellSize), 0)
	}

	screen.DrawImage(p.im, op)
}

func (p Port) In(pt BoardPoint) bool {
	return pt == p.origin || pt == p.direction.NextPoint(p.origin)
}
