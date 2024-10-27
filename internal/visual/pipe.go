package visual

import "github.com/hajimehoshi/ebiten/v2"

type Pipe struct {
	pipe []Cell
}

func (p Pipe) Draw(screen *ebiten.Image) {
	for _, c := range p.pipe {
		c.Draw(screen)
	}
}
