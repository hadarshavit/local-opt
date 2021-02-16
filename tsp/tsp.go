package tsp

import (
	"math"
	"math/rand"

	base "github.com/hadarshavit/local-opt/base/modules"
	"github.com/hadarshavit/local-opt/utils"
)

// Point with x,y axis
type Point struct {
	X int
	Y int
}

func (p1 Point) distance (p2 Point) int {
	return int(math.Sqrt(math.Pow(float64(p1.X - p2.X), 2) + math.Pow(float64(p1.Y - p2.Y), 2)))
}

// NewPoint creates a point
func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

// State Travelling Salesman Problem's state
type State struct {
	Path []Point
}


// Cost of the state
func (state State) Cost() int {
	if (len(state.Path)) < 2 {
		return 0
	}

	var total int = 0
	
	for i := 1; i < len(state.Path); i++ {
		total += state.Path[i - 1].distance(state.Path[i])
	}

	return total
}

// BasicProvider for TSP
type BasicProvider struct {
	Points []Point
}

// GenerateStartingState TSP random state
func (b BasicProvider) GenerateStartingState() base.State {
	state := State{}
	state.Path = b.Points
	rand.Shuffle(len(state.Path), func(i, j int) { state.Path[i], state.Path[j] = state.Path[j], state.Path[i] })

	return state
}

func (state State) swap(i int, j int) State {
	res := State{}
	res.Path = make([]Point, len(state.Path))
	copy(res.Path, state.Path)
	res.Path[i] = state.Path[j]
	res.Path[j] = state.Path[i]

	return res
}

// CostIfSwap Cost of the state if swapping p1 and p2
func (state State) CostIfSwap(p1 int, p2 int) int {
	cost := 0

	if p1 > 0 {
		for i := 1; i < p1; i++ {
			cost += state.Path[i - 1].distance(state.Path[i])
		}
	
		cost += state.Path[p1 - 1].distance(state.Path[p2])
	}

	cost += state.Path[p2].distance(state.Path[p1 + 1])

	for i := p1 + 1; i < p2; i++ {
		cost += state.Path[i - 1].distance(state.Path[i])
	}

	cost += state.Path[p2 - 1].distance(state.Path[p1])

	if p2 < len(state.Path) - 1 {
		cost += state.Path[p1].distance(state.Path[p2 + 1])

		for i := p2 + 1; i < len(state.Path); i++ {
			cost += state.Path[i - 1].distance(state.Path[i])
		}
	}


	return cost
}

// GetBestNeighbor Returns best neighbor
func (b BasicProvider) GetBestNeighbor(state base.State) base.State {
	tspState := state.(State)
	p1, p2 := 0, 0
	bestCost := utils.MaxInt
	for i := 0; i < len(tspState.Path) - 1; i++ {
		for j := i + 1; j <len(tspState.Path); j++ {
			costIfSwap := tspState.CostIfSwap(i, j)
			if costIfSwap < bestCost {
				p1, p2 = i, j
				bestCost = costIfSwap
			}
		}
	}

	return tspState.swap(p1, p2)
}

// GetRandomNeighbor Returns random neighbor
func (b BasicProvider) GetRandomNeighbor(state base.State) base.State {
	tspState := state.(State)
	p1 := rand.Intn(len(tspState.Path) - 1)
	p2 := p1 + rand.Intn(len(tspState.Path) - p1)
	return tspState.swap(p1, p2)
}

// NewProvider Constructs a new provider
func NewProvider(points[] Point) BasicProvider {
	return BasicProvider{Points: points}
}