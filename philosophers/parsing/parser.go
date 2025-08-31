package parsing

import (
	"fmt"
	"os"
	"strconv"
)

func Parser(args []string) {
	fmt.Println("Starting Parser...")
	fmt.Println(args)

	tmp, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid number format")
		os.Exit(1)
	}

	if tmp < 1 || tmp > 200 {
		fmt.Println("Invalid number of philosophers (must be 1-200)")
		os.Exit(1)
	}
}