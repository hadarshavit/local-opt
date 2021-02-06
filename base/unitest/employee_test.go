package base

import (
	"local-opt/base/modules"
	"testing"
)

func TestEmployee(t *testing.T) {
	t.Logf("hello")
	var e base.Employee
	e.LeavesRemaining()
}
