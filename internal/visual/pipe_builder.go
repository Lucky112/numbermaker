package visual

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type PipeBuilder struct {
	pipe      *Pipe
	color     color.RGBA
	direction Direction
	pipeEnd   *Cell
}

func NewPipeBuilder(pipeColor color.RGBA) PipeBuilder {
	return PipeBuilder{
		color: pipeColor,
	}
}

func (pb *PipeBuilder) Draw(screen *ebiten.Image, cellSize int) {
	if pb.pipe != nil {
		pb.pipe.Draw(screen, cellSize)
	}
}

func (pb *PipeBuilder) ElongatePipe(pt BoardPoint) {
	if pb.pipeEnd != nil {
		dr := pt.r - pb.pipeEnd.origin.r
		dc := pt.c - pb.pipeEnd.origin.c

		if math.Abs(float64(dr)) > 1 || math.Abs(float64(dc)) > 1 || math.Abs(float64(dr)) == 1 && math.Abs(float64(dc)) == 1 {
			return
		}
	}

	cell := NewCell(pt, pb.color)
	pb.pipeEnd = &cell

	pb.pipe.pipe = append(pb.pipe.pipe, cell)
}

func (pb *PipeBuilder) InitPipe(dir Direction) {
	pb.pipe = &Pipe{}
	pb.direction = dir
}

func (pb *PipeBuilder) CancelPipe() {
	pb.pipe = nil
	pb.pipeEnd = nil
}
