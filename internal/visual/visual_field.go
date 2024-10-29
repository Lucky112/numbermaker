package visual

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/exp/rand"
)

type VisualField struct {
	backColor  color.RGBA
	background []Cell
	cellSize   float64

	board       Board
	pipeBuilder PipeBuilder

	mousePressed bool
}

func NewVisualField(rows, cols int) VisualField {
	cellSize := 30
	cellColor := color.RGBA{255, 255, 0, 255}
	pipeColor := color.RGBA{255, 0, 255, 255}
	backColor := color.RGBA{0, 0, 0, 255}

	board := NewBoard(rows, cols)
	for range 2 {
		nSize := 2
		x := rand.Intn(cols - nSize)
		y := rand.Intn(rows - nSize)

		n := NewConsumer(x, y, nSize, color.RGBA{255, 0, 0, 255})
		board.Add(n)

		var p Port
		p = NewPort(BoardPoint{y, x}, UP, color.RGBA{0, 255, 0, 255})
		board.AddPort(p)
		p = NewPort(BoardPoint{y, x}, DOWN, color.RGBA{0, 255, 0, 255})
		board.AddPort(p)
		p = NewPort(BoardPoint{y, x}, LEFT, color.RGBA{0, 255, 0, 255})
		board.AddPort(p)
		p = NewPort(BoardPoint{y, x}, RIGHT, color.RGBA{0, 255, 0, 255})
		board.AddPort(p)
	}

	bg := makeBackground(rows, cols, cellSize, cellColor)
	return VisualField{
		backColor:   backColor,
		background:  bg,
		cellSize:    float64(cellSize),
		board:       board,
		pipeBuilder: NewPipeBuilder(pipeColor),
	}
}

func (vf VisualField) Draw(screen *ebiten.Image) {
	screen.Fill(vf.backColor)

	for _, c := range vf.background {
		c.Draw(screen, int(vf.cellSize))
	}

	vf.board.Draw(screen, int(vf.cellSize))
	vf.pipeBuilder.Draw(screen, int(vf.cellSize))
}

func (vf *VisualField) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		vf.mousePressed = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		vf.mousePressed = false

		x, y := ebiten.CursorPosition()

		pt := BoardPoint{y / int(vf.cellSize), x / int(vf.cellSize)}
		_, ok := vf.board.PortWithin(pt)
		if ok {
			vf.board.AddPipe(*vf.pipeBuilder.pipe)
		}

		vf.pipeBuilder.CancelPipe()
	}

	if vf.mousePressed {
		x, y := ebiten.CursorPosition()
		pt := BoardPoint{y / int(vf.cellSize), x / int(vf.cellSize)}

		if vf.pipeBuilder.pipe == nil {
			port, ok := vf.board.PortWithin(pt)
			if ok {
				vf.pipeBuilder.InitPipe(port.direction)
			}
		} else {
			vf.pipeBuilder.ElongatePipe(pt)
		}
	}

	return nil
}

func makeBackground(rows, cols int, cellSize int, cellColor color.RGBA) []Cell {
	cells := make([]Cell, 0, rows*cols)
	// padding := 1

	for r := range rows {
		for c := range cols {
			cell := NewCell(BoardPoint{r, c}, cellColor)

			cells = append(cells, cell)
		}
	}

	return cells
}
