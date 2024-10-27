package main

// import (
// 	"image"
// 	"image/color"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/inpututil"
// )

// // Sprite represents an image.
// type Sprite struct {
// 	image      *ebiten.Image
// 	alphaImage *image.Alpha
// 	x          int
// 	y          int
// 	inPath     bool
// }

// // In returns true if (x, y) is in the sprite, and false otherwise.
// func (s *Sprite) In(x, y int) bool {
// 	// Check the actual color (alpha) value at the specified position
// 	// so that the result of In becomes natural to users.
// 	//
// 	// Use alphaImage (*image.Alpha) instead of image (*ebiten.Image) here.
// 	// It is because (*ebiten.Image).At is very slow as this reads pixels from GPU,
// 	// and should be avoided whenever possible.
// 	return s.alphaImage.At(x-s.x, y-s.y).(color.Alpha).A > 0
// }

// // MoveTo moves the sprite to the position (x, y).
// func (s *Sprite) MoveTo(x, y int) {
// 	w, h := s.image.Bounds().Dx(), s.image.Bounds().Dy()

// 	s.x = x
// 	s.y = y
// 	if s.x < 0 {
// 		s.x = 0
// 	}
// 	if s.x > screenWidth-w {
// 		s.x = screenWidth - w
// 	}
// 	if s.y < 0 {
// 		s.y = 0
// 	}
// 	if s.y > screenHeight-h {
// 		s.y = screenHeight - h
// 	}
// }

// // Draw draws the sprite.
// func (s *Sprite) Draw(screen *ebiten.Image, alpha float32) {
// 	if s.inPath {
// 		// s.image.Fill(color.White)
// 	} else {
// 		// s.image.Fill(cellColor)
// 	}

// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(float64(s.x), float64(s.y))
// 	op.ColorScale.ScaleAlpha(alpha)

// 	screen.DrawImage(s.image, op)
// }

// // StrokeSource represents a input device to provide strokes.
// type StrokeSource interface {
// 	Position() (int, int)
// 	IsJustReleased() bool
// }

// // MouseStrokeSource is a StrokeSource implementation of mouse.
// type MouseStrokeSource struct{}

// func (m *MouseStrokeSource) Position() (int, int) {
// 	return ebiten.CursorPosition()
// }

// func (m *MouseStrokeSource) IsJustReleased() bool {
// 	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
// }

// // TouchStrokeSource is a StrokeSource implementation of touch.
// type TouchStrokeSource struct {
// 	ID ebiten.TouchID
// }

// func (t *TouchStrokeSource) Position() (int, int) {
// 	return ebiten.TouchPosition(t.ID)
// }

// func (t *TouchStrokeSource) IsJustReleased() bool {
// 	return inpututil.IsTouchJustReleased(t.ID)
// }

// // Stroke manages the current drag state by mouse.
// type Stroke struct {
// 	source StrokeSource

// 	// offsetX and offsetY represents a relative value from the sprite's upper-left position to the cursor position.
// 	offsetX int
// 	offsetY int

// 	// sprite represents a sprite being dragged.
// 	sprite *Sprite
// }

// func NewStroke(source StrokeSource, sprite *Sprite) *Stroke {
// 	x, y := source.Position()
// 	return &Stroke{
// 		source:  source,
// 		offsetX: x - sprite.x,
// 		offsetY: y - sprite.y,
// 		sprite:  sprite,
// 	}
// }

// func (s *Stroke) Update() {
// 	s.sprite.image.Fill(color.White)
// }

// func (s *Stroke) Sprite() *Sprite {
// 	return s.sprite
// }

// var (
// 	ebitenImage      *ebiten.Image
// 	ebitenAlphaImage *image.Alpha
// )

// func init() {
// 	ebitenImage = ebiten.NewImage(cellSize, cellSize)
// 	ebitenImage.Fill(cellColor)

// 	// Clone an image but only with alpha values.
// 	// This is used to detect a user cursor touches the image.
// 	b := ebitenImage.Bounds()
// 	ebitenAlphaImage = image.NewAlpha(b)
// 	for j := b.Min.Y; j < b.Max.Y; j++ {
// 		for i := b.Min.X; i < b.Max.X; i++ {
// 			ebitenAlphaImage.Set(i, j, cellColor)
// 		}
// 	}
// }

// func NewGame() *Game {
// 	// Initialize the sprites.
// 	sprites := []*Sprite{}
// 	for row := 0; row < rows; row++ {
// 		for col := 0; col < cols; col++ {
// 			x := col*(cellSize+cellSpacing) + cellSpacing
// 			y := row*(cellSize+cellSpacing) + cellSpacing

// 			s := &Sprite{
// 				image:      ebiten.NewImageFromImage(ebitenImage),
// 				alphaImage: ebitenAlphaImage,
// 				x:          x,
// 				y:          y,
// 			}
// 			sprites = append(sprites, s)

// 		}
// 	}

// 	// Initialize the game.
// 	return &Game{
// 		strokes: map[*Stroke]struct{}{},
// 		sprites: sprites,
// 	}
// }

// func (g *Game) spriteAt(x, y int) *Sprite {
// 	// As the sprites are ordered from back to front,
// 	// search the clicked/touched sprite in reverse order.
// 	for i := len(g.sprites) - 1; i >= 0; i-- {
// 		s := g.sprites[i]
// 		if s.In(x, y) {
// 			return s
// 		}
// 	}
// 	return nil
// }

// func (g *Game) moveSpriteToFront(sprite *Sprite) {
// 	index := -1
// 	for i, ss := range g.sprites {
// 		if ss == sprite {
// 			index = i
// 			break
// 		}
// 	}
// 	g.sprites = append(g.sprites[:index], g.sprites[index+1:]...)
// 	g.sprites = append(g.sprites, sprite)
// }

// func (g *Game) Update() error {
// 	// Обрабатываем начало нажатия
// 	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
// 		g.mousePressed = true
// 	}

// 	return nil
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return screenWidth, screenHeight
// }
