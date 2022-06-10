package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"math/rand"
	"math"
	"vrouilhac/snake/timer"
	"vrouilhac/snake/input"
)

type BodyPart struct {
	Pos Coord
	Index int
	Symbol rune
}

type Snake struct {
	Head BodyPart
	Direction Coord // (0, 1) => go to bottom (-1, 0) => go to left
	BodyParts []BodyPart
}

func (snake *Snake) ChangeDirection(coord Coord) {
	if math.Abs(float64(coord.X)) + math.Abs(float64(snake.Direction.X)) > 1 {
		return
	}
	if math.Abs(float64(coord.Y)) + math.Abs(float64(snake.Direction.Y)) > 1 {
		return
	}

	snake.Direction = coord
}

type Game struct {
	snake Snake
	terrain Terrain
	apple Apple
}

func (game *Game) MoveSnake() {
	previousPos := game.snake.Head.Pos

	nextX := ((previousPos.X+game.snake.Direction.X)%game.terrain.Size.X)
	nextY := ((previousPos.Y+game.snake.Direction.Y)%game.terrain.Size.Y)

	if nextX < 0 {
		nextX = game.terrain.Size.X - 1
	}

	if nextY < 0 {
		nextY = game.terrain.Size.Y - 1
	}

	if nextX >= game.terrain.Size.X {
		nextX = 0
	}

	if nextY >= game.terrain.Size.Y {
		nextY = 0
	}

	game.snake.Head.Pos = Coord{
		X: nextX,
		Y: nextY,
	}

	game.terrain.ClearCell(previousPos)

	sortedBodyPart := game.snake.BodyParts
	sort.Slice(sortedBodyPart, func(i, j int) bool {
		return sortedBodyPart[i].Index < sortedBodyPart[j].Index;
	})

	updatedBodyParts := make([]BodyPart, 0)

	for _, bodyPart := range sortedBodyPart {
		tmp := bodyPart.Pos
		bodyPart.Pos = previousPos
		previousPos = tmp
		game.terrain.ClearCell(previousPos)
		updatedBodyParts = append(updatedBodyParts, bodyPart)
	}

	game.snake.BodyParts = updatedBodyParts
	game.terrain.SetSnakeAtPos(game.snake)
}

func (game *Game) IsEnd() bool {
	return game.terrain.SnakeBitHimself(game.snake)
}

func (game *Game) EvaluateInput(input string) {
	switch input {
		case "h":
			game.snake.ChangeDirection(Coord{-1, 0})
		case "j":
			game.snake.ChangeDirection(Coord{0, 1})
		case "k":
			game.snake.ChangeDirection(Coord{0, -1})
		case "l":
			game.snake.ChangeDirection(Coord{1, 0})
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

func (coord *Coord) IsEqual(pos Coord) bool {
	if coord.X == pos.X && coord.Y == pos.Y {
		return true
	}

	return false
}

func (game *Game) GenerateNewApple() {
	randomX := rand.Intn(game.terrain.Size.X)
	randomY := rand.Intn(game.terrain.Size.Y)
	newApple := Apple{
		Coord: Coord{randomX, randomY},
		Symbol: game.apple.Symbol,
	}

	game.apple = newApple
}

func (coord *Coord) Remove(pos Coord) {
	coord.X -= pos.X
	coord.Y -= pos.Y
}

func (game *Game) AddBodyPart() {
	lastBodyPart := game.snake.BodyParts[len(game.snake.BodyParts)-1]
	newCood := Coord{
		lastBodyPart.Pos.X,
		lastBodyPart.Pos.Y,
	}
	newCood.Remove(game.snake.Direction)
	bodyPart := BodyPart{
		Symbol: game.snake.BodyParts[0].Symbol,
		Pos: Coord{0, 0},
		Index: 1,
	}

	game.snake.BodyParts = append(game.snake.BodyParts, bodyPart)
}

func (game *Game) EvaluateSnakeGrowth() {
	if game.snake.Head.Pos.IsEqual(game.apple.Coord) {
		game.GenerateNewApple()	
		game.AddBodyPart()
	}
}

func ClearTUI() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	ClearTUI()

	terrain := NewTerrain(Coord{16, 16})
	apple := Apple{Coord{10, 3}, '•'}

	snake := Snake{
		Head: BodyPart{Coord{2, 2}, 0, '0'},
		Direction: Coord{0, 1},
		BodyParts: append(make([]BodyPart, 0), BodyPart{Coord{2, 1}, 0, 'O'}, BodyPart{Coord{2, 0}, 1, 'O'}),
	}
	terrain.SetAtPos(apple.Symbol, apple.Coord)
	terrain.SetSnakeAtPos(snake)
	terrain.Render()

	game := Game{
		snake,
		*terrain,
		apple,
	}

	rawDirection := make(chan string)

	go input.GetInput(rawDirection)

	Timer.ExecuteAtInterval(func (done func() ()) {
		ClearTUI()
		game.HandleInput(&rawDirection)
		game.MoveSnake()
		game.EvaluateSnakeGrowth()
		game.terrain.SetAtPos(game.apple.Symbol, game.apple.Coord)
		game.terrain.Render()

		if end := game.IsEnd(); end {
			fmt.Println(end)
			done()
		}
	})
}


type Apple struct {
	Coord Coord
	Symbol rune
}

// Game manager
// Snake entity -> Head and body with pos
// AppleGenerator
// Lose Condition
// Menu

type Terrain struct {
	Size Coord
	Grid [][]rune
}

func (terrain *Terrain) IsValidCoord(coord Coord) bool {
	return coord.Y < len(terrain.Grid) && coord.X < len(terrain.Grid[0])
}

func (terrain *Terrain) SetAtPos(c rune, pos Coord) {
	if terrain.IsValidCoord(pos) {
		terrain.Grid[pos.Y][pos.X] = c
	}
}

func (terrain *Terrain) SnakeBitHimself(snake Snake) bool {
	if terrain.Grid[snake.Head.Pos.Y][snake.Head.Pos.X] == snake.BodyParts[0].Symbol {
		return true
	}

	return false
}

func (terrain *Terrain) ClearCell(pos Coord) {
	if terrain.IsValidCoord(pos) {
		terrain.Grid[pos.Y][pos.X] = ' '
	}
}

func (terrain *Terrain) SetSnakeAtPos(snake Snake) {
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
			if terrain.IsValidCoord(Coord{i, j}) {
				fmt.Printf("%c", terrain.Grid[i][j])	
			}
		}
		fmt.Printf("\n")
	}
}

func NewTerrain(size Coord) *Terrain {
	terrain := Terrain{
		Size: size,
	}

	terrain.Init()

	return &terrain
}

// (0, 0) is the top left corner
type Coord struct {
	X, Y int
}

