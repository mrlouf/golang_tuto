package main

import (
	"fmt"
	"os"

	"tutugo/mod/philosophers/types"
	"tutogo/mod/philosophers/parser"
	"tutogo/mod/philosophers/initialisation"
	"tutogo/mod/philosophers/simulation"
)

func main() {
	args := os.Args[1:]

	argLen := len(args)
	if argLen < 4 || argLen > 5 {
		fmt.Println("Usage: './philosophers <nb_philos> <time_to_eat> <time_to_sleep> opt:<nb_of_meals'")
		os.Exit(1)
	}

	banket := &types.Banket{}

	parser.Parser(args)
	initialisation.Initialiser(args, banket)
	simulation.Start(banket)

	os.Exit(0)
}