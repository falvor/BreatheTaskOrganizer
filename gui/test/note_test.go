package test

import (
	"gui/gui"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoteTitle(t *testing.T) {
	note := &gui.Note{Content: "Hello"}
	assert.Equal(t, "Hello", note.Title())
	
	note = &gui.Note{Content: "Note\nLine2"}
	assert.Equal(t, "Note", note.Title())
	
	note = &gui.Note{Content: ""}
	assert.Equal(t, "Empty", note.Title())
}

func TestNoteGetContent(t *testing.T) {
	note := &gui.Note{Content:"Note\nTest Content"}
	assert.Equal(t, note.GetContent(), "Test Content")
}

func TestNoteUpdateContent(t *testing.T) {
	note := &gui.Note{Content: "Note\nTest Content"}
	note.UpdateContent("New Content")
	assert.Equal(t, note.GetContent(), "New Content")
}

func TestNoteEquals(t *testing.T) {
	note1 := gui.NewNote("Test", "Content")
	note2 := gui.NewNote("Test", "Content")

	b := note1.Equals(note2)
	assert.True(t,b)
}