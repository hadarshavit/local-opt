package gradientdescent

import (
	"time"

	base "github.com/hadarshavit/local-opt/base/modules"
	"github.com/hadarshavit/local-opt/utils"
)

// Optimizer uses Gradient Descent Algorithm
type Optimizer struct {
	Provider base.StateProvider
	Timeout time.Duration
	StopDiff int
	cur base.State
	start time.Time
}

// NewOptimizer Constructs a new optimizer
func NewOptimizer(provider base.StateProvider, timeout time.Duration, stopDiff int) Optimizer {
	return Optimizer{Provider: provider, Timeout: timeout, StopDiff: stopDiff, cur: nil, start: time.Time{}}
}

func (o Optimizer) hasTimePassed() bool{
	return time.Now().Sub(o.start)  >= o.Timeout
}

// Run the optimizer
func (o Optimizer) Run() base.State {
	utils.Assert(o.cur == nil, "Invalid State")

	o.start = time.Now()
	o.cur = o.Provider.GenerateStartingState()

	for !o.hasTimePassed() {
		next := o.Provider.GetBestNeighbor(o.cur)

		diff := next.Cost() - o.cur.Cost()
		if diff <= o.StopDiff {
			if next.Cost() < o.cur.Cost() {
				return next
			} 

			return o.cur
		}

		o.cur = next
	}

	return o.cur
}