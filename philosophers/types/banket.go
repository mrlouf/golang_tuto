package types

import (
	"sync"
)

type Philosopher struct {
	Id 			int
	IsAlive 	bool
	LeftFork 	*sync.Mutex
	RightFork 	*sync.Mutex
	Meals 		int32
	LastMeal 	uint64
	Banket 		*Banket
}

type Banket struct {
	NbOfPhilosophers	uint8
	Philosophers 		[]Philosopher
	Forks				[]sync.Mutex
	TimeToDie 			int32
	TimeToEat 			int32
	TimeToSleep 		int32
	NbMeals 			int32
	Start 				uint64
	PrintMutex			sync.Mutex
	StatusMutex			sync.Mutex
}