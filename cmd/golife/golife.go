package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/danwhitford/golife/pkg/game"
	"github.com/danwhitford/golife/pkg/rule"
	"golang.org/x/term"
)

func main() {

	width, height, _ := term.GetSize(int(os.Stdin.Fd()))

	automata := flag.String("automata", "conway", "The automata to use")
	rulestring := flag.String("rulestring", "", "Define a B/S rulestring to use, see https://conwaylife.com/wiki/Rulestring")
	listAutomata := flag.Bool("list", false, "List all automata")

	flag.Parse()

	if *listAutomata {
		for k, r := range rule.Patterns {
			fmt.Printf("%s (%s)\n", k, r)
		}
		os.Exit(0)
	}

	var r rule.RuleStruct
	if *rulestring != "" {
		var err error
		r, err = rule.ParseRule(*rulestring)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse the rulestring '%s'\n", *rulestring)
			os.Exit(1)
		}
	} else {
		var ok bool
		r, ok = rule.Patterns[*automata]
		if !ok {
			fmt.Fprintf(os.Stderr, "Could not find automata '%s'. Try '-list' to list all available.\n", *automata)
			os.Exit(1)
		}
	}

	game := game.NewGame(uint(width), uint(height), r)
	game.Run()
}
