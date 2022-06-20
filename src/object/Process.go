package object

import (
	"fmt"
	"math"
)

type ProcessState int

const (
	READY ProcessState = iota
	RUNNING
	BLOCKED
    // When it was suspended already
    BLOCKED_NOT_SUSPENDED
	SUSPENDED_BLOCKED
	SUSPENDED_READY
	FINISHED
)

type Process struct {
	Name                      string
	Time                      int
	IsBlocked                 bool
	IsSuspendedAtRunning      bool
	IsSuspendedAtBlocked      bool
	IsSuspendedAtIOCompletion bool
	State                     ProcessState

	timeRemaining int
}

func (p *Process) Process(time int) {
	p.timeRemaining = int(math.Max(0.0, float64(p.timeRemaining-time)))
}

func (p *Process) HasFinished() bool {
	return p.timeRemaining == 0
}

func (p *Process) ToString() string {
	return fmt.Sprintf("%s %d %t\n", p.Name, p.Time, p.IsBlocked)
}
