package test

import (
	"fmt"
	"gui/gui"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskNote(t *testing.T) {
	note := gui.NewNote("Test", "Content")
	task := &gui.Task{Note: &note}
	assert.True(t, task.Note.Equals(note))
}

func TestTaskUpdateNote(t *testing.T) {
	task := gui.NewTask("Test","")
	s := "Content"
	task.UpdateNote(s)
	assert.Equal(t, task.Note.Content, s)
}

func TestTaskParent(t *testing.T) {
	parent := gui.NewTask("Parent", "Test")
	child := gui.NewTask("Child","Test", &parent)

	assert.True(t, child.Parent.Note.Equals(*parent.Note))
	assert.True(t, child.Note.Equals(*parent.Children[0].Note))
}

func TestTaskAddChild(t *testing.T) {
	parent := gui.NewTask("Parent", "Test")
	children := make([]*gui.Task,0)
	for i := 0; i < 7; i++ {
		title := fmt.Sprintf("Child%v",i)
		child := gui.NewTask(title,"Content")
		parent.AddChild(&child)
		children = append(children, &child)
	}
	for i,v := range parent.Children {
		assert.True(t, children[i].Equals(v))
	}
}

