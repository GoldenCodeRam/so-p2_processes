package view

import "github.com/gotk3/gotk3/gtk"

func CreateProcessFrame(label string) *gtk.Frame {
	frame := CreateFrame(label)
	box := CreateBox(gtk.ORIENTATION_HORIZONTAL)
	grid := CreateGrid()
	processNameLabel := CreateLabel("Name")
	processTimeLabel := CreateLabel("Time")
	processName := CreateEntry()
	processTime := CreateEntry()
	processBlocked := CreateCheckButton("Is blocked?")
    addProcessButton := CreateButton("Create")

	grid.Attach(processNameLabel, 0, 0, 1, 1)
	grid.Attach(processTimeLabel, 0, 1, 1, 1)
	grid.Attach(processName, 1, 0, 1, 1)
	grid.Attach(processTime, 1, 1, 1, 1)
	grid.Attach(processBlocked, 0, 2, 2, 1)
    grid.Attach(addProcessButton, 0, 3, 2, 1)

	box.SetCenterWidget(grid)
	frame.Add(box)

	return frame
}

func CreateLabeledEntry(text string) *gtk.Box {
	box := CreateBox(gtk.ORIENTATION_HORIZONTAL)
	label := CreateLabel(text)
	entry := CreateEntry()

	box.Add(label)
	box.Add(entry)

	return box
}
