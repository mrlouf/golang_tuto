package simulation

import (
	"fmt"
	"time"
	"os"

	"tutogo/mod/philosophers/types"
)

func philosopherRoutine(philo *types.Philosopher) {
	for philo.IsAlive {
		think(philo)
		eat(philo)
		sleep(philo)
	}
}

func think(philo *types.Philosopher) {
	timestamp := uint64(time.Now().UnixMilli()) - philo.Banket.Start
	fmt.Printf("[%d] %d is thinking\n", timestamp, philo.Id)
}

func eat(philo *types.Philosopher) {

	timestamp := uint64(time.Now().UnixMilli()) - philo.Banket.Start
	fmt.Printf("[%d] %d has taken a fork\n", timestamp, philo.Id)


	timestamp = uint64(time.Now().UnixMilli()) - philo.Banket.Start
	fmt.Printf("[%d] %d is eating\n", timestamp, philo.Id)

	time.Sleep(time.Duration(philo.Banket.TimeToEat) * time.Millisecond)

}

func sleep(philo *types.Philosopher) {

	timestamp := uint64(time.Now().UnixMilli()) - philo.Banket.Start
	fmt.Printf("[%d] %d is sleeping\n", timestamp, philo.Id)

	time.Sleep(time.Duration(philo.Banket.TimeToSleep) * time.Millisecond)

}

func checkPhilosophers(banket *types.Banket) bool {
	return false
}

func Simulator(banket *types.Banket) {
	fmt.Println("Starting Simulation...")

	banket.Start = uint64(time.Now().UnixMilli())

	for i := 0; i < int(banket.NbOfPhilosophers); i++ {
		go philosopherRoutine(&banket.Philosophers[i])
	}

	for {
		if checkPhilosophers(banket) {
			os.Exit(0)
		}
		time.Sleep(1 * time.Millisecond)
	}

	fmt.Println("Simulation ended")
}