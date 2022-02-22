package main

import (
	"fmt"
	"gui/gui"
	// "gui/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/widget"
)

func uMsg(s string) {
	fmt.Println(s)
}

var current *gui.Note


func main() {
	fmt.Println("Initializing...")
	a := app.New()
	w := a.NewWindow("Breathe Task Organizer")
	
	w.Resize(fyne.NewSize(500,600))
	ui := gui.UI{}
	list := gui.TaskList{}
	t1 := gui.NewTask(
		"Test 1",
		"This is a regular old test.", 
	)
	t2 := gui.NewTask(
		"Test 2",
		"Lorem Ipsum Fo, Random noise. Random words.", 
	)
	t3 := gui.NewTask(
		"Test 3",
		"This is also a wordy, but necessary test.",
	)
	t1.Children = gui.TestTasks(3)
	t2.Children = gui.TestTasks(5)
	t3.Children = gui.TestTasks(9)
	list.AddTask(&t1)
	list.AddTask(&t2)
	list.AddTask(&t3)
	ui.List = &list
	w.SetContent(ui.LoadUI())

	w.ShowAndRun()
}

