package visual

type Direction int

const (
	UNKNOWN Direction = iota
	UP
	DOWN
	LEFT
	RIGHT
)

func (d Direction) NextPoint(pt BoardPoint) BoardPoint {
	switch d {
	case UP:
		return BoardPoint{pt.r - 1, pt.c}
	case DOWN:
		return BoardPoint{pt.r + 1, pt.c}
	case LEFT:
		return BoardPoint{pt.r, pt.c - 1}
	case RIGHT:
		return BoardPoint{pt.r, pt.c + 1}
	case UNKNOWN:
		return BoardPoint{pt.r, pt.c}
	default:
		panic("unexpected direction")
	}
}
