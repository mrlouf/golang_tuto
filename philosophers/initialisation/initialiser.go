package initialisation

import (
	"strconv"
	"sync"

	"tutogo/mod/philosophers/types"
)

func Initialiser(args []string, banket *types.Banket) {

	nbPhilos, _ := strconv.Atoi(args[0])
	banket.NbOfPhilosophers = uint8(nbPhilos)

	banket.Philosophers = make([]types.Philosopher, nbPhilos)
	banket.Forks = make([]sync.Mutex, nbPhilos)

	for i := 0; i < nbPhilos; i++ {
		banket.Philosophers[i] = types.Philosopher{
			Id:        i + 1,
			IsAlive:   true,
			LeftFork:  &banket.Forks[i],
			RightFork: &banket.Forks[(i+1)%nbPhilos],
			Banket:    banket,
		}
	}

	tmp, _ := strconv.Atoi(args[1])
	banket.TimeToDie = int32(tmp)

	tmp, _ = strconv.Atoi(args[2])
	banket.TimeToEat = int32(tmp)

	tmp, _ = strconv.Atoi(args[3])
	banket.TimeToSleep = int32(tmp)

	if len(args) == 5 {
		tmp, _ = strconv.Atoi(args[4])
		banket.NbMeals = int32(tmp)
	}
}