package tsp

import (
	// "math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCost(t *testing.T) {
	p1 := Point{X: 2, Y: -1}
	p2 := Point{X: -2, Y: 2}
	state := State{Path: []Point{p1}}
	assert.Equal(t, 0, state.Cost(), "Unexpected cost")
	state.Path = append(state.Path, p2)
	assert.Equal(t, 5, state.Cost(), "Unexpected cost")
}

func equalPoints(t *testing.T, s1 []Point, s2 []Point) {
	assert.Equal(t, len(s1), len(s2))

	var pointsMap map[Point]bool = make(map[Point]bool)

	for _, v := range s1 {
		pointsMap[v] = true
	}

	for _, v := range s2 {
		if _, ok := pointsMap[v]; !ok {
			t.Error("Point not in original array", v)
		}

		if pointsMap[v] {
			pointsMap[v] = false
		} else {
			t.Error("Point appeared twice in array", v)
		}
	}
}

func TestGenerator(t *testing.T) {
	// n := rand.Intn(10000) + 1
	points := []Point{NewPoint(0, 0), NewPoint(1, 1), NewPoint(2, 2)}
	provider := NewProvider(points)

	startingState := provider.GenerateStartingState().(State)

	equalPoints(t, points, startingState.Path)
}

func TestBestNeighbor(t *testing.T) {
	points := []Point{NewPoint(0, 0), NewPoint(1, 1), NewPoint(2, 2)}
	provider := NewProvider(points)
	s := provider.GenerateStartingState()
	neighbor := provider.GetBestNeighbor(s).(State)

	equalPoints(t, points, neighbor.Path)
}

func TestCostIfSwap(t *testing.T) {
	points := []Point{NewPoint(0, 0), NewPoint(1, 1), NewPoint(2, 2), NewPoint(3, 3)}
	provider := NewProvider(points)
	s := provider.GenerateStartingState().(State)
	s.CostIfSwap(0, 1)
	s.CostIfSwap(1, 2)
	s.CostIfSwap(2, 3)

}
