package base

import (
	"testing"
	"base/modules"
)

func TestEmployee(t *testing.T) {
	t.Logf("hello");
	var e base.Employee
	e.LeavesRemaining()
}