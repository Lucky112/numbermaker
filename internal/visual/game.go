package visual

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	background   color.RGBA
	board        Board
	screenWidth  int
	screenHeight int
}

func NewGame(screenWidth, screenHeight int) *Game {
	rows := 15
	cols := 20
	padding := 2
	cellSize := 30
	cellColor := color.RGBA{200, 200, 0, 255}
	backColor := color.RGBA{0, 0, 0, 255}

	board := NewBoard(rows, cols, cellSize, padding, cellColor)

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
	k := rand.Int31n(10)

	if k == 5 {
		x := rand.Intn(g.board.cols - 3)
		y := rand.Intn(g.board.rows - 3)

		n := NewConsumer(x, y, 3, color.RGBA{255, 0, 0, 255})
		g.board.Add(n)
	}

	return nil
}
