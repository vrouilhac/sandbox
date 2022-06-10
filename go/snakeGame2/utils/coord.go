package utils

// (0, 0) is the top left corner
type Coord struct {
	X, Y int
}

func (coord *Coord) IsEqual(pos Coord) bool {
	if coord.X == pos.X && coord.Y == pos.Y {
		return true
	}

	return false
}

func (coord *Coord) Remove(pos Coord) {
	coord.X -= pos.X
	coord.Y -= pos.Y
}
