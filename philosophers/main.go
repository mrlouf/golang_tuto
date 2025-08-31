package main

import (
	"fmt"
	"os"

	"tutogo/mod/philosophers/types"
	"tutogo/mod/philosophers/parsing"
	"tutogo/mod/philosophers/initialisation"
	"tutogo/mod/philosophers/simulation"
)

func main() {
	args := os.Args[1:]

	argLen := len(args)
	if argLen < 4 || argLen > 5 {
		fmt.Println("Usage: './philosophers <nb_philos> <time_to_die> <time_to_eat> <time_to_sleep> opt:<nb_of_meals'")
		os.Exit(1)
	}

	banket := &types.Banket{}

	parsing.Parser(args)
	initialisation.Initialiser(args, banket)
	simulation.Simulator(banket)

	os.Exit(0)
}