package view

import (
	"github.com/goldencoderam/so-p2_processes/src/view/process"
	"github.com/gotk3/gotk3/gtk"
)

type LogProcessesNotebook struct {
	Box      *gtk.Box
	Notebook *gtk.Notebook

	// States
	readyProcessesTreeView            *process.ProcessTreeView
	runningProcessesTreeView          *process.ProcessTreeView
	blockedProcessesTreeView          *process.ProcessTreeView
	suspendedBlockedProcessesTreeView *process.ProcessTreeView
	suspendedReadyProcessesTreeView   *process.ProcessTreeView
	finishedProcessesTreeView         *process.ProcessTreeView
	// Transitions
	dispatchTransitionTreeView                        *process.ProcessTreeView
	timeoutTransitionTreeView                         *process.ProcessTreeView
	waitEventTransitionTreeView                       *process.ProcessTreeView
	completionEventTransitionTreeView                 *process.ProcessTreeView
	suspendBlockedTransitionTreeView                  *process.ProcessTreeView
	resumeSuspendedBlockedTransitionTreeView          *process.ProcessTreeView
	completionEventSuspendedBlockedTransitionTreeView *process.ProcessTreeView
	suspendRunningTransitionTreeView                  *process.ProcessTreeView
	resumeSuspendedReadyTransitionTreeView            *process.ProcessTreeView
	// This transition is not needed yet
	//suspendReadyTransitionTreeView                  *process.ProcessTreeView
}

func CreateOutputProcessesNotebook(listeners OutputProcessesNotebookListeners) *LogProcessesNotebook {
	outputNotebook := LogProcessesNotebook{
		Box:      CreateBox(gtk.ORIENTATION_VERTICAL, ZeroMargin),
		Notebook: CreateNotebook(),

		readyProcessesTreeView:                            process.NewTreeView(),
		runningProcessesTreeView:                          process.NewTreeView(),
		blockedProcessesTreeView:                          process.NewTreeView(),
		suspendedBlockedProcessesTreeView:                 process.NewTreeView(),
		suspendedReadyProcessesTreeView:                   process.NewTreeView(),
		finishedProcessesTreeView:                         process.NewTreeView(),
		dispatchTransitionTreeView:                        process.NewTreeView(),
		timeoutTransitionTreeView:                         process.NewTreeView(),
		waitEventTransitionTreeView:                       process.NewTreeView(),
		completionEventTransitionTreeView:                 process.NewTreeView(),
		suspendBlockedTransitionTreeView:                  process.NewTreeView(),
		resumeSuspendedBlockedTransitionTreeView:          process.NewTreeView(),
		completionEventSuspendedBlockedTransitionTreeView: process.NewTreeView(),
		suspendRunningTransitionTreeView:                  process.NewTreeView(),
		resumeSuspendedReadyTransitionTreeView:            process.NewTreeView(),
	}

	headerBar := CreateHeaderBar()

	// States
	outputNotebook.Notebook.AppendPage(outputNotebook.readyProcessesTreeView.TreeView, CreateLabel("Listos"))
	outputNotebook.Notebook.AppendPage(outputNotebook.runningProcessesTreeView.TreeView, CreateLabel("En ejecución"))
	outputNotebook.Notebook.AppendPage(outputNotebook.blockedProcessesTreeView.TreeView, CreateLabel("Bloqueados"))
	outputNotebook.Notebook.AppendPage(outputNotebook.suspendedBlockedProcessesTreeView.TreeView, CreateLabel("Suspendido-bloqueado"))
	outputNotebook.Notebook.AppendPage(outputNotebook.suspendedReadyProcessesTreeView.TreeView, CreateLabel("Suspendido-listo"))
	outputNotebook.Notebook.AppendPage(outputNotebook.finishedProcessesTreeView.TreeView, CreateLabel("Finalizado"))
	// Transitions
	outputNotebook.Notebook.AppendPage(outputNotebook.dispatchTransitionTreeView.TreeView, CreateLabel("Listos a en ejecución"))
	outputNotebook.Notebook.AppendPage(outputNotebook.timeoutTransitionTreeView.TreeView, CreateLabel("En ejecución a listos"))
	outputNotebook.Notebook.AppendPage(outputNotebook.waitEventTransitionTreeView.TreeView, CreateLabel("En ejecución a bloqueados"))
	outputNotebook.Notebook.AppendPage(outputNotebook.completionEventTransitionTreeView.TreeView, CreateLabel("Bloqueados a listos"))
	outputNotebook.Notebook.AppendPage(outputNotebook.suspendBlockedTransitionTreeView.TreeView, CreateLabel("Bloqueados a suspendidos"))
	outputNotebook.Notebook.AppendPage(outputNotebook.resumeSuspendedBlockedTransitionTreeView.TreeView, CreateLabel("Suspendido-bloqueado a bloqueado"))
	outputNotebook.Notebook.AppendPage(outputNotebook.completionEventSuspendedBlockedTransitionTreeView.TreeView, CreateLabel("Suspendido-bloqueado a suspendido-listo"))
	outputNotebook.Notebook.AppendPage(outputNotebook.suspendRunningTransitionTreeView.TreeView, CreateLabel("En ejecución a suspendido"))
	outputNotebook.Notebook.AppendPage(outputNotebook.resumeSuspendedReadyTransitionTreeView.TreeView, CreateLabel("Suspendido-listo a listo"))

	headerBar.SetTitle("Resultado de los procesos")
	outputNotebook.Box.Add(headerBar)
	outputNotebook.Box.PackEnd(outputNotebook.Notebook, true, true, 0)

	return &outputNotebook
}

func (o *LogProcessesNotebook) ResetTreeViews() {
	o.readyProcessesTreeView.Clear()
    o.runningProcessesTreeView.Clear()
    o.blockedProcessesTreeView.Clear()
    o.suspendedBlockedProcessesTreeView.Clear()
    o.suspendedReadyProcessesTreeView.Clear()
    o.finishedProcessesTreeView.Clear()
    o.dispatchTransitionTreeView.Clear()
    o.timeoutTransitionTreeView.Clear()
    o.waitEventTransitionTreeView.Clear()
    o.completionEventTransitionTreeView.Clear()
    o.suspendBlockedTransitionTreeView.Clear()
    o.resumeSuspendedBlockedTransitionTreeView.Clear()
    o.completionEventSuspendedBlockedTransitionTreeView.Clear()
    o.suspendRunningTransitionTreeView.Clear()
    o.resumeSuspendedReadyTransitionTreeView.Clear()
}
