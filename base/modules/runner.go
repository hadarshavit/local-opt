package base

import (
	"fmt"
	"runtime"

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
	for i := 0; i < runtime.NumCPU(); i++ {
		res := <- c
		fmt.Println(res, res.Cost())
		if res.Cost() < minVal {
			min = res
			minVal = min.Cost()
		}
	}

	return min
}