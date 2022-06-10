package view

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func GetMainWindow() *gtk.Window {
    win := CreateWindow()

    paned, err := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)
    win.Add(paned)


    box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
    frame := CreateProcessFrame("test")

	l1, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

    box.SetCenterWidget(frame)

	// Add the label to the window.
    paned.Pack1(box, true, false)
    paned.Pack2(l1, true, true)

	return win
}
