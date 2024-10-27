package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Настройки игрового поля
const (
	screenWidth  = 800
	screenHeight = 600
	cellSize     = 40     // Размер одной клетки
	cellSpacing  = 8      // Отступ между клетками
	cellRadius   = 8      // Радиус скругления углов
	rows, cols   = 10, 15 // Количество строк и столбцов клеток
)

var backgroundColor = color.RGBA{30, 30, 30, 255} // Цвет фона
var cellColor = color.RGBA{100, 200, 250, 255}    // Цвет клеток

// Game - структура игры
type Game struct {
	touchIDs     []ebiten.TouchID
	strokes      map[*Stroke]struct{}
	sprites      []*Sprite
	mousePressed bool
}

// Draw - метод отрисовки для игры
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor) // Заливка фона

	// rectImg := ebiten.NewImage(cellSize, cellSize)
	// rectImg.Fill(cellColor)

	// // Рисуем сетку клеток
	// for row := 0; row < rows; row++ {
	// 	for col := 0; col < cols; col++ {
	// 		// Координаты верхнего левого угла клетки
	// 		x := float64(col*(cellSize+cellSpacing) + cellSpacing)
	// 		y := float64(row*(cellSize+cellSpacing) + cellSpacing)

	// 		op := &ebiten.DrawImageOptions{}
	// 		op.GeoM.Translate(x, y) // Перемещаем прямоугольник на нужные координаты
	// 		screen.DrawImage(rectImg, op)
	// 	}
	// }

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.mousePressed = true
	}

	if g.mousePressed && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, s := range g.sprites {
			x, y := ebiten.CursorPosition()

			if s.In(x, y) {
				s.inPath = true
			}
		}
	}

	for _, s := range g.sprites {
		s.Draw(screen, 1)

		// if s.inPath {
		// 	s.Draw(screen, 0.5)
		// } else {
		// 	s.Draw(screen, 1)
		// }
	}

	// Сбрасываем флаг нажатия мыши
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.mousePressed = false
		for _, s := range g.sprites {
			s.inPath = false
		}
	}

	// Отображаем FPS для отладки
	ebitenutil.DebugPrint(screen, "FPS: "+fmt.Sprintf("%.2f", ebiten.ActualFPS()))
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rounded Grid Example")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
