package parsing

import (
	"fmt"
	"os"
	"strconv"
	"math" // contains the constants like 'MaxInt"
)

func checkNumberPhilos(arg string) {
	tmp, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid number format")
		os.Exit(1)
	}

	if tmp < 1 || tmp > 200 {
		fmt.Println("Invalid number of philosophers (must be 1-200)")
		os.Exit(1)
	}
}

func checkTimeToDie(arg string) {
	tmp, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid number format")
		os.Exit(1)
	}

	if tmp < 60 || tmp > math.MaxInt {
		fmt.Println("Invalid time to die (must be 60-2148473647)")
		os.Exit(1)
	}
}

func checkTimeToEat(arg string) {
	tmp, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid number format")
		os.Exit(1)
	}

	if tmp < 60 || tmp > math.MaxInt {
		fmt.Println("Invalid time to eat (must be 60-2148473647)")
		os.Exit(1)
	}
}

func checkTimeToSleep(arg string) {
	tmp, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid number format")
		os.Exit(1)
	}

	if tmp < 60 || tmp > math.MaxInt {
		fmt.Println("Invalid time to sleep (must be 60-2148473647)")
		os.Exit(1)
	}
}

func checkNumberOfMeals(arg string) {
	tmp, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid number format")
		os.Exit(1)
	}

	if tmp < 0 || tmp > math.MaxInt {
		fmt.Println("Invalid number of meals (must be 0-2148473647)")
		os.Exit(1)
	}
}

func Parser(args []string) {

	checkNumberPhilos(args[0])
	checkTimeToDie(args[1])
	checkTimeToEat(args[2])
	checkTimeToSleep(args[3])
	if len(args) == 5  {
		checkNumberOfMeals(args[4])
	}
}