package model

import (
	"errors"

	"github.com/goldencoderam/so-p2_processes/src/object"
)

type ProcessorLogListeners interface {
	// Ready - Running
	LogProcessDispatched(process *object.Process)
	// Ready - SuspendedReady
	//LogProcessSuspendedReady(process *object.Process)
	// Running - Ready
	LogProcessTimeout(process *object.Process)
	// Running - Blocked
	LogProcessBlocked(process *object.Process)
	// Running - SuspendedReady
	LogProcessSuspendedRunning(process *object.Process)
	// Running - Finished
	LogProcessFinished(process *object.Process)
	// Blocked - Ready
	LogProcessIOBlockedCompleted(process *object.Process)
	// Blocked - SuspendedBlocked
	LogProcessSuspendedBlocked(process *object.Process)
	// SuspendedBlocked - Blocked
	LogProcessSuspendedBlockedResumed(process *object.Process)
	// SuspendedBlocked - SuspendedReady
	LogProcessIOSuspendedBlockedCompleted(process *object.Process)
	// SuspendedReady - Ready
	LogProcessSuspendedReadyResumed(process *object.Process)

	// Finished all processes
	LogFinishedProcessing()
}

const ProcessingTime = 5

type Processor struct {
	ReadyProcessesList []*object.Process
	CurrentProcess     *object.Process
	LogListeners       ProcessorLogListeners
}

func (p *Processor) AddProcessToReadyList(process *object.Process) error {
	for _, p := range p.ReadyProcessesList {
		if p.Name == process.Name {
			return errors.New("Process with repeated name")
		}
	}

	p.ReadyProcessesList = append(p.ReadyProcessesList, process)
	return nil
}

func (p *Processor) GetNextProcess() (*object.Process, error) {
	if len(p.ReadyProcessesList) > 0 {
		result := p.ReadyProcessesList[0]
		p.ReadyProcessesList = p.ReadyProcessesList[1:]
		return result, nil
	} else {
		return nil, errors.New("Empty ready processes list")
	}
}

func (p *Processor) Reset() {
	p.ReadyProcessesList = make([]*object.Process, 0)
	p.CurrentProcess = nil
}

func (p *Processor) MakeTick(listeners ProcessorLogListeners) {
	if p.CurrentProcess == nil {
		nextProcess, err := p.GetNextProcess()
		if err != nil {
			listeners.LogFinishedProcessing()
		} else {
			p.CurrentProcess = nextProcess
		}
	}

	switch p.CurrentProcess.State {
	case object.READY:
		p.makeProcessReadyTransition()
		break
	case object.RUNNING:
		p.makeProcessRunningTransition()
		break
	case object.BLOCKED:
		p.makeProcessBlockedTransition()
		break
    case object.BLOCKED_NOT_SUSPENDED:
        p.makeProcessBlockedNotSuspendedTransition()
        break
	case object.SUSPENDED_BLOCKED:
        p.makeProcessSuspendedBlockedTransition()
		break
	case object.SUSPENDED_RUNNING:
        p.makeProcessSuspendedReadyTransition()
		break
	case object.FINISHED:
		panic("State FINISHED should never happen!")
	default:
		panic("Process with an undefined state!")
	}
}

func (p *Processor) makeProcessReadyTransition() {
	p.CurrentProcess.State = object.RUNNING
	p.CurrentProcess.Process(ProcessingTime)
	p.LogListeners.LogProcessDispatched(p.CurrentProcess)
}

func (p *Processor) makeProcessRunningTransition() {
	if p.CurrentProcess.HasFinished() {
		p.CurrentProcess.State = object.FINISHED
		p.LogListeners.LogProcessFinished(p.CurrentProcess)
		p.CurrentProcess = nil
	} else if p.CurrentProcess.IsBlocked {
		p.CurrentProcess.State = object.BLOCKED
		p.LogListeners.LogProcessBlocked(p.CurrentProcess)
	} else if p.CurrentProcess.IsSuspendedAtRunning {
		p.CurrentProcess.State = object.SUSPENDED_RUNNING
		p.LogListeners.LogProcessSuspendedRunning(p.CurrentProcess)
	} else {
		p.CurrentProcess.State = object.READY
		p.AddProcessToReadyList(p.CurrentProcess)
		p.LogListeners.LogProcessTimeout(p.CurrentProcess)
		p.CurrentProcess = nil
	}
}

func (p *Processor) makeProcessBlockedTransition() {
	if p.CurrentProcess.IsSuspendedAtBlocked {
		p.CurrentProcess.State = object.SUSPENDED_BLOCKED
		p.LogListeners.LogProcessSuspendedBlocked(p.CurrentProcess)
	} else {
        p.CurrentProcess.State = object.READY
        p.AddProcessToReadyList(p.CurrentProcess)
        p.LogListeners.LogProcessIOBlockedCompleted(p.CurrentProcess)
        p.CurrentProcess = nil
	}
}

func (p *Processor) makeProcessBlockedNotSuspendedTransition() {
    p.CurrentProcess.State = object.READY
    p.AddProcessToReadyList(p.CurrentProcess)
    p.LogListeners.LogProcessIOBlockedCompleted(p.CurrentProcess)
    p.CurrentProcess = nil
}

func (p *Processor) makeProcessSuspendedBlockedTransition() {
    if p.CurrentProcess.IsSuspendedAtIOCompletion {
        p.CurrentProcess.State = object.SUSPENDED_RUNNING
        p.LogListeners.LogProcessIOSuspendedBlockedCompleted(p.CurrentProcess)
    } else {
        p.CurrentProcess.State = object.BLOCKED_NOT_SUSPENDED
        p.LogListeners.LogProcessSuspendedBlockedResumed(p.CurrentProcess)
    }
}

func (p *Processor) makeProcessSuspendedReadyTransition() {
    p.CurrentProcess.State = object.READY
    p.AddProcessToReadyList(p.CurrentProcess)
    p.LogListeners.LogProcessSuspendedReadyResumed(p.CurrentProcess)
    p.CurrentProcess = nil
}
