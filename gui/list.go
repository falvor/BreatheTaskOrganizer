// Holds a list of note items
package gui

import (
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// A view for all the available apex tasks
type TaskList struct {
	List []*Task
	vbox *fyne.Container
	backbutton *widget.Button
	Content *widget.Entry
	cardholder *fyne.Container
	active *Task
}

// Return the active task
func (list *TaskList) Active() *Task {
	return list.active
}

// Given either a index or a task pointer, set the active task. Return false if task doesn't
// exist within the list or the index is out of bounds
func (list *TaskList) SetActive(a interface{}) bool {
	switch v := a.(type) {
	case int:
		if v < 0 || v >= len(list.List) {
			fmt.Println("[TASK LIST] Invalid index supplied to update active task:",v)
			return false
		}
		list.active = list.List[v]
	case *Task:
		found := false
		for _,t := range list.List {
			if t.Equals(v) {
				found = true
				list.active = t
				if list.backbutton != nil {
					list.vbox.Remove(list.backbutton)
					list.backbutton = nil
				}
			}
		}
		if !found {
			// is child task
			fmt.Println("Setting a child task as the active task.")
			list.active = v
		}
	default:
		fmt.Println("[TASK LIST] Improper type supplied to set active task, needs index or task pointer.",reflect.TypeOf(v))
		return false
	}
	// fmt.Println("[TASK LIST] New active task:",list.active.Note.Title())
	return true
}

// Return the button that returns active to the parent of the current task
func (list *TaskList) ToParentButton(child *Task) *widget.Button {
	list.backbutton = widget.NewButtonWithIcon(child.Parent.Note.Title(),theme.NavigateBackIcon(),
		func(){
			if child.Parent == nil {
				panic("Error child has no parent? @ child = "+child.Note.Title())
			}
			list.ResetList()
			list.SetActive(child.Parent)
		})
	return list.backbutton
}

// This is called when a user clicks a child task. Now the child task populates the content window, and its children populate the cards below.
func (list *TaskList) SetChildActive(task *Task) {
	list.active = task
	if list.backbutton == nil {
		list.vbox.Add(list.ToParentButton(task))
	} else {
		fmt.Println("Back to parent button exists?",list.backbutton)
	}
}

func (list *TaskList) AddTask(t *Task) {
	list.List = append(list.List, t)
}

// Create and add a number of test tasks
func (list *TaskList) RandomTestTasks(n int) {
	for i := 0; i < n; i++ {
		title := fmt.Sprintf("Test %v", i)
		task := NewTask(title,"Content")
		list.AddTask(&task)
	}
}

// Set the current Content
func (list *TaskList) SetContent(content *widget.Entry) {
	list.Content = content
}

// Takes an array of tasks and builds cards for them, then appends them to
// bottom of the task entry window
func (list *TaskList) UpdateContentChildren(children []*Task) {
	list.cardholder.Objects = nil
	for _,v := range children {
		c := v.AsButton(list.ListButtonFunc(v, true))
		list.cardholder.Add(c)
	}
}

func (list *TaskList) UpdateActiveTask(newContent string) {
	// fmt.Println("Updating Content:", newContent)
	if list.active == nil {return}
	list.active.UpdateNote(newContent)
	list.UpdateContentChildren(list.active.Children)
}

// Return the task at index i
func (list *TaskList) GetTask(i int) *Task {
	return list.List[i]
}

// Takes the given child task, finds its parents, then sets the vbox options to the other children
func (list *TaskList) SetListChildren(task *Task) {
	list.vbox.Objects = nil
	list.backbutton = nil
	parent := task.Parent
	for _,v := range parent.Children {
		if !v.Equals(task) {
			button := v.AsButton(list.ListButtonFunc(v, true))
			list.vbox.Add(button)
		}
	}
}

func (list *TaskList) ListButtonFunc(task *Task, isChild bool) func() {
	return func() {
		// fmt.Println("Task button fired:",i, task.Note.Title())
		if isChild {
			list.SetListChildren(task)
			list.SetChildActive(task)
			// Set vbox options to the other children
		} else {
			list.SetActive(task)
		}
		list.Content.SetText(task.Note.GetContent())
	}
}

// Update the active task with the changed content
func (list *TaskList) ContentChanged(newContent string) {
	// Check if the title is the same as the active task, and the content is different, then update | list.active.HasTitle(newContent) && 
	if !list.active.HasContent(newContent) {
		list.UpdateActiveTask(newContent)
	}
}

// Builds the vbox with all the parent tasks 
func (list *TaskList) ResetList() {
	if list.vbox == nil {
		list.vbox = fyne.NewContainerWithLayout(layout.NewVBoxLayout())
	} else {
		list.vbox.Objects = nil
	}
	for _,v := range list.List {
		button := v.AsButton(list.ListButtonFunc(v, false))
		list.vbox.Add(button)
		for _,c := range v.Children {
			c.SetParent(v)
		}
	}
}

// Return the canvas object that is drawn on the gui
func (list *TaskList) LoadList() fyne.CanvasObject {
	list.cardholder = fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(80,100)))
	list.ResetList()
	fmt.Println("Task List loaded.")
	return list.vbox
}