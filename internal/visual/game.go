package visual

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	background   color.RGBA
	board        Board
	screenWidth  int
	screenHeight int

	mousePressed   bool
	pipeInProgress bool
}

func NewGame(screenWidth, screenHeight int) *Game {
	rows := 15
	cols := 20
	padding := 2
	cellSize := 30
	cellColor := color.RGBA{200, 200, 0, 255}
	backColor := color.RGBA{0, 0, 0, 255}

	board := NewBoard(rows, cols, cellSize, padding, cellColor)

	for range 2 {
		nSize := 2
		x := rand.Intn(cols - nSize)
		y := rand.Intn(rows - nSize)

		n := NewConsumer(x, y, nSize, color.RGBA{255, 0, 0, 255})
		board.Add(n)

		p := NewPort(x, y, 1, color.RGBA{255, 255, 0, 255})
		board.AddPort(p)
	}

	return &Game{
		background:   backColor,
		board:        board,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.background)

	g.board.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.mousePressed = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.mousePressed = false

		x, y := ebiten.CursorPosition()
		port := g.board.PortWithin(x, y)
		if port.im != nil { // TODO : better check on port existance
			g.pipeInProgress = false
			g.board.PersistPipe()
		} else {
			g.board.CancelPipe()
		}

	}

	if g.mousePressed {
		x, y := ebiten.CursorPosition()

		if !g.pipeInProgress {
			port := g.board.PortWithin(x, y)
			if port.im != nil { // TODO : better check on port existance
				g.pipeInProgress = true
			}
		} else {

			cell := g.board.CellWithin(x, y)
			g.board.ElongatePipe(cell)
		}

	}

	// k := rand.Int31n(10)

	// if k == 5 {
	// 	x := rand.Intn(g.board.cols - 3)
	// 	y := rand.Intn(g.board.rows - 3)

	// 	n := NewConsumer(x, y, 3, color.RGBA{255, 0, 0, 255})
	// 	g.board.Add(n)
	// }

	return nil
}
