package visual

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	rows, cols int

	cells []Cell
	nodes []Node
	pipes []Pipe
}

func NewBoard(rows, cols int, cellSize int, padding int, cellColor color.RGBA) Board {
	cells := make([]Cell, 0, rows*cols)

	for i := range rows {
		for j := range cols {
			x := cellSize * j
			y := cellSize * i
			cell := NewCell(x, y, cellSize, padding, cellColor)

			cells = append(cells, cell)
		}
	}

	return Board{
		rows:  rows,
		cols:  cols,
		cells: cells,
	}
}

func (b Board) Draw(screen *ebiten.Image) {
	for _, c := range b.cells {
		c.Draw(screen)
	}

	for _, n := range b.nodes {
		n.Draw(screen, 30)
	}

	for _, p := range b.pipes {
		p.Draw(screen)
	}
}

func (b *Board) Add(n Node) {
	b.nodes = append(b.nodes, n)
}

func (b *Board) CanAdd(n Node) bool {
	return true
}
