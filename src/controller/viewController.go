package controller

import (
	"log"
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

func (v *viewController) DispatchProcess(process *object.Process) {
	v.MainWindow.RemoveFromReadyProcessesList(process)
	v.MainWindow.AddToDispatchedProcessesList(process)
}

func (v *viewController) Timeout(process *object.Process) {
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
	view.ShowInfoDialog("Finished")
}

func (v *viewController) CommunicateWithProcess(process *object.Process) {
	if process.CommunicateWith != "" && process.CommunicateWith != "None" && process.CommunicateWith != process.Name {
		for _, element := range GetMainControllerInstance().Processor.ReadyProcessesList {
			if element.State == object.READY && element.Name == process.CommunicateWith {
				v.MainWindow.LogCommunication(process.Name + " communicated with " + process.CommunicateWith)
			}
		}
	}
	process.CommunicateWith = "None"
}

func (v *viewController) AddProcessButtonListener(process *object.Process) {
	if GetMainControllerInstance().AddProcessToProcessor(process) {
		v.MainWindow.AddToReadyProcessesList(process)
	}
}

func (v *viewController) OnCommunicateWithProcessChanged(processName string) {
	log.Default().Println(processName)
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
