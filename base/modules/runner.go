package base

import (
	"fmt"
	"runtime"
	"time"

	"github.com/hadarshavit/local-opt/utils"
)

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
		fmt.Println(res, res.Cost(), runtime)
		if res.Cost() < minVal {
			min = res
			minVal = min.Cost()
		}
	}

	return min
}