package main

import (
	"fmt"
	"os"
	"os/exec"

	"vrouilhac/snake/timer"
	"vrouilhac/snake/utils"
	g "vrouilhac/snake/game"
	"vrouilhac/snake/tools"
	"vrouilhac/snake/terrain"
	"vrouilhac/snake/snake"
	"vrouilhac/snake/input"
)

func ClearTUI() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	ClearTUI()

	// TODO: Settings for size
	// TODO: Increase number of apple ( and add to settings )
	terrain := terrain.NewTerrain(utils.Coord{X: 30, Y: 30})
	apple := tools.Apple{
		Coord: utils.Coord{X: 10, Y: 3}, 
		Symbol: '•',
	}

	snake := snake.Snake{
		Head: snake.BodyPart{Pos: utils.Coord{X: 2, Y: 2}, Index: 0, Symbol:'0'},
		Direction: utils.Coord{X: 0, Y: 1},
		BodyParts: append(
			make([]snake.BodyPart, 0), 
			snake.BodyPart{ Pos: utils.Coord{X: 2, Y: 1}, Index: 0, Symbol: 'O'}, 
			snake.BodyPart{ Pos: utils.Coord{X: 2, Y: 0}, Index: 1, Symbol: 'O'},
		),
	}
	terrain.SetAtPos(apple.Symbol, apple.Coord)
	terrain.SetSnakeAtPos(snake)
	terrain.Render()

	game := g.Game{
		Snake: snake,
		Terrain: *terrain,
		Apple: apple,
	}

	rawDirection := make(chan string)

	go input.GetInput(rawDirection)

	Timer.ExecuteAtInterval(func (done func() ()) {
		ClearTUI()
		game.HandleInput(&rawDirection)
		game.MoveSnake()
		game.EvaluateSnakeGrowth()
		game.Terrain.SetAtPos(game.Apple.Symbol, game.Apple.Coord)
		game.Terrain.Render()

		if end := game.IsEnd(); end {
			fmt.Println(end)
			done()
			os.Exit(0)
		}
	})
}



// Game manager
// Snake entity -> Head and body with pos
// AppleGenerator
// Lose Condition
// Menu

