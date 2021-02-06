package base

import (
	"testing"


	modules "github.com/hadarshavit/local-opt/base/modules"
)

func TestEmployee(t *testing.T) {
	t.Logf("hello")
	var e modules.Employee
	e.LeavesRemaining()
}
