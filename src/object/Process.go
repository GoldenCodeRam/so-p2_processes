package object

import "fmt"

type ProcessState int
const (
	READY ProcessState = iota
	RUNNING
	BLOCKED
    SUSPENDED
)

type Process struct {
	Name      string
	Time      int
	IsBlocked bool
	State     ProcessState
	priority  int
}

func (p *Process) ToString() string {
	return fmt.Sprintf("%s %d %t\n", p.Name, p.Time, p.IsBlocked)
}
