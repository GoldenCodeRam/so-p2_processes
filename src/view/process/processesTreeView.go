package process

import (
	"github.com/goldencoderam/so-p2_processes/src/object"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	PROCESS_NAME = iota
	PROCESS_TIME
    PROCESS_DELETED
	PROCESS_BLOCKED
	PROCESS_SUSPENDED_RUNNING
	PROCESS_SUSPENDED_BLOCKED
	PROCESS_TIME_REMAINING
	PROCESS_STATUS
)

type ProcessTreeView struct {
	TreeView  *gtk.TreeView
	listStore *gtk.ListStore
}

func NewTreeView() *ProcessTreeView {
	treeView, listStore := setupTreeView()
	return &ProcessTreeView{
		TreeView:  treeView,
		listStore: listStore,
	}
}

func (p *ProcessTreeView) AddRow(process *object.Process) {
	iter := p.listStore.Append()
	p.listStore.Set(
		iter,
		[]int{
            PROCESS_NAME,
            PROCESS_TIME,
            PROCESS_BLOCKED,
            PROCESS_SUSPENDED_RUNNING,
            PROCESS_SUSPENDED_BLOCKED,
            PROCESS_TIME_REMAINING,
            PROCESS_STATUS,
        },
		[]interface{}{
            process.Name,
            process.Time,
            process.IsBlocked,
            process.IsSuspendedAtRunning,
            process.IsSuspendedAtBlocked,
            process.GetTimeRemaining(),
            process.State,
        },
	)
}

func (p *ProcessTreeView) Clear() {
	p.listStore.Clear()
}

func (p *ProcessTreeView) RemoveRow(process *object.Process) {
	p.listStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		value, _ := model.GetValue(iter, PROCESS_NAME)
		valueString, _ := value.GetString()

		if valueString == process.Name {
			p.listStore.Remove(iter)
		}
		return false
	})
}

func setupTreeView() (*gtk.TreeView, *gtk.ListStore) {
	treeView, _ := gtk.TreeViewNew()

	treeView.AppendColumn(createColumn("Name", PROCESS_NAME))
	treeView.AppendColumn(createColumn("Time", PROCESS_TIME))
	treeView.AppendColumn(createColumn("Blocked", PROCESS_BLOCKED))
	treeView.AppendColumn(createColumn("Suspended at running", PROCESS_SUSPENDED_RUNNING))
	treeView.AppendColumn(createColumn("Suspended at blocked", PROCESS_SUSPENDED_BLOCKED))
	treeView.AppendColumn(createColumn("Time remaining", PROCESS_TIME_REMAINING))
	treeView.AppendColumn(createColumn("Status", PROCESS_STATUS))

	listStore, _ := gtk.ListStoreNew(
        glib.TYPE_STRING,
        glib.TYPE_INT,
        glib.TYPE_BOOLEAN,
        glib.TYPE_BOOLEAN,
        glib.TYPE_BOOLEAN,
        glib.TYPE_INT,
        glib.TYPE_INT,
    )
	treeView.SetModel(listStore)

	return treeView, listStore
}

func createColumn(columnTitle string, id int) *gtk.TreeViewColumn {
	cellRenderer, _ := gtk.CellRendererTextNew()
	column, _ := gtk.TreeViewColumnNewWithAttribute(columnTitle, cellRenderer, "text", id)
	return column
}
