package renderer

import (
	"fmt"
	"math/rand"

	"vrouilhac/snake/utils"
)

type SnakeBody struct {
	Direction utils.Vector
	Position utils.Vector
	Next *SnakeBody
	Previous *SnakeBody
}

type Renderer struct {
 Size int
 Grid [][]string
 SnakeLength int
 Snake *SnakeBody
}

func (renderer *Renderer) Build() {
	grid := make([][]string, 0)

	for i := 0; i < renderer.Size; i++ {
		row := make([]string, 0)
		for j := 0; j < renderer.Size; j++ {
			row = append(row, " ")
		}

		grid = append(grid, row)
	}

	renderer.Grid = grid
}

func (renderer *Renderer) Render() {
	for i := 0; i < renderer.Size; i++ {
		for j := 0; j < renderer.Size; j++ {
			fmt.Printf(renderer.Grid[i][j])
		}
		fmt.Printf("\n")
	}
}

func (renderer *Renderer) GenerateApple() utils.Vector {
	posX := rand.Intn(renderer.Size)
	posY := rand.Intn(renderer.Size)
	renderer.Grid[posX][posY] = "•"
	return utils.Vector{
		X: posX,
		Y: posY,
	}
}

func (renderer *Renderer) GenerateSnakePos() {
	posX := rand.Intn(renderer.Size)
	posY := rand.Intn(renderer.Size)

	snakeHead := SnakeBody{
		Position: utils.Vector{
			X: posX,
			Y: posY,
		},
		Direction: utils.Vector{
			X: 0,
			Y: 1,
		},
	}

	bodyPosX := 0
	bodyPosY := 0

	if posX + 1 < renderer.Size {
		bodyPosX = posX + 1
	} else {
		bodyPosX = posX - 1
	}

	bodyPosY = posY

	snakeBody := SnakeBody{
		Position: utils.Vector{
			X: bodyPosX,
			Y: bodyPosY,
		},
		Direction: utils.Vector{
			X: 0,
			Y: 1,
		},
		Previous: &snakeHead,
	}

	snakeHead.Next = &snakeBody

	renderer.Snake = &snakeHead

	renderer.AddSnakeToGrid(snakeHead)
}

func (renderer *Renderer) AddSnakeToGrid(snake SnakeBody) {
	end := false
	snakeBody := &snake
	for !end {
		renderer.Grid[snakeBody.Position.X][snakeBody.Position.Y] = "■"

		fmt.Println(snakeBody)

		if snakeBody.Next == nil {
			end = true
		}

		snakeBody = snakeBody.Next
	}
}
