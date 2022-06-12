package view

import "github.com/gotk3/gotk3/gtk"

type ProcessActionPanel struct {

}

func CreateProcessActionPanel(listeners ProcessActionPanelListeners) *gtk.Box {
    actionPanelBox := CreateBox(gtk.ORIENTATION_VERTICAL, SmallMargin)
    centerBox := CreateBox(gtk.ORIENTATION_VERTICAL, ZeroMargin)

    startSimulationButton := CreateButton("Start processor")
    makeProcessorTickButton := CreateButton("Tick processor")
    resetSimulationButton := CreateButton("Reset processor")

    createProcessFrame := CreateProcessFrame("Create process", listeners)

    centerBox.SetSpacing(int(MediumMargin))
    startSimulationButton.Connect("clicked", listeners.StartProcessor)
    makeProcessorTickButton.Connect("clicked", listeners.MakeProcessorTick)
    resetSimulationButton.Connect("clicked", listeners.ResetProcessor)

    centerBox.Add(createProcessFrame.Frame)
    centerBox.Add(startSimulationButton)
    centerBox.Add(makeProcessorTickButton)
    centerBox.Add(resetSimulationButton)

    actionPanelBox.SetCenterWidget(centerBox)
    return actionPanelBox
}
