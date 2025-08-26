package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	argLen := len(args)
	if argLen < 4 || argLen > 5 {
		fmt.Println("Usage: './philosophers <nb_philos> <time_to_eat> <time_to_sleep> opt:<nb_of_meals'")
		os.Exit(1)
	}

}