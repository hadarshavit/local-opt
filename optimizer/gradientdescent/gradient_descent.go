package gradientdescent

import (
	"math/rand"
	"time"

	"github.com/hadarshavit/local-opt/optimizer/base"
	"github.com/hadarshavit/local-opt/utils/assert"
)

const (
	deterministicCooldown = 1
)

// Optimizer uses Gradient Descent Algorithm
type Optimizer struct {
	provider base.StateProvider
	timeout time.Duration
	stopDiff int
	cur base.State
	best base.State
	start time.Time
	temperature float64
	cooldown float64
}

// NewGradientDescentOptimizer Constructs a new optimizer
func NewGradientDescentOptimizer(provider base.StateProvider, timeout time.Duration, stopDiff int) Optimizer {
	return Optimizer{provider: provider, timeout: timeout, stopDiff: stopDiff, cur: nil, start: time.Time{}, cooldown: deterministicCooldown}
}

// NewStochasticGradientDescentOptimizer Constructs a new optimizer
func NewStochasticGradientDescentOptimizer(provider base.StateProvider, timeout time.Duration, stopDiff int, cooldown float64) Optimizer {
	return Optimizer{provider: provider, timeout: timeout, stopDiff: stopDiff, cur: nil, start: time.Time{}, cooldown: cooldown, temperature: 0.99}
}

func (o Optimizer) hasTimePassed() bool {
	return time.Now().Sub(o.start)  >= o.timeout
}

func (o Optimizer) shouldGoRandom() bool {
	return float64(rand.Intn(100)) < o.temperature * 100
}

// Run the optimizer
func (o Optimizer) Run() base.State {
	assert.Assert(o.cur == nil, "Invalid State")

	o.start = time.Now()
	o.temperature = 0.99
	o.cur = o.provider.GenerateStartingState()
	o.best = o.cur

	for !o.hasTimePassed() {
		if !o.shouldGoRandom() {
			next := o.provider.GetBestNeighbor(o.cur)

			diff := next.Cost() - o.cur.Cost()
			if diff <= o.stopDiff {
				if next.Cost() < o.cur.Cost() {
					return next
				} 
	
				return o.cur
			}

			if o.cur.Cost() < o.best.Cost() {
				o.best = o.cur
			}

			o.cur = next
		} else {
			o.cur = o.provider.GetRandomNeighbor(o.cur)
		}
	}

	return o.best
}

// GetName Optimizer Name
func (o Optimizer) GetName() string {
	if o.cooldown == deterministicCooldown {
		return "Gradient Descent"
	}

	return "Sochastic Gradient Descent"
}