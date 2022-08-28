package main

import (
	"os"

	"github.com/danwhitford/danterm/pkg/game"
	"github.com/danwhitford/danterm/pkg/steppers"
	"golang.org/x/term"
)

func main() {
	width, height, _ := term.GetSize(int(os.Stdin.Fd()))
	game := game.NewGame(uint(width), uint(height))
	game.Run(steppers.ConwayStep)
}
