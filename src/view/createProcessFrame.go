package view

import (
	"strconv"

	"github.com/goldencoderam/so-p2_processes/src/object"
	"github.com/gotk3/gotk3/gtk"
)

type ProcessFrame struct {
	Frame *gtk.Frame
}

func CreateProcessFrame(label string, listeners ProcessFrameListeners) *ProcessFrame {
	processFrame := ProcessFrame{
		Frame: CreateFrame(label),
	}
	box := CreateBox(gtk.ORIENTATION_HORIZONTAL, SmallMargin)
	grid := CreateGrid()
	processNameLabel := CreateLabel("Name")
	processTimeLabel := CreateLabel("Time")
	processName := CreateEntry()
	processTime := CreateEntry()
	processDeleted := CreateCheckButton("Is deleted?")
	processBlocked := CreateCheckButton("Is blocked?")
	processSuspendedAtReady := CreateCheckButton("Is suspended at ready?")
	processSuspendedAtRunning := CreateCheckButton("Is suspended at running?")
	processSuspendedAtBlocked := CreateCheckButton("Is suspended at blocked?")
	addProcessButton := CreateButton("Create")

	addProcessButton.Connect("clicked", func() {
		name, _ := processName.GetText()
		timeText, _ := processTime.GetText()

		time, err := strconv.Atoi(timeText)
		if err != nil {
			ShowErrorDialog(err)
		} else {
			listeners.AddProcessButtonListener(&object.Process{
				Name:                 name,
				Time:                 time,
				IsDeleted:            processDeleted.GetActive(),
				IsBlocked:            processBlocked.GetActive(),
				IsSuspendedAtReady:   processSuspendedAtReady.GetActive(),
				IsSuspendedAtRunning: processSuspendedAtRunning.GetActive(),
				IsSuspendedAtBlocked: processSuspendedAtBlocked.GetActive(),
				TimeRemaining:        time,
			})
		}
	})

	grid.Attach(processNameLabel, 0, 0, 1, 1)
	grid.Attach(processTimeLabel, 0, 1, 1, 1)
	grid.Attach(processName, 1, 0, 1, 1)
	grid.Attach(processTime, 1, 1, 1, 1)
	grid.Attach(processDeleted, 0, 2, 2, 1)
	grid.Attach(processBlocked, 0, 3, 2, 1)
	grid.Attach(processSuspendedAtReady, 0, 4, 2, 1)
	grid.Attach(processSuspendedAtRunning, 0, 5, 2, 1)
	grid.Attach(processSuspendedAtBlocked, 0, 6, 2, 1)
	grid.Attach(addProcessButton, 0, 7, 2, 1)

	box.SetCenterWidget(grid)
	processFrame.Frame.Add(box)

	return &processFrame
}
