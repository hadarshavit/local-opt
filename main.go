package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hadarshavit/local-opt/optimizer/base"
	"github.com/hadarshavit/local-opt/optimizer/gradientdescent"
	"github.com/hadarshavit/local-opt/problems/tsp"
	"github.com/hadarshavit/local-opt/runner"
	"github.com/hadarshavit/local-opt/runner/database"
	"github.com/hadarshavit/local-opt/runner/results"
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

	db := database.NewMongoAdapter()

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 30)
	var best *results.RunResults = nil
	
	err := db.Connect(ctx)
	if (err != nil) { log.Panicln(err) }

	for _, optimizer := range optimizers {
		cur := runner.ParallelRunner(optimizer)

		db.SaveRunResults(ctx, cur)
		
		if best == nil || cur.Result < best.Result {
			best = &cur
		}
	}

	fmt.Println()
	fmt.Println("Best Optimizer: ", best.OptimizerName)
	fmt.Println("Cost: ", best.Result, " Total run time: ", best.Runtime)
}