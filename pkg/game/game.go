package game

import (
	"fmt"
	"time"

	"github.com/danwhitford/danterm/pkg/board"
	"github.com/danwhitford/danterm/pkg/rule"
)

type Game struct {
	Board *board.Board
	Rule  rule.RuleStruct
}

type SurvivalFn func(neighbours int, alive bool) bool

func NewGame(width, height uint, rule rule.RuleStruct) Game {
	board := board.NewRandom(width, height)
	return Game{Board: board, Rule: rule}
}

func contains(a []int, find int) bool {
	for _, v := range a {
		if v == find {
			return true
		}
	}
	return false
}

func (game Game) survives(neighbours int, alive bool) bool {
	if alive && contains(game.Rule.Survive, neighbours) {
		return true
	} else if !alive && contains(game.Rule.Born, neighbours) {
		return true
	} else {
		return false
	}
}

func (game *Game) Step() {
	next := game.Board.Copy()

	for y := uint(0); y < game.Board.Height; y++ {
		for x := uint(0); x < game.Board.Width; x++ {
			neighbours := game.Board.CountNeighbours(x, y)
			state := game.Board.GetAt(x, y)

			if game.survives(neighbours, state) {
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

func (game Game) Run() {
	history := make(map[uint32]struct{})

	defer fmt.Print("\033[0m")
	for {
		_, ok := history[game.Hash()]
		if ok {
			break
		}
		history[game.Hash()] = struct{}{}

		fmt.Print(game)
		game.Step()
		time.Sleep(time.Second / 15)
	}
}
