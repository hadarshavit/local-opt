package runner

import (
	"log"
	"runtime"
	"time"

	"github.com/hadarshavit/local-opt/optimizer/base"
	"github.com/hadarshavit/local-opt/runner/results"
	"github.com/hadarshavit/local-opt/utils/limits"
)

// ParallelRunner Parallel runner for optimizer
func ParallelRunner(optimizer base.Optimizer) results.RunResults {
	numRoutines := runtime.NumCPU()
	c := make(chan base.State) 

	start := time.Now()

	for i := 0; i < numRoutines; i++ {
		go func () {
			c <- optimizer.Run()
		}()
	}
	var min base.State
	minVal := limits.MaxInt
	for i := 0; i < numRoutines; i++ {
		res := <- c
		runtime := time.Now().Sub(start)
		log.Println("Optimizer: ", optimizer.GetName(), "Path: ", res, " Cost: ", res.Cost(), " Runtime: ", runtime)
		if res.Cost() < minVal {
			min = res
			minVal = min.Cost()
		}
	}

	totalTime := time.Now().Sub(start)

	log.Println("Optimizer: ", optimizer.GetName(), "Path: ", min, " Cost: ", min.Cost(), " Runtime: ", totalTime)

	return results.NewResults(0, time.Now().Sub(start), min.Cost(), optimizer.GetName())
}