package gui

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// This is how the user interacts with the tasks and their children
type UI struct {
	Current *Task
	Content *widget.Entry
	List *TaskList
}

func(ui *UI) LoadUI() fyne.CanvasObject{

	ui.Content = widget.NewMultiLineEntry()

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {fmt.Println("Tool bar item tapped.")}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {fmt.Println("Tool bar item tapped.")}),
	)

	side := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(bar, nil, nil, nil), bar, ui.List.LoadList())

	ui.List.SetContent(ui.Content)	
	ui.Content.PlaceHolder = "type out a task here."
	ui.Content.Wrapping = fyne.TextWrapBreak
	ui.Content.OnChanged = ui.List.ContentChanged
	entry_final := container.NewMax(ui.Content)
	task_bot_view := container.NewBorder(
		NewFillObject("Edit Task Name"),
		NewFillObject("New Child"),
		NewFillObject("Edit Task"),
		NewFillObject("Set Completed"), entry_final)
	task_view := container.NewVSplit(task_bot_view, ui.List.cardholder)
	// Content updater
	go func() {
		// pastContent := ui.Content.Text
		// for {
		// 	if pastContent != ui.Content.Text {
		// 		ui.Content.Refresh()
		// 		pastContent = ui.Content.Text
		// 		ui.List.UpdateActiveTask(pastContent)
		// 	}
		// }
	}()

	split := container.NewHSplit(side, task_view)
	return split
}
