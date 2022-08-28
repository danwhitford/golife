package game

import (
	"fmt"
	"time"

	"github.com/danwhitford/danterm/pkg/board"
)

type Game struct {
	Board *board.Board
	Survives func(x, y uint) bool
}

type SurvivalFn func (neighbours int, alive bool) bool

func NewGame(width, height uint) Game {
	board := board.NewRandom(width, height)
	return Game{Board: board}
}

func (game *Game) Step(survives SurvivalFn) {
	next := game.Board.Copy()

	for y := uint(0); y < game.Board.Height; y++ {
		for x := uint(0); x < game.Board.Width; x++ {
			neighbours := game.Board.CountNeighbours(x, y)
			state := game.Board.GetAt(x, y)

			if survives(neighbours, state) {
				next.SetAt(x, y)
			} else {
				next.SetAtTo(x, y, false)
			}
		}
	}

	game.Board = next
}

func (game Game) Hash() uint32 {
	return game.Board.Hash()
}

func (game Game) String() string {
	return game.Board.String()
}

func (game Game) Run(survives SurvivalFn) {
	history := make(map[uint32]struct{})

	defer fmt.Print("\033[0m")
	for {
		_, ok := history[game.Hash()]
		if ok {
			break
		}
		history[game.Hash()] = struct{}{}

		fmt.Print(game)
		game.Step(survives)
		time.Sleep(time.Second / 15)
	}
}