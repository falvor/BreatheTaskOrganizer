package gui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
)

// Errors

// Panic after printing the string
func taskFatal(a ...string) {
	s := a[0]
	for i,v := range a {
		msg := fmt.Sprintf("%v \t %v", i, v)
		s += msg + "\n"
	}
	panic("Task Fatal Error.")
}

// Task Object
type Task struct {
	Children []*Task
	Parent   *Task
	Note     *Note
}

func NewTask(title, content string, parent ...*Task) (t Task) {
	if len(parent) > 1 {
		fmt.Println("Wanted up to one parent for task, instead recieved two. Empty task returned")
		return
	} else if len(parent) == 1 {
		t.Parent = parent[0]
		parent[0].AddChild(&t)
	}
	note := NewNote(title, content)
	fmt.Println("Note created.",note.Title())
	t.Note = &note
	return
}

// Return a number of random tasks
func TestTasks(num int) (t []*Task) {
	for i := 0; i < num; i++ {
		title := fmt.Sprintf("Test Task %v",i)
		task := NewTask(title, "Random Test")
		t = append(t, &task)
	}
	return t
}

// Returns the task as a button
func (t *Task) AsButton(buttonFunc func()) (button *widget.Button) {
	button = widget.NewButton(
		t.Note.Title(), buttonFunc)
	return
}

// See if two tasks are equal (have the same title)
func (t *Task) Equals(other *Task) bool {
	return t.Note.Title() == other.Note.Title()
}

// Replace the content of this tasks note when the given content
func (t *Task) UpdateNote(content string) {
	t.Note.UpdateContent(content)
}

// Return true if child exists within the task's children
func (t *Task) HasChild(child *Task) bool {
	for _,v := range t.Children {
		if v.Equals(child) {
			return true
		}
	}
	return false
}

// Return true if note has the same title as the passed content
func (t *Task) HasTitle(content string) bool {
	return t.Note.Title() == strings.SplitN(content, "\n", 2)[0]
}

// Returns the title of the underlying note
func (t *Task) Title() string {
	return t.Note.Title()
}

// Returns the title of the underlying note
func (t *Task) Content() string {
	return t.Note.GetContent()
}



// Return true if passed content equals the content of the task's note
func (t *Task) HasContent(content string) bool {
	return t.Note.Content == content
}

// Append a new child task to the current task
func (t *Task) AddChild(child *Task) {
	if t.HasChild(child) {return}
	t.Children = append(t.Children, child)
}

// Set a task as a new Parent
func (t *Task) SetParent(parent *Task) {
	t.Parent = parent
}

// Remove a current child task, panic if task is not found (this should never happen)
func (t * Task) RemoveChild(child *Task) {
	new_children := make([]*Task,0)
	removed := false
	for _,v := range t.Children {
		if v.Equals(child) {
			removed = true
		} else {
			new_children = append(new_children, v)
		}
	}
	if !removed{
		taskFatal("Trying to remove a child that doesn't exist (adult child):", t.Note.Title(), child.Note.Title())
	}
}