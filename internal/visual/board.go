package visual

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	rows, cols int
	cellColor  color.RGBA

	cells []Cell
	nodes []Node
	pipes []Pipe
	ports []Port

	currentPipe *Pipe
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
		rows:      rows,
		cols:      cols,
		cells:     cells,
		cellColor: cellColor,
	}
}

func (b Board) Draw(screen *ebiten.Image) {
	for _, c := range b.cells {
		c.Draw(screen, b.cellColor)
	}

	for _, n := range b.nodes {
		n.Draw(screen, 30)
	}

	for _, p := range b.ports {
		p.Draw(screen, 30)
	}

	for _, p := range b.pipes {
		p.Draw(screen)
	}

	if b.currentPipe != nil {
		b.currentPipe.Draw(screen)
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

func (b Board) CellWithin(x, y int) Cell {
	for _, cell := range b.cells {
		if cell.In(x, y) {
			return cell
		}
	}

	return Cell{} // TODO : is it impossible ?
}

func (b Board) PortWithin(x, y int) Port {
	for _, p := range b.ports {
		if p.In(x/30, y/30) {
			return p
		}
	}

	return Port{} // TODO : is it impossible ?
}

func (b *Board) ElongatePipe(cell Cell) {
	if b.currentPipe == nil {
		b.currentPipe = &Pipe{}
	}

	b.currentPipe.pipe = append(b.currentPipe.pipe, NewCell(cell.x, cell.y, cell.size, cell.padding, cell.color))
}

func (b *Board) PersistPipe() {
	b.pipes = append(b.pipes, *b.currentPipe)
	b.currentPipe = nil
}

func (b *Board) CancelPipe() {
	b.currentPipe = nil
}
