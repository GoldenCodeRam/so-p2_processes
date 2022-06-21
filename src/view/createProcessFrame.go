package view

import (
	"strconv"

	"github.com/goldencoderam/so-p2_processes/src/object"
	"github.com/gotk3/gotk3/gtk"
)

type ProcessFrame struct {
	Frame *gtk.Frame

	ProcessNameEntry                            *gtk.Entry
	ProcessTimeEntry                            *gtk.Entry
	IsProcessBlockedCheckButton                 *gtk.CheckButton
	IsProcessSuspendedAtRunningCheckButton      *gtk.CheckButton
	IsProcessSuspendedAtBlockedCheckButton      *gtk.CheckButton
	IsProcessSuspendedAtIOCompletionCheckButton *gtk.CheckButton
}

func CreateProcessFrame(label string, listeners ProcessFrameListeners) *ProcessFrame {
	processFrame := ProcessFrame{
		Frame:                                       CreateFrame(label),
		ProcessNameEntry:                            CreateEntry(),
		ProcessTimeEntry:                            CreateEntry(),
		IsProcessBlockedCheckButton:                 CreateCheckButton("¿Se bloquea?"),
		IsProcessSuspendedAtRunningCheckButton:      CreateCheckButton("¿Se suspende en ejecución?"),
		IsProcessSuspendedAtBlockedCheckButton:      CreateCheckButton("¿Se suspende bloqueado?"),
		IsProcessSuspendedAtIOCompletionCheckButton: CreateCheckButton("¿Se desbloquea suspendido?"),
	}
	box := CreateBox(gtk.ORIENTATION_HORIZONTAL, SmallMargin)
	grid := CreateGrid()

	processNameLabel := CreateLabel("Nombre")
	processTimeLabel := CreateLabel("Tiempo")

	addProcessButton := CreateButton("Crear")
	addProcessButton.Connect("clicked", func() {
		processFrame.addProcess(listeners)
	})

	grid.Attach(processNameLabel, 0, 0, 1, 1)
	grid.Attach(processTimeLabel, 0, 1, 1, 1)
	grid.Attach(processFrame.ProcessNameEntry, 1, 0, 1, 1)
	grid.Attach(processFrame.ProcessTimeEntry, 1, 1, 1, 1)
	grid.Attach(processFrame.IsProcessBlockedCheckButton, 0, 2, 2, 1)
	grid.Attach(processFrame.IsProcessSuspendedAtRunningCheckButton, 0, 3, 2, 1)
	grid.Attach(processFrame.IsProcessSuspendedAtBlockedCheckButton, 0, 4, 2, 1)
	grid.Attach(processFrame.IsProcessSuspendedAtIOCompletionCheckButton, 0, 5, 2, 1)

	grid.Attach(addProcessButton, 0, 6, 2, 1)

	box.SetCenterWidget(grid)
	processFrame.Frame.Add(box)

	return &processFrame
}

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
		Name:                      name,
		Time:                      time,
		IsBlocked:                 p.IsProcessBlockedCheckButton.GetActive() || p.IsProcessSuspendedAtBlockedCheckButton.GetActive() || p.IsProcessSuspendedAtIOCompletionCheckButton.GetActive(),
		IsSuspendedAtRunning:      p.IsProcessSuspendedAtRunningCheckButton.GetActive(),
		IsSuspendedAtBlocked:      p.IsProcessSuspendedAtBlockedCheckButton.GetActive() || p.IsProcessSuspendedAtIOCompletionCheckButton.GetActive(),
		IsSuspendedAtIOCompletion: p.IsProcessSuspendedAtIOCompletionCheckButton.GetActive(),
		State:                     object.READY,
		TimeRemaining:             time,
	})
	p.resetFields()
}

func (p *ProcessFrame) resetFields() {
	p.ProcessNameEntry.SetText("")
	p.ProcessTimeEntry.SetText("")
	p.IsProcessBlockedCheckButton.SetActive(false)
	p.IsProcessSuspendedAtRunningCheckButton.SetActive(false)
	p.IsProcessSuspendedAtBlockedCheckButton.SetActive(false)
	p.IsProcessSuspendedAtIOCompletionCheckButton.SetActive(false)
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
