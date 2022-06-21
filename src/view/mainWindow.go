package view

import (
	"github.com/goldencoderam/so-p2_processes/src/object"
	"github.com/gotk3/gotk3/gtk"
)

type MainWindowListeners interface {
	ProcessActionPanelListeners
	OutputProcessesNotebookListeners
}

type ProcessActionPanelListeners interface {
	ProcessFrameListeners

	StartProcessor()
	MakeProcessorTick()
	ResetProcessor()
}

type OutputProcessesNotebookListeners interface {
}

type ProcessFrameListeners interface {
	AddProcessButtonListener(process *object.Process)
}

type MainWindow struct {
	Window                  *gtk.Window
	ProcessActionPanel      *ProcessActionPanel
	OutputProcessesNotebook *LogProcessesNotebook
}

func (m *MainWindow) AddToReadyProcessesList(process *object.Process) {
	m.OutputProcessesNotebook.readyProcessesTreeView.AddRow(process)
}

func (m *MainWindow) RemoveFromReadyProcessesList(process *object.Process) {
	m.OutputProcessesNotebook.readyProcessesTreeView.RemoveRow(process)
}

func (m *MainWindow) AddToRunningProcessesList(process *object.Process) {
	m.OutputProcessesNotebook.runningProcessesTreeView.AddRow(process)
}

func (m *MainWindow) AddToBlockedProcessesList(process *object.Process) {
	m.OutputProcessesNotebook.blockedProcessesTreeView.AddRow(process)
}

func (m *MainWindow) AddToSuspendedReadyProcessesList(process *object.Process) {
	m.OutputProcessesNotebook.suspendedReadyProcessesTreeView.AddRow(process)
}

func (m *MainWindow) AddToSuspendedBlockedProcessesList(process *object.Process) {
	m.OutputProcessesNotebook.suspendedBlockedProcessesTreeView.AddRow(process)
}

func (m *MainWindow) AddToFinishedProcessesList(process *object.Process) {
	m.OutputProcessesNotebook.finishedProcessesTreeView.AddRow(process)
}

func (m *MainWindow) AddToDispatchTransitionList(process *object.Process) {
    m.OutputProcessesNotebook.dispatchTransitionTreeView.AddRow(process)
}

func (m *MainWindow) AddToTimeoutTransitionList(process *object.Process) {
    m.OutputProcessesNotebook.timeoutTransitionTreeView.AddRow(process)
}

func (m *MainWindow) AddToWaitEventTransitionList(process *object.Process) {
    m.OutputProcessesNotebook.waitEventTransitionTreeView.AddRow(process)
}

func (m *MainWindow) AddToCompletionEventTransitionList(process *object.Process) {
    m.OutputProcessesNotebook.completionEventTransitionTreeView.AddRow(process)
}

func (m *MainWindow) AddToSuspendBlockedTransitionList(process *object.Process) {
    m.OutputProcessesNotebook.suspendBlockedTransitionTreeView.AddRow(process)
}

func (m *MainWindow) AddToResumeSuspendedBlockedTransitionList(process *object.Process) {
    m.OutputProcessesNotebook.resumeSuspendedBlockedTransitionTreeView.AddRow(process)
}

func (m *MainWindow) AddToCompletionEventSuspendedBlockedTransitionList(process *object.Process) {
    m.OutputProcessesNotebook.completionEventSuspendedBlockedTransitionTreeView.AddRow(process)
}

func (m *MainWindow) AddToSuspendRunningTransitionList(process *object.Process) {
    m.OutputProcessesNotebook.suspendRunningTransitionTreeView.AddRow(process)
}

func (m *MainWindow) AddToResumeSuspendedReadyTransitionList(process *object.Process) {
    m.OutputProcessesNotebook.resumeSuspendedReadyTransitionTreeView.AddRow(process)
}

func (m *MainWindow) ResetLogs() {
	m.OutputProcessesNotebook.ResetTreeViews()
}

func CreateMainWindow(listeners MainWindowListeners) *MainWindow {
	mainWindow := MainWindow{
		Window:                  CreateWindow(),
		ProcessActionPanel:      CreateProcessActionPanel(listeners),
		OutputProcessesNotebook: CreateOutputProcessesNotebook(listeners),
	}
	paned := CreatePaned(gtk.ORIENTATION_HORIZONTAL)
	mainWindow.Window.Add(paned)

	// Add the label to the window.
	paned.Pack1(mainWindow.ProcessActionPanel.PanelBox, true, false)
	paned.Pack2(mainWindow.OutputProcessesNotebook.Box, true, true)

	return &mainWindow
}
