package controller

import (
	"github.com/goldencoderam/so-p2_processes/src/object"
	"github.com/goldencoderam/so-p2_processes/src/view"
	"github.com/gotk3/gotk3/gtk"
)

type viewController struct {
	MainWindow *view.MainWindow
}

func (v *viewController) AddProcessButtonListener(process *object.Process) {
	GetMainControllerInstance().AddProcessToProcessor(process)
	v.MainWindow.SetReadyProcessesListText(
		GetMainControllerInstance().Processor.ReadyProcessesLog,
	)
}

func (v *viewController) StartProcessor() {
	GetMainControllerInstance().Processor.MakeTick()
}

func (v *viewController) ResetProcessor() {
	GetMainControllerInstance().Processor.Reset()
	v.MainWindow.ResetLogs()
}

var viewControllerInstance *viewController

func GetViewControllerInstance() *viewController {
	if viewControllerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if viewControllerInstance == nil {
			viewControllerInstance = &viewController{}
			generateMainWindow()
		}
	}
	return viewControllerInstance
}

func generateMainWindow() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	viewControllerInstance.MainWindow = view.CreateMainWindow(viewControllerInstance)
	// Set the default window size.
	viewControllerInstance.MainWindow.Window.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	viewControllerInstance.MainWindow.Window.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
