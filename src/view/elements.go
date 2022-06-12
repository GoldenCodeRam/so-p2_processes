package view

import "github.com/gotk3/gotk3/gtk"

func CreateFrame(label string) *gtk.Frame {
	frame, err := gtk.FrameNew(label)
	if err != nil {
		panic(err)
	}

	frame.SetMarginStart(int(SmallMargin))
	frame.SetMarginEnd(int(SmallMargin))

	return frame
}

func CreateTextBuffer(text string) *gtk.TextBuffer {
    textBuffer, err := gtk.TextBufferNew(nil)
    if err != nil {
        panic(err)
    }

    textBuffer.SetText(text)
    return textBuffer
}

func CreateNotebook() *gtk.Notebook {
	notebook, err := gtk.NotebookNew()
	if err != nil {
		panic(err)
	}

	notebook.SetScrollable(true)
	return notebook
}

func CreateTextView() *gtk.TextView {
	textView, err := gtk.TextViewNew()
	if err != nil {
		panic(err)
	}

	textView.SetEditable(false)
	return textView
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

	label.SetMarginStart(int(SmallMargin))
	label.SetMarginEnd(int(SmallMargin))

	return label
}

func CreateEntry() *gtk.Entry {
	entry, err := gtk.EntryNew()
	if err != nil {
		panic(err)
	}
	return entry
}

func CreateBox(orientation gtk.Orientation, margins MarginType) *gtk.Box {
	box, err := gtk.BoxNew(orientation, int(margins))
	if err != nil {
		panic(err)
	}

	box.SetMarginStart(int(margins))
	box.SetMarginEnd(int(margins))
	box.SetMarginTop(int(margins))
	box.SetMarginBottom(int(margins))

	return box
}

func CreateHeaderBar() *gtk.HeaderBar {
    header, err := gtk.HeaderBarNew()
    if err != nil {
        panic(err)
    }

    return header
}

func CreatePaned(orientation gtk.Orientation) *gtk.Paned {
	paned, err := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)
	if err != nil {
		panic(err)
	}
	return paned
}

func CreateGrid() *gtk.Grid {
	grid, err := gtk.GridNew()
	if err != nil {
		panic(err)
	}

	grid.SetMarginStart(int(SmallMargin))
	grid.SetMarginEnd(int(SmallMargin))
	grid.SetMarginTop(int(SmallMargin))
	grid.SetMarginBottom(int(SmallMargin))

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
