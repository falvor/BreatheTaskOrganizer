package test

import (
	"gui/gui"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskListAdd(t *testing.T) {
	list := &gui.TaskList{}
	task := gui.NewTask("Test", "Content")
	task1 := gui.NewTask("Test2", "Content")
	list.AddTask(&task)
	list.AddTask(&task1)
	assert.True(t,list.List[0].Equals(&task))
	assert.True(t,list.List[1].Equals(&task1))
}

func TestTaskListUpdateActiveTask(t *testing.T) {
	list := &gui.TaskList{}
	list.RandomTestTasks(3)
	assert.True(t,list.Active() == nil)
	list.SetActive(2)
	assert.True(t, list.Active().Equals(list.List[2]))
	newContent := "Yo this is new content."
	list.UpdateActiveTask(newContent)
	assert.Equal(t,list.Active().Note.Content, newContent)
	
}