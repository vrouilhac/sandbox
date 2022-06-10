package game

import (
	"sort"
	"math/rand"

	"vrouilhac/snake/snake"
	"vrouilhac/snake/terrain"
	"vrouilhac/snake/utils"
	"vrouilhac/snake/tools"
)

type Game struct {
	Snake snake.Snake
	Terrain terrain.Terrain
	Apple tools.Apple
}

func (game *Game) MoveSnake() {
	previousPos := game.Snake.Head.Pos

	nextX := ((previousPos.X+game.Snake.Direction.X)%game.Terrain.Size.X)
	nextY := ((previousPos.Y+game.Snake.Direction.Y)%game.Terrain.Size.Y)

	if nextX < 0 {
		nextX = game.Terrain.Size.X - 1
	}

	if nextY < 0 {
		nextY = game.Terrain.Size.Y - 1
	}

	if nextX >= game.Terrain.Size.X {
		nextX = 0
	}

	if nextY >= game.Terrain.Size.Y {
		nextY = 0
	}

	game.Snake.Head.Pos = utils.Coord{
		X: nextX,
		Y: nextY,
	}

	game.Terrain.ClearCell(previousPos)

	sortedBodyPart := game.Snake.BodyParts
	sort.Slice(sortedBodyPart, func(i, j int) bool {
		return sortedBodyPart[i].Index < sortedBodyPart[j].Index;
	})

	updatedBodyParts := make([]snake.BodyPart, 0)

	for _, bodyPart := range sortedBodyPart {
		tmp := bodyPart.Pos
		bodyPart.Pos = previousPos
		previousPos = tmp
		game.Terrain.ClearCell(previousPos)
		updatedBodyParts = append(updatedBodyParts, bodyPart)
	}

	game.Snake.BodyParts = updatedBodyParts
	game.Terrain.SetSnakeAtPos(game.Snake)
}

func (game *Game) IsEnd() bool {
	return game.Terrain.SnakeBitHimself(game.Snake)
}

func (game *Game) EvaluateInput(input string) {
	switch input {
	case "h":
		game.Snake.ChangeDirection(utils.Coord{X: -1, Y: 0})
	case "j":
		game.Snake.ChangeDirection(utils.Coord{X: 0, Y:1})
	case "k":
		game.Snake.ChangeDirection(utils.Coord{X: 0, Y: -1})
	case "l":
		game.Snake.ChangeDirection(utils.Coord{X: 1, Y: 0})
	default:
		return

	}
}

func (game *Game) HandleInput(c *chan string) {
	select {
	case stdin, _ := <-*c:
		game.EvaluateInput(stdin)
	default:
		return
	}
}

// TODO: this function can be better
func (game *Game) GenerateNewApple() {
	validGeneration := false

	newApple := tools.Apple{}

	randomX := rand.Intn(game.Terrain.Size.X)
	randomY := rand.Intn(game.Terrain.Size.Y)
	newApple = tools.Apple{
		Coord: utils.Coord{X: randomX, Y:randomY},
		Symbol: game.Apple.Symbol,
	}

	if game.Terrain.Grid[randomY][randomX] != game.Snake.Head.Symbol &&
	game.Terrain.Grid[randomY][randomX] != game.Snake.BodyParts[0].Symbol {
		validGeneration = true
	}

	if !validGeneration {
		randomX = rand.Intn(game.Terrain.Size.X)
		randomY = rand.Intn(game.Terrain.Size.Y)
		newApple = tools.Apple{
			Coord: utils.Coord{X: randomX, Y: randomY},
			Symbol: game.Apple.Symbol,
		}
	}

	game.Apple = newApple
}

func (game *Game) AddBodyPart() {
	lastBodyPart := game.Snake.BodyParts[len(game.Snake.BodyParts)-1]
	newCood := utils.Coord{
		X: lastBodyPart.Pos.X,
		Y: lastBodyPart.Pos.Y,
	}
	newCood.Remove(game.Snake.Direction)
	bodyPart := snake.BodyPart{
		Symbol: game.Snake.BodyParts[0].Symbol,
		Pos: utils.Coord{X: 0, Y: 0},
		Index: 1,
	}

	game.Snake.BodyParts = append(game.Snake.BodyParts, bodyPart)
}

func (game *Game) EvaluateSnakeGrowth() {
	if game.Snake.Head.Pos.IsEqual(game.Apple.Coord) {
		game.GenerateNewApple()	
		game.AddBodyPart()
	}
}
