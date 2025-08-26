package types

import (
	"sync"
)

type Philosopher struct {
	Id int
	IsAlive bool
	LeftFork *sync.Mutex
	RightFork *sync.Mutex
	Meals uint64
	LastMeal uint64
	Banket *Banket
}

type Banket struct {
	NbOfPhilosophers uint8
	Philosophers *Philosopher{}
	TimeToEat uint16
	TimeToSleep uint16
	NbMeals uint8
	Start uint64
}