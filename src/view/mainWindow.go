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
	ProcessActionPanel      *gtk.Box
	OutputProcessesNotebook *OutputProcessesNotebook
}

func (m *MainWindow) AddToReadyProcessesList(process *object.Process) {
    m.OutputProcessesNotebook.readyProcessesTreeView.AddRow(process)
}

func (m *MainWindow) RemoveFromReadyProcessesList(process *object.Process) {
    m.OutputProcessesNotebook.readyProcessesTreeView.RemoveRow(process)
}

func (m *MainWindow) AddToDispatchedProcessesList(process *object.Process) {
    m.OutputProcessesNotebook.dispatchedProcessesTreeView.AddRow(process)
}

func (m *MainWindow) AddToProcessedProcessesList(process *object.Process) {
    m.OutputProcessesNotebook.processedProcessesTreeView.AddRow(process)
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

func (m *MainWindow) AddToDestroyedProcessesList(process *object.Process) {
    m.OutputProcessesNotebook.destroyedProcessesTreeView.AddRow(process)
}

func (m *MainWindow) ResetLogs() {
    m.OutputProcessesNotebook.ResetTextViews()
}

func CreateMainWindow(listeners MainWindowListeners) *MainWindow {
	mainWindow := MainWindow{
		Window: CreateWindow(),
        ProcessActionPanel: CreateProcessActionPanel(listeners),
        OutputProcessesNotebook: CreateOutputProcessesNotebook(listeners),
	}
    paned := CreatePaned(gtk.ORIENTATION_HORIZONTAL)
	mainWindow.Window.Add(paned)

	// Add the label to the window.
	paned.Pack1(mainWindow.ProcessActionPanel, true, false)
	paned.Pack2(mainWindow.OutputProcessesNotebook.Box, true, true)

	return &mainWindow
}
