package model

import (
	"log"

	"github.com/goldencoderam/so-p2_processes/src/object"
)

const ProcessingTime = 5

type Processor struct {
	ReadyProcessesList []*object.Process
	RunningProcess     *object.Process

	ReadyProcessesLog      string
	DispatchedProcessesLog string
    ProcessedProcessesLog  string
    BlockedProcessesLog    string
    AwokenProcessesLog     string

    ResumedProcessesLog    string
    SuspendedProcessesLog  string
    DestroyedProcessesLog  string
}

func (p *Processor) AddProcessToReadyList(process *object.Process) {
	p.ReadyProcessesList = append(p.ReadyProcessesList, process)
    p.ReadyProcessesLog += process.ToString()
}

func (p *Processor) Reset() {
    p.ReadyProcessesList = make([]*object.Process, 0)
    p.RunningProcess = nil

    p.ReadyProcessesLog = ""
    p.DispatchedProcessesLog = ""
    p.ProcessedProcessesLog = ""
    p.BlockedProcessesLog = ""
    p.AwokenProcessesLog = ""

    p.ResumedProcessesLog = ""
    p.SuspendedProcessesLog = ""
    p.DestroyedProcessesLog = ""
}

func (p *Processor) MakeTick() {
    process := p.ReadyProcessesList[0]
    log.Default().Println(process)
}
