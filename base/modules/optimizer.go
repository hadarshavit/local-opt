package base

import (
	
)

// State of optimization problem
type State interface {
	Cost() int
}

// StateProvider abstract
type StateProvider interface {
	GenerateStartingState() State
	GetBestNeighbor(state State) State
}

// Optimizer for optimization problem
type Optimizer interface {
	Run() State
}