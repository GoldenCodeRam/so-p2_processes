package view

import "github.com/gotk3/gotk3/gtk"

func CreateFrame(label string) *gtk.Frame {
	frame, err := gtk.FrameNew(label)
	if err != nil {
		panic(err)
	}

	frame.SetMarginStart(SmallMargin)
	frame.SetMarginEnd(SmallMargin)

	return frame
}

func CreateCheckButton(label string) *gtk.CheckButton {
	var checkButton *gtk.CheckButton
    var err error

	if label == "" {
		checkButton, err = gtk.CheckButtonNew()
	} else {
		checkButton, err = gtk.CheckButtonNewWithLabel(label)
	}

	if err != nil {
		panic(err)
	}

	return checkButton
}

func CreateButton(label string) *gtk.Button {
    button, err := gtk.ButtonNewWithLabel(label)
    if err != nil {
        panic(err)
    }

    return button
}

func CreateToggleButton(label string) *gtk.ToggleButton {
	toggleButton, err := gtk.ToggleButtonNewWithLabel(label)
	if err != nil {
		panic(err)
	}

	return toggleButton
}

func CreateLabel(text string) *gtk.Label {
	label, err := gtk.LabelNew(text)
	if err != nil {
		panic(err)
	}

	label.SetMarginStart(SmallMargin)
	label.SetMarginEnd(SmallMargin)

	return label
}

func CreateEntry() *gtk.Entry {
	entry, err := gtk.EntryNew()
	if err != nil {
		panic(err)
	}
	return entry
}

func CreateBox(orientation gtk.Orientation) *gtk.Box {
	box, err := gtk.BoxNew(orientation, SmallMargin)
	if err != nil {
		panic(err)
	}

	box.SetMarginStart(SmallMargin)
	box.SetMarginEnd(SmallMargin)
	box.SetMarginTop(SmallMargin)
	box.SetMarginBottom(SmallMargin)

	return box
}

func CreateGrid() *gtk.Grid {
	grid, err := gtk.GridNew()
	if err != nil {
		panic(err)
	}

	grid.SetMarginStart(SmallMargin)
	grid.SetMarginEnd(SmallMargin)
	grid.SetMarginTop(SmallMargin)
	grid.SetMarginBottom(SmallMargin)

	grid.SetRowSpacing(uint(SmallMargin))
	grid.SetColumnSpacing(uint(SmallMargin))

	return grid
}

func CreateWindow() *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		panic(err)
	}

	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	return win
}
