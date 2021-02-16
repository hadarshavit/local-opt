package main

import (
	"fmt"
	"log"
	"os"
	"time"

	base "github.com/hadarshavit/local-opt/base/modules"
	"github.com/hadarshavit/local-opt/gradientdescent"
	"github.com/hadarshavit/local-opt/tsp"
)

func configLogger() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

	log.SetOutput(file)
}

func main() {
	configLogger()
	tspProvider := tsp.NewProvider([]tsp.Point{tsp.NewPoint(0, 0), tsp.NewPoint(1, 2), tsp.NewPoint(2, 2),
		  tsp.NewPoint(2, 5), tsp.NewPoint(5, 2), tsp.NewPoint(4, 4), tsp.NewPoint(5, 5), tsp.NewPoint(6, 6),
		   tsp.NewPoint(7, 7), tsp.NewPoint(8, 8), tsp.NewPoint(9, 9), tsp.NewPoint(10, 10), tsp.NewPoint(11, 11), 
		tsp.NewPoint(50, 60), tsp.NewPoint(52, 67), tsp.NewPoint(30, 30)})
	optimizers := []base.Optimizer{gradientdescent.NewGradientDescentOptimizer(tspProvider, time.Second * 60, 1),
		gradientdescent.NewStochasticGradientDescentOptimizer(tspProvider, time.Second * 60, 1, 0.5),
	}

	var bestOptimizer base.Optimizer = nil
	var bestResult base.State
	var bestTime time.Duration
	
	for _, optimizer := range optimizers {
		start := time.Now()
		best := base.Runner(optimizer)
		end := time.Now()

		totalTime := end.Sub(start)

		fmt.Println("Optimizer: ", optimizer.GetName())
		fmt.Println("Cost: ", best.Cost(), " Total run time: ", totalTime)
		fmt.Println("Route: ", best)

		if bestResult == nil  || best.Cost() < bestResult.Cost() {
			bestOptimizer = optimizer
			bestResult = best
			bestTime = totalTime
		}
	}

	fmt.Println()
	fmt.Println("Best Optimizer: ", bestOptimizer.GetName())
	fmt.Println("Cost: ", bestResult.Cost(), " Total run time: ", bestTime)
	fmt.Println("Route: ", bestResult)	
}