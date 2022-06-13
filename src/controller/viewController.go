package controller

import (
	"log"

	"github.com/goldencoderam/so-p2_processes/src/object"
	"github.com/goldencoderam/so-p2_processes/src/view"
	"github.com/gotk3/gotk3/gtk"
)

type viewController struct {
	MainWindow *view.MainWindow
}

func (v *viewController) DispatchProcess(process *object.Process) {
	v.MainWindow.RemoveFromReadyProcessesList(process)
    v.MainWindow.AddToDispatchedProcessesList(process)
}

func (v *viewController) TimerRunoutProcess(process *object.Process) {
	GetMainControllerInstance().AddProcessToProcessor(process)
	v.MainWindow.AddToReadyProcessesList(process)
}

func (v *viewController) FinishedProcess(process *object.Process) {
    v.MainWindow.AddToProcessedProcessesList(process)
}

func (v *viewController) BlockedProcess(process *object.Process) {
    v.MainWindow.AddToBlockedProcessesList(process)
}

func (v *viewController) SuspendedReadyProcess(process *object.Process) {
    v.MainWindow.RemoveFromReadyProcessesList(process)
    v.MainWindow.AddToSuspendedReadyProcessesList(process)
}

func (v *viewController) SuspendedBlockedProcess(process *object.Process) {
    v.MainWindow.AddToSuspendedBlockedProcessesList(process)
}

func (v *viewController) DestroyedProcess(process *object.Process) {
    v.MainWindow.AddToDestroyedProcessesList(process)
}

func (v *viewController) FinishedProcessing() {
    // TODO: This should be done soon
    log.Default().Println("Finished")
}

func (v *viewController) AddProcessButtonListener(process *object.Process) {
	GetMainControllerInstance().AddProcessToProcessor(process)
	v.MainWindow.AddToReadyProcessesList(process)
}

func (v *viewController) StartProcessor() {
    controllerInstance := GetMainControllerInstance()
    for len(controllerInstance.Processor.ReadyProcessesList) > 0 || controllerInstance.Processor.CurrentProcess != nil {
        controllerInstance.Processor.MakeTick(v)
    }
}

func (v *viewController) MakeProcessorTick() {
	GetMainControllerInstance().Processor.MakeTick(v)
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
