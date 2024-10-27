package visual

import "github.com/hajimehoshi/ebiten/v2"

type Node interface {
	Draw(*ebiten.Image, int)
}
