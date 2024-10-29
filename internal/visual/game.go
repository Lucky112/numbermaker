package visual

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	screenWidth  int
	screenHeight int

	vf VisualField
}

func NewGame(screenWidth, screenHeight int) *Game {
	rows := 15
	cols := 20

	vf := NewVisualField(rows, cols)

	return &Game{
		vf:           vf,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.vf.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) Update() error {
	err := g.vf.Update()
	if err != nil {
		return err
	}

	return nil
}
