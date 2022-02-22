package gui

import (
	"fmt"
	"strings"
)

type Note struct {
	Content string
}

func NewNote(title, content string) (n Note) {
	fmt.Println("Creating new note:",title, content)
	n.Content = title +"\n" + content
	return
}

// Return the content without the title of the note
func (n *Note) GetContent() string {
	if n.Content == "" {
		return ""
	}
	return strings.SplitN(n.Content, "\n", 2)[1] 
}

func (n *Note) UpdateContent(s string)  {
	n.Content = n.Title() + "\n" + s
}

func (n *Note) Title() string {
	if n.Content == "" {
		return "Empty"
	}
	return strings.SplitN(n.Content, "\n", 2)[0]
}

// Return true if the content of one note exactly matches that of another
func (n *Note) Equals(other Note) bool {
	return n.Content == other.Content
}