package main

import (
	"fmt"
	"time"

	base "github.com/hadarshavit/local-opt/base/modules"
	"github.com/hadarshavit/local-opt/gradientdescent"
	"github.com/hadarshavit/local-opt/tsp"
)

func main() {
	tspProvider := tsp.NewProvider([]tsp.Point{tsp.NewPoint(0, 0), tsp.NewPoint(1, 2), tsp.NewPoint(2, 2),
		  tsp.NewPoint(2, 5), tsp.NewPoint(5, 2), tsp.NewPoint(4, 4), tsp.NewPoint(5, 5), tsp.NewPoint(6, 6),
		   tsp.NewPoint(7, 7), tsp.NewPoint(8, 8), tsp.NewPoint(9, 9), tsp.NewPoint(10, 10), tsp.NewPoint(11, 11), 
		tsp.NewPoint(50, 60), tsp.NewPoint(52, 67), tsp.NewPoint(30, 30)})
	optimizer := gradientdescent.NewOptimizer(tspProvider, time.Second * 60, 1)

	start := time.Now()
	best := base.Runner(optimizer)
	end := time.Now()

	fmt.Println(best, best.Cost())
	fmt.Println("Total run time", (end.Sub(start)))
}