package main

import (
	"log"

	"github.com/Lucky112/numbermaker/internal/visual"
	"github.com/hajimehoshi/ebiten/v2"
)

// Настройки игрового поля
const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rounded Grid Example")

	game := visual.NewGame(screenWidth, screenHeight)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
