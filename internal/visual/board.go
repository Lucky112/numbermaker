package visual

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	rows, cols int

	nodes []Node
	pipes []Pipe
	ports []Port
}

func NewBoard(rows, cols int) Board {

	return Board{
		rows: rows,
		cols: cols,
	}
}

func (b Board) Draw(screen *ebiten.Image, cellSize int) {
	for _, n := range b.nodes {
		n.Draw(screen, cellSize)
	}

	for _, p := range b.ports {
		p.Draw(screen, cellSize)
	}

	for _, p := range b.pipes {
		p.Draw(screen, cellSize)
	}
}

func (b *Board) Add(n Node) {
	b.nodes = append(b.nodes, n)
}

func (b Board) CanAdd(n Node) bool {
	return true
}

func (b *Board) AddPort(p Port) {
	b.ports = append(b.ports, p)
}

func (b *Board) AddPipe(p Pipe) {
	b.pipes = append(b.pipes, p)
}

func (b Board) PortWithin(pt BoardPoint) (Port, bool) {
	for _, p := range b.ports {
		if p.In(pt) {
			return p, true
		}
	}

	return Port{}, false
}
