package model

import (
	"github.com/goldencoderam/so-p2_processes/src/object"
)

type ProcessorLogListeners interface {
	UpdateListsText()
}

const ProcessingTime = 5

type Processor struct {
	ReadyProcessesList []*object.Process
	CurrentProcess     *object.Process

	DispatchedProcessesLog string
	ProcessedProcessesLog  string
	BlockedProcessesLog    string
	AwokenProcessesLog     string

	ResumedProcessesLog   string
	SuspendedProcessesLog string
	DestroyedProcessesLog string
}

func (p *Processor) AddProcessToReadyList(process *object.Process) {
	p.ReadyProcessesList = append(p.ReadyProcessesList, process)
}

func (p *Processor) Reset() {
	p.ReadyProcessesList = make([]*object.Process, 0)
	p.CurrentProcess = nil

	p.DispatchedProcessesLog = ""
	p.ProcessedProcessesLog = ""
	p.BlockedProcessesLog = ""
	p.AwokenProcessesLog = ""

	p.ResumedProcessesLog = ""
	p.SuspendedProcessesLog = ""
	p.DestroyedProcessesLog = ""
}

func (p *Processor) MakeTick(listeners ProcessorLogListeners) {
	if len(p.ReadyProcessesList) > 0 {
		p.CurrentProcess = p.ReadyProcessesList[0]
		p.ReadyProcessesList = p.ReadyProcessesList[1:]

		switch p.CurrentProcess.State {
		case object.READY:
			p.DispatchedProcessesLog += "Dispatched: " + p.CurrentProcess.ToString()
			listeners.UpdateListsText()
			break
		}
	}
}
