package main

import (
	"os"
	"os/exec"
	"fmt"

	tui "vrouilhac/snake/renderer"
	"vrouilhac/snake/utils"
)


type Game struct {
	renderer tui.Renderer
	snakeDirection utils.Vector
	applePos utils.Vector
}

func (game *Game) Init() {
	game.renderer.Build()
	applePos := game.renderer.GenerateApple()
	game.applePos = applePos
	game.renderer.GenerateSnakePos()
}

func (game *Game) Play() {
	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		game.renderer.Render()
		game.PlayMove()
	}
}

func (game *Game) PlayMove() {
	end := false

	snake := game.renderer.Snake

	for !end {
		game.renderer.Grid[snake.Position.X][snake.Position.Y] = " "

		nextX, nextY := 0, 0

		if snake.Position.X + snake.Direction.X < game.renderer.Size {
			nextX = snake.Position.X + snake.Direction.X
		} else {
			nextX = 0
		}

		if snake.Position.Y + snake.Direction.Y < game.renderer.Size {
			nextY = snake.Position.Y + snake.Direction.Y
		} else {
			nextY = 0
		}

		game.renderer.Grid[nextX][nextY] = "■"
		fmt.Printf("Snake: (%v, %v)\tDirection: (%v, %v)", nextX, nextY, snake.Direction.X, snake.Direction.Y)

		if snake.Next == nil {
			end = true
		}

		snake = snake.Next
	}
}

func main() {
	renderer := tui.Renderer{Size: 16, Grid: make([][]string, 0), SnakeLength: 2}
	snakeDirection := utils.Vector{X: 0, Y: 1}
	applePos := utils.Vector{X: 0, Y: 0}
	game := Game{renderer, snakeDirection, applePos}

	game.Init()
	game.Play()
}
