package base

import (
	"base/modules"
)

func main() {
	e := base.Employee {
		FirstName: "Hadar",
		LastName: "Shavit",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	e.LeavesRemaining()
}