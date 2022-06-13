package view

import (
	"github.com/goldencoderam/so-p2_processes/src/view/process"
	"github.com/gotk3/gotk3/gtk"
)

type OutputProcessesNotebook struct {
	Box      *gtk.Box
	Notebook *gtk.Notebook

	readyProcessesTreeView            *process.ProcessTreeView
	dispatchedProcessesTreeView       *process.ProcessTreeView
	processedProcessesTreeView        *process.ProcessTreeView
	blockedProcessesTreeView          *process.ProcessTreeView
	suspendedReadyProcessesTreeView   *process.ProcessTreeView
	suspendedBlockedProcessesTreeView *process.ProcessTreeView
	destroyedProcessesTreeView        *process.ProcessTreeView
	communicationProcessesTextView    *gtk.TextView
}

func CreateOutputProcessesNotebook(listeners OutputProcessesNotebookListeners) *OutputProcessesNotebook {
	outputNotebook := OutputProcessesNotebook{
		Box:      CreateBox(gtk.ORIENTATION_VERTICAL, ZeroMargin),
		Notebook: CreateNotebook(),

		readyProcessesTreeView:            process.NewTreeView(),
		dispatchedProcessesTreeView:       process.NewTreeView(),
		processedProcessesTreeView:        process.NewTreeView(),
		blockedProcessesTreeView:          process.NewTreeView(),
		suspendedReadyProcessesTreeView:   process.NewTreeView(),
		suspendedBlockedProcessesTreeView: process.NewTreeView(),
		destroyedProcessesTreeView:        process.NewTreeView(),
		communicationProcessesTextView:    CreateTextView(),
	}

	headerBar := CreateHeaderBar()

	outputNotebook.Notebook.AppendPage(outputNotebook.readyProcessesTreeView.TreeView, CreateLabel("Ready"))
	outputNotebook.Notebook.AppendPage(outputNotebook.dispatchedProcessesTreeView.TreeView, CreateLabel("Dispatched"))
	outputNotebook.Notebook.AppendPage(outputNotebook.processedProcessesTreeView.TreeView, CreateLabel("Processed"))
	outputNotebook.Notebook.AppendPage(outputNotebook.blockedProcessesTreeView.TreeView, CreateLabel("Blocked"))
	outputNotebook.Notebook.AppendPage(outputNotebook.suspendedReadyProcessesTreeView.TreeView, CreateLabel("Suspended-ready"))
	outputNotebook.Notebook.AppendPage(outputNotebook.suspendedBlockedProcessesTreeView.TreeView, CreateLabel("Suspended-blocked"))
	outputNotebook.Notebook.AppendPage(outputNotebook.destroyedProcessesTreeView.TreeView, CreateLabel("Destroyed"))
    outputNotebook.Notebook.AppendPage(outputNotebook.communicationProcessesTextView, CreateLabel("Communication"))

	headerBar.SetTitle("Output processes")
	outputNotebook.Box.Add(headerBar)
	outputNotebook.Box.PackEnd(outputNotebook.Notebook, true, true, 0)

	return &outputNotebook
}

func (o *OutputProcessesNotebook) ResetTextViews() {
	o.readyProcessesTreeView.Clear()
	o.dispatchedProcessesTreeView.Clear()
	o.processedProcessesTreeView.Clear()
	o.blockedProcessesTreeView.Clear()
	o.suspendedReadyProcessesTreeView.Clear()
	o.suspendedBlockedProcessesTreeView.Clear()
	o.destroyedProcessesTreeView.Clear()
}
