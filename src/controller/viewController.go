package controller

import (
	"sync"

	"github.com/goldencoderam/so-p2_processes/src/model"
	"github.com/goldencoderam/so-p2_processes/src/object"
	"github.com/goldencoderam/so-p2_processes/src/view"
	"github.com/gotk3/gotk3/gtk"
)

type viewController struct {
	MainWindow *view.MainWindow
	Processor  *model.Processor
}

var viewControllerInstance *viewController

var lock = &sync.Mutex{}

func GetViewControllerInstance() *viewController {
	if viewControllerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if viewControllerInstance == nil {
			viewControllerInstance = &viewController{}
            viewControllerInstance.Processor = &model.Processor{
                LogListeners: viewControllerInstance,
            }
			generateMainWindow()
		}
	}
	return viewControllerInstance
}

func (v *viewController) LogProcessDispatched(process *object.Process) {
    v.MainWindow.AddToDispatchTransitionList(process)
    v.MainWindow.AddToRunningProcessesList(process)
    v.MainWindow.RemoveFromReadyProcessesList(process)
}

func (v *viewController) LogProcessTimeout(process *object.Process) {
    v.MainWindow.AddToTimeoutTransitionList(process)
    v.MainWindow.AddToReadyProcessesList(process)
}

func (v *viewController) LogProcessBlocked(process *object.Process) {
    v.MainWindow.AddToWaitEventTransitionList(process)
    v.MainWindow.AddToBlockedProcessesList(process)
}

func (v *viewController) LogProcessSuspendedRunning(process *object.Process) {
    v.MainWindow.AddToSuspendRunningTransitionList(process)
    v.MainWindow.AddToSuspendedReadyProcessesList(process)
}

func (v *viewController) LogProcessFinished(process *object.Process) {
    v.MainWindow.AddToFinishedProcessesList(process)
}

func (v *viewController) LogProcessIOBlockedCompleted(process *object.Process) {
    v.MainWindow.AddToCompletionEventTransitionList(process)
    v.MainWindow.AddToReadyProcessesList(process)
}

func (v *viewController) LogProcessSuspendedBlocked(process *object.Process) {
    v.MainWindow.AddToSuspendBlockedTransitionList(process)
    v.MainWindow.AddToSuspendedBlockedProcessesList(process)
}

func (v *viewController) LogProcessSuspendedBlockedResumed(process *object.Process) {
    v.MainWindow.AddToResumeSuspendedBlockedTransitionList(process)
    v.MainWindow.AddToBlockedProcessesList(process)
}

func (v *viewController) LogProcessIOSuspendedBlockedCompleted(process *object.Process) {
    v.MainWindow.AddToCompletionEventSuspendedBlockedTransitionList(process)
    v.MainWindow.AddToSuspendedReadyProcessesList(process)
}

func (v *viewController) LogProcessSuspendedReadyResumed(process *object.Process) {
    v.MainWindow.AddToResumeSuspendedReadyTransitionList(process)
    v.MainWindow.AddToReadyProcessesList(process)
}

func (v *viewController) LogFinishedProcessing() {
    view.ShowInfoDialog("Finalizado")
}

func (v *viewController) AddProcessButtonListener(process *object.Process) {
    err := v.Processor.AddProcessToReadyList(process)
    if err != nil {
        view.ShowErrorDialog(err)
    } else {
        v.MainWindow.AddToReadyProcessesList(process)
    }
}

func (v *viewController) StartProcessor() {
	for len(v.Processor.ReadyProcessesList) > 0 || v.Processor.CurrentProcess != nil {
		v.Processor.MakeTick(v)
	}
	v.Processor.MakeTick(v)
}

func (v *viewController) MakeProcessorTick() {
	v.Processor.MakeTick(v)
}

func (v *viewController) ResetProcessor() {
	v.Processor.Reset()
	v.MainWindow.ResetLogs()
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
