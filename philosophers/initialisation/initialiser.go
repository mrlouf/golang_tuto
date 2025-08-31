package initialisation

import (
	"fmt"

	"tutogo/mod/philosophers/types"
)

func Initialiser(args []string, banket *types.Banket) {
	fmt.Println("Starting Initialiser...")

	fmt.Println(args)
	fmt.Println(banket)
}