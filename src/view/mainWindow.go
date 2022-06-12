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

func (m *MainWindow) SetReadyProcessesListText(text string) {
	m.OutputProcessesNotebook.readyProcessesTextView.SetBuffer(CreateTextBuffer(text))
}

func (m *MainWindow) SetDispatchedProcessesListText(text string) {
	m.OutputProcessesNotebook.dispatchedProcessesTextView.SetBuffer(CreateTextBuffer(text))
}

func (m *MainWindow) SetProcessedProcessesListText(text string) {
	m.OutputProcessesNotebook.processedProcessesTextView.SetBuffer(CreateTextBuffer(text))
}

func (m *MainWindow) SetBlockedProcessesListText(text string) {
	m.OutputProcessesNotebook.blockedProcessesTextView.SetBuffer(CreateTextBuffer(text))
}

func (m *MainWindow) SetAwokenProcessesListText(text string) {
	m.OutputProcessesNotebook.awokenProcessesTextView.SetBuffer(CreateTextBuffer(text))
}

func (m *MainWindow) SetResumedProcessesListText(text string) {
	m.OutputProcessesNotebook.resumedProcessesTextView.SetBuffer(CreateTextBuffer(text))
}

func (m *MainWindow) SetSuspendedProcessesListText(text string) {
	m.OutputProcessesNotebook.suspendedProcessesTextView.SetBuffer(CreateTextBuffer(text))
}

func (m *MainWindow) SetDestroyedProcessesListText(text string) {
	m.OutputProcessesNotebook.destroyedProcessesTextView.SetBuffer(CreateTextBuffer(text))
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
