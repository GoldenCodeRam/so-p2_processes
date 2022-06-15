package view

import (
	"strconv"

	"github.com/goldencoderam/so-p2_processes/src/object"
	"github.com/gotk3/gotk3/gtk"
)

type ProcessFrame struct {
	Frame *gtk.Frame

	ProcessNameEntry                       *gtk.Entry
	ProcessTimeEntry                       *gtk.Entry
	IsProcessBlockedCheckButton            *gtk.CheckButton
	IsProcessSuspendedAtRunningCheckButton *gtk.CheckButton
	IsProcessSuspendedAtBlockedCheckButton *gtk.CheckButton
}

func CreateProcessFrame(label string, listeners ProcessFrameListeners) *ProcessFrame {
	processFrame := ProcessFrame{
		Frame:                                  CreateFrame(label),
		ProcessNameEntry:                       CreateEntry(),
		ProcessTimeEntry:                       CreateEntry(),
		IsProcessBlockedCheckButton:            CreateCheckButton("Is blocked?"),
		IsProcessSuspendedAtRunningCheckButton: CreateCheckButton("Is suspended at running?"),
		IsProcessSuspendedAtBlockedCheckButton: CreateCheckButton("Is suspended at blocked?"),
	}
	box := CreateBox(gtk.ORIENTATION_HORIZONTAL, SmallMargin)
	grid := CreateGrid()

	processPriorityLabel := CreateLabel("Priority")
	processChangePriorityLabel := CreateLabel("Change priority")
	processNameLabel := CreateLabel("Name")
	processTimeLabel := CreateLabel("Time")

	addProcessButton := CreateButton("Create")
	addProcessButton.Connect("clicked", func() {
		processFrame.addProcess(listeners)
	})

	grid.Attach(processPriorityLabel, 0, 0, 1, 1)
	grid.Attach(processChangePriorityLabel, 0, 1, 1, 1)
	grid.Attach(processNameLabel, 0, 2, 1, 1)
	grid.Attach(processTimeLabel, 0, 3, 1, 1)
	grid.Attach(processName, 1, 2, 1, 1)
	grid.Attach(processTime, 1, 3, 1, 1)
	grid.Attach(processDeleted, 0, 4, 2, 1)
	grid.Attach(processBlocked, 0, 5, 2, 1)
	grid.Attach(processSuspendedAtReady, 0, 6, 2, 1)
	grid.Attach(processSuspendedAtRunning, 0, 7, 2, 1)
	grid.Attach(processSuspendedAtBlocked, 0, 8, 2, 1)

	grid.Attach(CreateLabel("Communicate with"), 0, 9, 2, 1)
	grid.Attach(addProcessButton, 0, 11, 2, 1)

	box.SetCenterWidget(grid)
	processFrame.Frame.Add(box)

	return &processFrame
} 

func

func (p *ProcessFrame) addProcess(listeners ProcessFrameListeners) {
	name, err := extractTextFromEntry(p.ProcessNameEntry)
	if err != nil {
		return
	}

	time, err := extractIntFromEntry(p.ProcessTimeEntry)
	if err != nil {
		return
	}

	listeners.AddProcessButtonListener(&object.Process{
		Name:                 name,
		Time:                 time,
		IsBlocked:            p.IsProcessBlockedCheckButton.GetActive(),
		IsSuspendedAtRunning: p.IsProcessSuspendedAtRunningCheckButton.GetActive(),
		IsSuspendedAtBlocked: p.IsProcessSuspendedAtBlockedCheckButton.GetActive(),
		State:                object.READY,
		TimeRemaining:        time,
	})
	p.resetFields()
}

func (p *ProcessFrame) resetFields() {
	p.ProcessNameEntry.SetText("")
	p.ProcessTimeEntry.SetText("")
	p.IsProcessBlockedCheckButton.SetActive(false)
	p.IsProcessSuspendedAtRunningCheckButton.SetActive(false)
	p.IsProcessSuspendedAtBlockedCheckButton.SetActive(false)
}

func extractIntFromEntry(entry *gtk.Entry) (int, error) {
	text, err := extractTextFromEntry(entry)
	if err != nil {
		return 0, err
	} else {
		number, err := strconv.Atoi(text)
		if err != nil {
			return 0, err
		} else {
			return number, nil
		}
	}
}

func extractTextFromEntry(entry *gtk.Entry) (string, error) {
	text, err := entry.GetText()
	if err != nil {
		return "", err
	} else {
		return text, err
	}
}
