package results

import "time"

// RunResults of an optimizer
type RunResults struct {
	ProblemID int
	Runtime time.Duration
	Result int
	OptimizerName string
}

// NewResults constructor
func NewResults(problemID int, runtime time.Duration, result int, OptimizerName string) RunResults{
	return RunResults{ProblemID: problemID, Runtime: runtime, Result: result, OptimizerName: OptimizerName}
}