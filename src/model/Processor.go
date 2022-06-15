package model

import (
	"sort"
	"github.com/goldencoderam/so-p2_processes/src/object"
)

type ProcessorLogListeners interface {
    // Ready - Running
	DispatchProcess(process *object.Process)
    // Ready - SuspendedReady
    SuspendReadyProcess(process *object.Process)
    // Running - Ready
	TimerRunoutProcess(process *object.Process)
    // Running - Blocked
	BlockProcess(process *object.Process)
    // Running - SuspendedReady
    SuspendRunningProcess(process *object.Process)
    // Running - Finished
	FinishedProcess(process *object.Process)
    // Blocked - Ready
    IOBlockedCompletionProcess(process *object.Process)
    // Blocked - SuspendedBlocked
    SuspendBlockedProcess(process *object.Process)
    // SuspendedBlocked - Blocked
    ResumeSuspendedBlockedProcess(process *object.Process)
    // SuspendedBlocked - SuspendedReady
    IOSuspendedBlockedCompletionProcess(process *object.Process)
    // SuspendedReady - Ready
    ResumeSuspendedReadyProcess(process *object.Process)
}

const ProcessingTime = 5

type Processor struct {
	ReadyProcessesList []*object.Process
	CurrentProcess     *object.Process
}

func (p *Processor) AddProcessToReadyList(process *object.Process) bool {
    for _, p := range p.ReadyProcessesList {
        if p.Name == process.Name {
            return false
        }
    }
	p.ReadyProcessesList = append(p.ReadyProcessesList, process)
    return true
}

func (p *Processor) Reset() {
	p.ReadyProcessesList = make([]*object.Process, 0)
	p.CurrentProcess = nil
}

func (p *Processor) MakeTick(listeners ProcessorLogListeners) {
	if p.CurrentProcess == nil {
		if len(p.ReadyProcessesList) > 0 {
            sort.SliceStable(p.ReadyProcessesList, func(i, j int) bool {
                return p.ReadyProcessesList[i].Priority > p.ReadyProcessesList[j].Priority
            })
			p.CurrentProcess = p.ReadyProcessesList[0]
			p.ReadyProcessesList = p.ReadyProcessesList[1:]
		} else {
			listeners.FinishedProcessing()
			return
		}
	}

	switch p.CurrentProcess.State {
	case object.READY:
		if p.CurrentProcess.IsSuspendedAtReady {
			p.CurrentProcess.State = object.SUSPENDED_READY
            p.CurrentProcess.IsSuspendedAtReady = false
			listeners.SuspendedReadyProcess(p.CurrentProcess)
		} else {
			p.CurrentProcess.State = object.RUNNING
			p.CurrentProcess.Process(ProcessingTime)
            listeners.CommunicateWithProcess(p.CurrentProcess)
			listeners.DispatchProcess(p.CurrentProcess)
		}
		break
	case object.SUSPENDED_READY:
		p.CurrentProcess.State = object.READY
		listeners.TimerRunoutProcess(p.CurrentProcess)
		p.CurrentProcess = nil
		break
	case object.RUNNING:
		if p.CurrentProcess.HasFinished() {
			p.CurrentProcess.State = object.FINISHED
			listeners.FinishedProcess(p.CurrentProcess)
			p.CurrentProcess = nil
		} else if p.CurrentProcess.IsDeleted {
			p.CurrentProcess.State = object.DESTROYED
			listeners.DestroyedProcess(p.CurrentProcess)
			p.CurrentProcess = nil
		} else if p.CurrentProcess.IsBlocked {
			p.CurrentProcess.State = object.BLOCKED
			listeners.BlockedProcess(p.CurrentProcess)
		} else if p.CurrentProcess.IsSuspendedAtRunning {
			p.CurrentProcess.State = object.SUSPENDED_READY
			listeners.SuspendedReadyProcess(p.CurrentProcess)
		} else {
			p.CurrentProcess.State = object.READY
			listeners.TimerRunoutProcess(p.CurrentProcess)
			p.CurrentProcess = nil
		}
		break
	case object.BLOCKED:
		if p.CurrentProcess.IsSuspendedAtBlocked {
			p.CurrentProcess.State = object.SUSPENDED_BLOCKED
			listeners.SuspendedBlockedProcess(p.CurrentProcess)
		} else {
			p.CurrentProcess.State = object.READY
			listeners.TimerRunoutProcess(p.CurrentProcess)
			p.CurrentProcess = nil
		}
		break
	case object.SUSPENDED_BLOCKED:
		p.CurrentProcess.State = object.READY
		listeners.TimerRunoutProcess(p.CurrentProcess)
		p.CurrentProcess = nil
		break
	}

}
