package snake

import (
	"math"

	"vrouilhac/snake/utils"
)

type BodyPart struct {
	Pos utils.Coord
	Index int
	Symbol rune
}

type Snake struct {
	Head BodyPart
	Direction utils.Coord // (0, 1) => go to bottom (-1, 0) => go to left
	BodyParts []BodyPart
}

func (snake *Snake) ChangeDirection(coord utils.Coord) {
	if math.Abs(float64(coord.X)) + math.Abs(float64(snake.Direction.X)) > 1 {
		return
	}
	if math.Abs(float64(coord.Y)) + math.Abs(float64(snake.Direction.Y)) > 1 {
		return
	}

	snake.Direction = coord
}
