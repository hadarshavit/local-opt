package base

import (
	"log"
	"runtime"
	"time"

	"github.com/hadarshavit/local-opt/utils"
)

// Runner Parallel runner for optimizer
func Runner(optimizer Optimizer) State {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan State) 

	for i := 0; i < runtime.NumCPU(); i++ {
		go func () {
			c <- optimizer.Run()
		}()
	}
	var min State
	minVal := utils.MaxInt
	start := time.Now()
	for i := 0; i < runtime.NumCPU(); i++ {
		res := <- c
		runtime := time.Now().Sub(start)
		log.Println("Path: ", res, " Cost: ", res.Cost(), " Runtime: ", runtime)
		if res.Cost() < minVal {
			min = res
			minVal = min.Cost()
		}
	}

	return min
}