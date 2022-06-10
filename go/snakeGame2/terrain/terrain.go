package terrain

import (
	"fmt"

	"vrouilhac/snake/utils"
	"vrouilhac/snake/snake"
)

type Terrain struct {
	Size utils.Coord
	Grid [][]rune
}

func (terrain *Terrain) IsValidCoord(coord utils.Coord) bool {
	return coord.Y < len(terrain.Grid) && coord.X < len(terrain.Grid[0])
}

func (terrain *Terrain) SetAtPos(c rune, pos utils.Coord) {
	if terrain.IsValidCoord(pos) {
		terrain.Grid[pos.Y][pos.X] = c
	}
}

func (terrain *Terrain) SnakeBitHimself(snake snake.Snake) bool {
	if terrain.Grid[snake.Head.Pos.Y][snake.Head.Pos.X] == snake.BodyParts[0].Symbol {
		return true
	}

	return false
}

func (terrain *Terrain) ClearCell(pos utils.Coord) {
	if terrain.IsValidCoord(pos) {
		terrain.Grid[pos.Y][pos.X] = ' '
	}
}

func (terrain *Terrain) SetSnakeAtPos(snake snake.Snake) {
	terrain.SetAtPos(snake.Head.Symbol, snake.Head.Pos)

	for _, v := range snake.BodyParts {
		terrain.SetAtPos(v.Symbol, v.Pos)
	}
}

func (terrain *Terrain) Init() {
	grid := make([][]rune, 0)

	for i := 0; i < terrain.Size.Y; i++ {
		row := make([]rune, 0)

		for j := 0; j < terrain.Size.X; j++ {
			row = append(row, ' ')
		}

		grid = append(grid, row)
	}

	terrain.Grid = grid
}

func (terrain *Terrain) Render() {
	for i := 0; i < terrain.Size.Y; i++ {
		for j := 0; j < terrain.Size.X; j++ {
			if terrain.IsValidCoord(utils.Coord{X: i, Y: j}) {
				fmt.Printf("%c", terrain.Grid[i][j])	
			}
		}
		fmt.Printf("\n")
	}
}

func NewTerrain(size utils.Coord) *Terrain {
	terrain := Terrain{
		Size: size,
	}

	terrain.Init()

	return &terrain
}
