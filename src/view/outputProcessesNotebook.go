package view

import "github.com/gotk3/gotk3/gtk"

type OutputProcessesNotebook struct {
	Box      *gtk.Box
	Notebook *gtk.Notebook

	readyProcessesTextView      *gtk.TextView
	dispatchedProcessesTextView *gtk.TextView
	processedProcessesTextView  *gtk.TextView
	blockedProcessesTextView    *gtk.TextView
	awokenProcessesTextView     *gtk.TextView
	resumedProcessesTextView    *gtk.TextView
	suspendedProcessesTextView  *gtk.TextView
	destroyedProcessesTextView  *gtk.TextView
}

func CreateOutputProcessesNotebook(listeners OutputProcessesNotebookListeners) *OutputProcessesNotebook {
	outputNotebook := OutputProcessesNotebook{
		Box:      CreateBox(gtk.ORIENTATION_VERTICAL, ZeroMargin),
		Notebook: CreateNotebook(),

		readyProcessesTextView:      CreateTextView(),
		dispatchedProcessesTextView: CreateTextView(),
		processedProcessesTextView:  CreateTextView(),
		blockedProcessesTextView:    CreateTextView(),
		awokenProcessesTextView:     CreateTextView(),
		resumedProcessesTextView:    CreateTextView(),
		suspendedProcessesTextView:  CreateTextView(),
		destroyedProcessesTextView:  CreateTextView(),
	}

	headerBar := CreateHeaderBar()

	outputNotebook.Notebook.AppendPage(outputNotebook.readyProcessesTextView, CreateLabel("Ready"))
	outputNotebook.Notebook.AppendPage(outputNotebook.dispatchedProcessesTextView, CreateLabel("Dispatched"))
	outputNotebook.Notebook.AppendPage(outputNotebook.processedProcessesTextView, CreateLabel("Processed"))
	outputNotebook.Notebook.AppendPage(outputNotebook.blockedProcessesTextView, CreateLabel("Blocked"))
	outputNotebook.Notebook.AppendPage(outputNotebook.awokenProcessesTextView, CreateLabel("Awoken"))
	outputNotebook.Notebook.AppendPage(outputNotebook.resumedProcessesTextView, CreateLabel("Resumed"))
	outputNotebook.Notebook.AppendPage(outputNotebook.suspendedProcessesTextView, CreateLabel("Suspended"))
	outputNotebook.Notebook.AppendPage(outputNotebook.destroyedProcessesTextView, CreateLabel("Destroyed"))

	headerBar.SetTitle("Output processes")
	outputNotebook.Box.Add(headerBar)
	outputNotebook.Box.PackEnd(outputNotebook.Notebook, true, true, 0)

	return &outputNotebook
}

func (o *OutputProcessesNotebook) ResetTextViews() {
    o.readyProcessesTextView.SetBuffer(nil)
    o.dispatchedProcessesTextView.SetBuffer(nil)
    o.processedProcessesTextView.SetBuffer(nil)
    o.blockedProcessesTextView.SetBuffer(nil)
    o.awokenProcessesTextView.SetBuffer(nil)
    o.resumedProcessesTextView.SetBuffer(nil)
    o.suspendedProcessesTextView.SetBuffer(nil)
    o.destroyedProcessesTextView.SetBuffer(nil)
}
