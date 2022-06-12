package view

import "github.com/gotk3/gotk3/gtk"

func ShowErrorDialog(err error) {
	dialog := gtk.MessageDialogNew(
        nil,
        gtk.DIALOG_DESTROY_WITH_PARENT,
        gtk.MESSAGE_ERROR,
        gtk.BUTTONS_CLOSE,
        err.Error(),
    )
    dialog.Run()
    dialog.Destroy()
}
