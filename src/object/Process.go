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
	SUSPENDED_RUNNING
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
	TimeRemaining int
}

func (p *Process) Process(time int) {
	p.TimeRemaining = int(math.Max(0.0, float64(p.TimeRemaining-time)))
}

func (p *Process) GetTimeRemaining() int {
    return p.TimeRemaining
}

func (p *Process) HasFinished() bool {
	return p.TimeRemaining == 0
}

func (p *Process) ToString() string {
	return fmt.Sprintf("%s %d %t\n", p.Name, p.Time, p.IsBlocked)
}
