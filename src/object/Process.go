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
	SUSPENDED_BLOCKED
	SUSPENDED_READY
	FINISHED
)

type Process struct {
	Name                 string
	Time                 int
	IsBlocked            bool
	IsSuspendedAtRunning bool
	IsSuspendedAtBlocked bool

	State                ProcessState
	TimeRemaining        int
}

func (p *Process) Process(time int) {
	p.TimeRemaining = int(math.Max(0.0, float64(p.TimeRemaining-time)))
}

func (p *Process) HasFinished() bool {
	return p.TimeRemaining == 0
}

func (p *Process) ToString() string {
	return fmt.Sprintf("%s %d %t\n", p.Name, p.Time, p.IsBlocked)
}
