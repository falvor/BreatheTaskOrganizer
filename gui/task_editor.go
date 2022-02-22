package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TaskEditor struct {
	// Task Name, Edit Task Name(icon)
	TaskName *widget.Label
	TaskNameEntry *widget.Entry
	NameEdit *widget.Button
	// Entry
	Entry *widget.Entry
	// Save Task (icon)
	Save *widget.Button
	Delete *widget.Button
	Check *widget.Button
	// New Child (icon) [toggles form]
	NewChild *widget.Button // Shows the form and then becomes a save button
	UndoChild *widget.Button // Shows up after new child pressed, closes the new child entry
	ChildForm *widget.Entry
}

func (editor *TaskEditor) SetCurrentTask(task *Task) {
	editor.TaskName = widget.NewLabel(task.Title())
	editor.Entry = widget.NewMultiLineEntry()
	editor.Entry.Text = task.Content()
}

// Builds the button that toggles editing of the current task name
func (editor *TaskEditor) BuildEditTaskNameButton() *widget.Button {
	editor.NameEdit = widget.NewButtonWithIcon("",theme.CancelIcon(), func() {
		if editor.TaskName.Hidden {
			editor.TaskName.Show()
			editor.NameEdit.Hide()
		} else {
			editor.TaskName.Hide()
			editor.NameEdit.Show()
		}
	})
	return editor.NameEdit
}

func (editor *TaskEditor) SaveCurrentTask(old *Task) (task Task) {
	if old == nil {
		// New Task
		task = NewTask(editor.TaskName.Text, editor.Entry.Text)
	} else {
		// Replace old task
		task = NewTask(editor.TaskName.Text, editor.Entry.Text, old.Parent)
		task.Children = old.Children
	}
	return 
}

// Get the canvas object that draws on the gui, assumes objects have been built already
func (editor *TaskEditor) LoadTaskEditor() fyne.CanvasObject {
	top_buttons := container.NewHBox(editor.TaskName, editor.TaskNameEntry, editor.NameEdit)
	entry_final := container.NewMax(editor.Entry)
	side_buttons := container.NewVBox(editor.Save, editor.Delete, editor.Check)
	bot_buttons := container.NewHBox(editor.NewChild, editor.UndoChild, editor.ChildForm)
	return container.NewBorder(
		top_buttons,
		bot_buttons,
		nil,
		side_buttons, entry_final)
}