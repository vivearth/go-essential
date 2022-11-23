package challenge

import (
	"fmt"
)

type Square struct {
	X    int
	Y    int
	side int
}

func NewSquare(X int, Y int, side int) (*Square, error) {
	if side < 0 {
		return nil, fmt.Errorf("Cannot create a square of negative lenth")
	}
	sq := Square{X, Y, side}
	return &sq, nil
}

func (sq Square) Area() int {
	return sq.side * sq.side
}

func (sq *Square) Move(dx int, dy int) {
	sq.X += dx
	sq.Y += dy
}
